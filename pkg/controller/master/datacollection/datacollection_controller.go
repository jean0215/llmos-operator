package datacollection

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"path"
	"reflect"

	ctlbatchv1 "github.com/rancher/wrangler/v3/pkg/generated/controllers/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	agentv1 "github.com/llmos-ai/llmos-operator/pkg/apis/agent.llmos.ai/v1"
	ctlagentv1 "github.com/llmos-ai/llmos-operator/pkg/generated/controllers/agent.llmos.ai/v1"
	"github.com/llmos-ai/llmos-operator/pkg/registry"
	"github.com/llmos-ai/llmos-operator/pkg/registry/backend"
	"github.com/llmos-ai/llmos-operator/pkg/server/config"
)

// TODO: Docling is triggered to process files through changes in sourceFiles. Frequent changes will result in
// the need to frequently start docling containers or processes. Therefore, it is necessary to consider how to
// reduce the number of docling processing times.

const (
	sourceFilesSubDirectory    = "sourceFiles"
	processedFilesSubDirectory = "processedFiles"
)

type handler struct {
	ctx context.Context

	dataCollectionClient ctlagentv1.DataCollectionClient
	jobClient            ctlbatchv1.JobClient

	rm *registry.Manager
}

func Register(_ context.Context, mgmt *config.Management, _ config.Options) error {
	registries := mgmt.LLMFactory.Ml().V1().Registry()
	secrets := mgmt.CoreFactory.Core().V1().Secret()
	datacollections := mgmt.AgentFactory.Agent().V1().DataCollection()
	jobs := mgmt.BatchFactory.Batch().V1().Job()

	h := handler{
		ctx: mgmt.Ctx,

		dataCollectionClient: datacollections,
		jobClient:            jobs,
	}
	h.rm = registry.NewManager(secrets.Cache().Get, registries.Cache().Get)

	datacollections.OnChange(mgmt.Ctx, "datacollection.OnChange", h.OnChange)
	datacollections.OnRemove(mgmt.Ctx, "datacollection.OnRemove", h.OnRemove)

	return nil
}

// OnChange is called when an DataCollection object is created or updated.
func (h *handler) OnChange(_ string, ad *agentv1.DataCollection) (*agentv1.DataCollection, error) {
	if ad == nil || ad.DeletionTimestamp != nil {
		return ad, nil
	}

	adCopy := ad.DeepCopy()
	err := h.ensureRootPath(adCopy)
	if err != nil {
		return h.updateDataCollectionStatus(adCopy, ad, err)
	}

	filesRemoved, filesToProcess := deltaFiles(ad)
	err = h.removeProcessedFiles(adCopy, filesRemoved)
	if err != nil {
		return h.updateDataCollectionStatus(adCopy, ad, err)
	}
	err = h.processNewFiles(adCopy, filesToProcess)
	if err != nil {
		return h.updateDataCollectionStatus(adCopy, ad, err)
	}

	return h.updateDataCollectionStatus(adCopy, ad, err)
}

func (h *handler) OnRemove(_ string, ad *agentv1.DataCollection) (*agentv1.DataCollection, error) {
	if ad == nil || ad.DeletionTimestamp == nil {
		return ad, nil
	}

	b, err := h.rm.NewBackendFromRegistry(h.ctx, ad.Spec.Registry)
	if err != nil {
		return nil, fmt.Errorf("failed to create backend from registry %s: %w", ad.Spec.Registry, err)
	}
	if err := b.Delete(h.ctx, getRootPath(ad.Namespace, ad.Name)); err != nil {
		return nil, fmt.Errorf("failed to delete directory %s in registry %s: %w", getRootPath(ad.Namespace, ad.Name), ad.Spec.Registry, err)
	}

	return ad, nil
}

func (h *handler) ensureRootPath(ad *agentv1.DataCollection) error {
	// the rootPath recorded in the status is the sourceFiles directory
	rootPath := getSourceFilePath(ad.Namespace, ad.Name)
	processedFilesPath := getProcessedFilePath(ad.Namespace, ad.Name)
	if agentv1.Ready.IsTrue(ad) && ad.Status.RootPath == rootPath {
		return nil
	}
	b, err := h.rm.NewBackendFromRegistry(h.ctx, ad.Spec.Registry)
	if err != nil {
		return fmt.Errorf("failed to create backend from registry %s: %w", ad.Spec.Registry, err)
	}
	if err := b.CreateDirectory(h.ctx, rootPath); err != nil {
		return fmt.Errorf("failed to create directory %s in registry %s: %w", rootPath, ad.Spec.Registry, err)
	}
	if err := b.CreateDirectory(h.ctx, processedFilesPath); err != nil {
		return fmt.Errorf("failed to create directory %s in registry %s: %w", processedFilesPath, ad.Spec.Registry, err)
	}

	ad.Status.RootPath = rootPath
	agentv1.Ready.True(ad)
	return nil
}

func (h *handler) updateDataCollectionStatus(adCopy, ad *agentv1.DataCollection, err error) (*agentv1.DataCollection, error) {
	if err == nil {
		agentv1.Ready.True(adCopy)
		agentv1.Ready.Message(adCopy, "")
	} else {
		agentv1.Ready.False(adCopy)
		agentv1.Ready.Message(adCopy, err.Error())
	}

	// don't update when no change happens
	if reflect.DeepEqual(adCopy.Status, ad.Status) {
		return adCopy, err
	}

	updatedAd, updateErr := h.dataCollectionClient.UpdateStatus(adCopy)
	if updateErr != nil {
		return nil, fmt.Errorf("update application data status failed: %w", updateErr)
	}
	return updatedAd, err
}

func fileUid(registry, path, etag string) string {
	data := fmt.Appendf([]byte{}, "%s~%s~%s", registry, path, etag)
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}

// compare ad.Spec.SourceFiles with ad.Status.PreprocessedFiles
// to find the removed files and the files to process.
func deltaFiles(ad *agentv1.DataCollection) (filesRemoved []string, filesToProcess []agentv1.PreprocessedFile) {
	pfs := make(map[string]bool)
	for _, pf := range ad.Status.PreprocessedFiles {
		pfs[pf.UID] = false
	}

	for _, sf := range ad.Spec.SourceFiles {
		uid := fileUid(ad.Spec.Registry, sf.Path, sf.Etag)
		if _, ok := pfs[uid]; !ok {
			filesToProcess = append(filesToProcess, agentv1.PreprocessedFile{
				UID: uid,
				SourceFile: agentv1.FileInfo{
					Path: sf.Path,
					ETag: sf.Etag,
				},
			})
		} else {
			pfs[uid] = true
		}
	}
	for uid, found := range pfs {
		if !found {
			filesRemoved = append(filesRemoved, uid)
		}
	}

	return
}

func (h *handler) removeProcessedFiles(ad *agentv1.DataCollection, filesRemoved []string) error {
	if len(filesRemoved) == 0 {
		return nil
	}

	// Create a map for faster lookup of UIDs to remove
	removeMap := make(map[string]bool)
	for _, uid := range filesRemoved {
		removeMap[uid] = true
	}

	// Get backend client for S3 operations
	b, err := h.rm.NewBackendFromRegistry(h.ctx, ad.Spec.Registry)
	if err != nil {
		return fmt.Errorf("failed to create backend from registry %s: %w", ad.Spec.Registry, err)
	}

	// Find processed files to delete and delete them from S3
	updatedProcessedFiles := make([]agentv1.PreprocessedFile, 0)
	for _, pf := range ad.Status.PreprocessedFiles {
		if removeMap[pf.UID] {
			// Delete the processed file from S3
			if err := b.Delete(h.ctx, pf.ProcessedFilePath); err != nil {
				return fmt.Errorf("failed to delete processed file %s from registry: %w", pf.ProcessedFilePath, err)
			}
		} else {
			// Keep files that are not in the removal list
			updatedProcessedFiles = append(updatedProcessedFiles, pf)
		}
	}

	// Update the status with filtered processed files
	ad.Status.PreprocessedFiles = updatedProcessedFiles

	return nil
}

func (h *handler) processNewFiles(ad *agentv1.DataCollection, filesToProcess []agentv1.PreprocessedFile) error {
	if len(filesToProcess) == 0 {
		return nil
	}

	// Get backend client for file operations
	b, err := h.rm.NewBackendFromRegistry(h.ctx, ad.Spec.Registry)
	if err != nil {
		return fmt.Errorf("failed to create backend from registry %s: %w", ad.Spec.Registry, err)
	}

	// Add files to status with validation
	if err := h.addFilesToStatus(ad, filesToProcess, b); err != nil {
		return fmt.Errorf("failed to add files to status: %w", err)
	}

	// Create docling job to process the files
	/*
		if err := h.createDoclingJob(ad, filesToProcess); err != nil {
			return fmt.Errorf("failed to create docling job: %w", err)
		}
	*/

	return nil
}

// addFilesToStatus validates files against registry and adds them to status
func (h *handler) addFilesToStatus(dc *agentv1.DataCollection, filesToProcess []agentv1.PreprocessedFile, b backend.Backend) error {
	validatedFiles := make([]agentv1.PreprocessedFile, 0, len(filesToProcess))

	for _, file := range filesToProcess {
		// List file from registry to validate etag
		registryFiles, err := b.List(h.ctx, file.SourceFile.Path, false, false)
		if err != nil {
			return fmt.Errorf("failed to list file %s from registry: %w", file.SourceFile.Path, err)
		}

		// Skip if file not found or etag mismatch
		if len(registryFiles) == 0 {
			continue
		}

		registryFile := registryFiles[0]
		if registryFile.ETag != file.SourceFile.ETag {
			// Ignore files with different etag
			continue
		}

		// Convert backend FileInfo to app FileInfo and complete the PreprocessedFile
		completedFile := file
		completedFile.SourceFile = h.convertToAppFileInfo(registryFile)
		// ProcessedFile remains empty as per requirement
		completedFile.ProcessedFilePath = path.Join(getProcessedFilePath(dc.Namespace, dc.Name), completedFile.SourceFile.Name+".md")
		// Set initial condition using ready condition
		agentv1.Ready.True(&completedFile)

		validatedFiles = append(validatedFiles, completedFile)
	}

	// Add validated files to status
	dc.Status.PreprocessedFiles = append(dc.Status.PreprocessedFiles, validatedFiles...)

	return nil
}

// convertToAppFileInfo converts backend.FileInfo to agentv1.FileInfo
func (h *handler) convertToAppFileInfo(backendFile backend.FileInfo) agentv1.FileInfo {
	return agentv1.FileInfo{
		Name:         backendFile.Name,
		Path:         backendFile.Path,
		Size:         backendFile.Size,
		IsDir:        backendFile.IsDir,
		LastModified: metav1.NewTime(backendFile.LastModified),
		ContentType:  backendFile.ContentType,
		ETag:         backendFile.ETag,
	}
}

// TODO: createDoclingJob creates a Kubernetes job to process the files
func (h *handler) createDoclingJob(dc *agentv1.DataCollection, filesToProcess []agentv1.PreprocessedFile) error {
	return nil
}

func getRootPath(namespace, name string) string {
	return path.Join(agentv1.DataCollectionResourceName, namespace, name)
}

func getSourceFilePath(namespace, name string) string {
	return path.Join(agentv1.DataCollectionResourceName, namespace, name, sourceFilesSubDirectory)
}

func getProcessedFilePath(namespace, name string) string {
	return path.Join(agentv1.DataCollectionResourceName, namespace, name, processedFilesSubDirectory)
}
