package registry

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"path"
	"strings"

	"github.com/gorilla/mux"
	"github.com/rancher/apiserver/pkg/apierror"
	"github.com/rancher/wrangler/v3/pkg/schemas/validation"
	"github.com/sirupsen/logrus"

	"github.com/llmos-ai/llmos-operator/pkg/registry"
	"github.com/llmos-ai/llmos-operator/pkg/registry/backend"
	"github.com/llmos-ai/llmos-operator/pkg/utils"
)

const (
	ActionUpload          = "upload"
	ActionDownload        = "download"
	ActionList            = "list"
	ActionRemove          = "remove"
	ActionCreateDirectory = "createDirectory"
)

type UploadInput struct {
	SourceFilePath string `json:"sourceFilePath"`
	// if empty, use version as target directory
	TargetDirectory string `json:"targetDirectory"`
}
type DownloadInput struct {
	TargetFilePath string `json:"targetFilePath"`
}
type ListInput DownloadInput
type RemoveInput DownloadInput

type CreateDirectoryInput struct {
	TargetDirectory string `json:"targetDirectory"`
}

type BaseHandler struct {
	Ctx                    context.Context
	RegistryManager        *registry.Manager
	GetRegistryAndRootPath func(namespace, name string) (string, string, error)
}

func (h BaseHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if err := h.do(rw, req); err != nil {
		logrus.Errorf("do action failed: %v", err)
		status := http.StatusInternalServerError
		var e *apierror.APIError
		if errors.As(err, &e) {
			status = e.Code.Status
		}
		utils.ResponseAPIError(rw, status, e)
		return
	}
	// if rw has content, it will be written to response writer and return 200, otherwise, return 204
	utils.ResponseOKWithNoContent(rw)
}

func (h BaseHandler) do(rw http.ResponseWriter, req *http.Request) error {
	vars := utils.EncodeVars(mux.Vars(req))

	if req.Method == http.MethodPost {
		return h.doPost(rw, req, vars)
	}

	return apierror.NewAPIError(validation.InvalidAction, fmt.Sprintf("Unsupported method %s", req.Method))
}

func (h BaseHandler) doPost(rw http.ResponseWriter, req *http.Request, vars map[string]string) error {
	action := vars["action"]
	namespace, name := vars["namespace"], vars["name"]

	switch action {
	case ActionUpload:
		return h.upload(req, namespace, name)
	case ActionDownload:
		return h.download(rw, req, namespace, name)
	case ActionList:
		return h.list(rw, req, namespace, name)
	case ActionRemove:
		return h.remove(req, namespace, name)
	case ActionCreateDirectory:
		return h.createDirectory(req, namespace, name)
	default:
		return apierror.NewAPIError(validation.InvalidAction, fmt.Sprintf("Unsupported action %s", action))
	}
}

func (h BaseHandler) upload(req *http.Request, namespace, name string) error {
	input := &UploadInput{}

	err := decodeAndValidateInput(req, input, input.TargetDirectory)
	if err != nil {
		return err
	}

	b, rootPath, err := h.getBackendAndRootPath(namespace, name)
	if err != nil {
		return err
	}

	if err := b.Upload(h.Ctx, input.SourceFilePath, path.Join(rootPath, input.TargetDirectory)); err != nil {
		return fmt.Errorf("upload file %s failed: %w", input.SourceFilePath, err)
	}

	return nil
}

func (h BaseHandler) download(rw http.ResponseWriter, req *http.Request, namespace, name string) error {
	input := &DownloadInput{}

	err := decodeAndValidateInput(req, input, input.TargetFilePath)
	if err != nil {
		return err
	}

	b, rootPath, err := h.getBackendAndRootPath(namespace, name)
	if err != nil {
		return err
	}

	objectName := path.Join(rootPath, input.TargetFilePath)

	if err := b.Download(h.Ctx, objectName, rw); err != nil {
		return fmt.Errorf("download file %s failed: %w", input.TargetFilePath, err)
	}

	return nil
}

func (h BaseHandler) list(rw http.ResponseWriter, req *http.Request, namespace, name string) error {
	input := &ListInput{}
	err := decodeAndValidateInput(req, input, input.TargetFilePath)
	if err != nil {
		return err
	}

	b, rootPath, err := h.getBackendAndRootPath(namespace, name)
	if err != nil {
		return err
	}

	objectName := path.Join(rootPath, input.TargetFilePath)
	output, err := b.List(h.Ctx, objectName, false, true)
	if err != nil {
		return fmt.Errorf("failed to list %s: %w", input.TargetFilePath, err)
	}

	utils.ResponseOKWithBody(rw, output)

	return nil
}

func (h BaseHandler) remove(req *http.Request, namespace, name string) error {
	input := &RemoveInput{}

	err := decodeAndValidateInput(req, input, input.TargetFilePath)
	if err != nil {
		return err
	}

	b, rootPath, err := h.getBackendAndRootPath(namespace, name)
	if err != nil {
		return err
	}

	objectName := path.Join(rootPath, input.TargetFilePath)
	if err := b.Delete(h.Ctx, objectName); err != nil {
		return fmt.Errorf("remove file %s failed: %v", objectName, err)
	}

	return nil
}

func (h BaseHandler) createDirectory(req *http.Request, namespace, name string) error {
	input := &CreateDirectoryInput{}

	err := decodeAndValidateInput(req, input, input.TargetDirectory)
	if err != nil {
		return err
	}

	ctx := req.Context()
	b, rootPath, err := h.getBackendAndRootPath(namespace, name)
	if err != nil {
		return err
	}

	directory := path.Join(rootPath, input.TargetDirectory)
	if err := b.CreateDirectory(ctx, directory); err != nil {
		return fmt.Errorf("create directory %s failed: %w", directory, err)
	}

	return nil
}

func decodeAndValidateInput(req *http.Request, input interface{}, pathField string) error {
	if err := json.NewDecoder(req.Body).Decode(input); err != nil {
		return apierror.NewAPIError(validation.InvalidBodyContent, fmt.Sprintf("Failed to parse body: %v", err))
	}

	if !isValidPath(pathField) {
		return apierror.NewAPIError(validation.InvalidBodyContent, "Invalid path")
	}

	return nil
}

func (h BaseHandler) getBackendAndRootPath(namespace, name string) (backend.Backend, string, error) {
	reg, rootPath, err := h.GetRegistryAndRootPath(namespace, name)
	if err != nil {
		return nil, "", fmt.Errorf("get registry and root path failed: %w", err)
	}

	b, err := h.RegistryManager.NewBackendFromRegistry(h.Ctx, reg)
	if err != nil {
		return nil, "", fmt.Errorf("new backend for registry %s failed: %w", reg, err)
	}

	return b, rootPath, nil
}

// isValidPath checks if the path is valid and doesn't contain directory traversal attempts
func isValidPath(p string) bool {
	// Prevent paths with ".." which could lead to directory traversal
	if strings.Contains(p, "..") {
		return false
	}

	// Prevent absolute paths
	if strings.HasPrefix(p, "/") {
		return false
	}

	return true
}
