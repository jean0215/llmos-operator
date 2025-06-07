package datacollection

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	cr "github.com/llmos-ai/llmos-operator/pkg/api/common/registry"
	agentv1 "github.com/llmos-ai/llmos-operator/pkg/apis/agent.llmos.ai/v1"
	ctlagentv1 "github.com/llmos-ai/llmos-operator/pkg/generated/controllers/agent.llmos.ai/v1"
	ctlmlv1 "github.com/llmos-ai/llmos-operator/pkg/generated/controllers/ml.llmos.ai/v1"
	"github.com/llmos-ai/llmos-operator/pkg/registry"
	"github.com/llmos-ai/llmos-operator/pkg/registry/backend"
	"github.com/llmos-ai/llmos-operator/pkg/server/config"
	"github.com/llmos-ai/llmos-operator/pkg/utils"
)

type DataCollectionGetter func(namespace, name string) (*agentv1.DataCollection, error)

type Handler struct {
	dcCache       ctlagentv1.DataCollectionCache
	dcClient      ctlagentv1.DataCollectionClient
	registryCache ctlmlv1.RegistryCache

	cr.BaseHandler
}

func NewHandler(scaled *config.Scaled) Handler {
	dcs := scaled.Management.AgentFactory.Agent().V1().DataCollection()
	h := Handler{
		dcClient: dcs,
		dcCache:  dcs.Cache(),
	}

	registryCache := scaled.Management.LLMFactory.Ml().V1().Registry().Cache()
	secretCache := scaled.CoreFactory.Core().V1().Secret().Cache()

	h.BaseHandler = cr.BaseHandler{
		Ctx:                    scaled.Ctx,
		GetRegistryAndRootPath: h.GetRegistryAndRootPath,
		RegistryManager:        registry.NewManager(secretCache.Get, registryCache.Get),
		PostHooks: map[string]cr.PostHook{
			cr.ActionUpload: h.SyncFiles,
			cr.ActionRemove: h.SyncFiles,
		},
	}

	return h
}

func (h Handler) SyncFiles(req *http.Request, b backend.Backend) error {
	vars := utils.EncodeVars(mux.Vars(req))
	namespace, name := vars["namespace"], vars["name"]

	logrus.Infof("sync files for datacollection %s/%s", namespace, name)
	dc, err := h.dcCache.Get(namespace, name)
	if err != nil {
		return fmt.Errorf("get datacollection %s/%s failed: %w", namespace, name, err)
	}
	// Use req.Context(), not h.Ctx. I don't know why.
	files, err := b.List(req.Context(), dc.Status.RootPath, true, true)
	if err != nil {
		return fmt.Errorf("failed to list files in %s from registry %s: %w", dc.Spec.Registry, dc.Status.RootPath, err)
	}

	sourceFiles := make([]agentv1.SourceFile, 0, len(files))
	for _, file := range files {
		sourceFiles = append(sourceFiles, agentv1.SourceFile{
			Etag: file.ETag,
			Path: file.Path,
		})
	}

	dcCopy := dc.DeepCopy()
	dcCopy.Spec.SourceFiles = sourceFiles
	if _, err = h.dcClient.Update(dcCopy); err != nil {
		return fmt.Errorf("failed to update application data %s/%s: %w", namespace, name, err)
	}

	return nil
}

func (h Handler) GetRegistryAndRootPath(namespace, name string) (string, string, error) {
	return GetRegistryAndRootPath(h.dcCache.Get, namespace, name)
}

func GetRegistryAndRootPath(adGetter DataCollectionGetter, namespace, name string) (string, string, error) {
	ad, err := adGetter(namespace, name)
	if err != nil {
		return "", "", fmt.Errorf("failed to get application data %s/%s: %w", namespace, name, err)
	}

	if !agentv1.Ready.IsTrue(ad) {
		return "", "", fmt.Errorf("application data %s/%s is not ready", namespace, name)
	}

	return ad.Spec.Registry, ad.Status.RootPath, nil
}
