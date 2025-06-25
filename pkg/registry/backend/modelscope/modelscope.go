package modelscope

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/llmos-ai/llmos-operator/pkg/registry/backend"
)

// ModelScopeClient represents a ModelScope client
type ModelScopeClient struct {
	token string
}

var _ backend.Backend = (*ModelScopeClient)(nil)

// NewModelScopeClient initializes a new ModelScope client
func NewModelScopeClient(token string) (backend.Backend, error) {
	return &ModelScopeClient{
		token: token,
	}, nil
}

// IncrementalDownload downloads a model from ModelScope using modelscope CLI
func (ms *ModelScopeClient) IncrementalDownload(ctx context.Context, modelName, outputDir string, concurrency int) error {
	// Ensure output directory exists
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("create output directory %s failed: %w", outputDir, err)
	}

	// Check if modelscope CLI is available
	if _, err := exec.LookPath("modelscope"); err != nil {
		return fmt.Errorf("modelscope CLI not found in PATH, please install it: %w", err)
	}

	// Prepare the download command
	cmd := exec.CommandContext(ctx, "modelscope", "download", "--model", modelName, "--local_dir", outputDir)
	
	// Set environment variables
	env := os.Environ()
	if ms.token != "" {
		env = append(env, fmt.Sprintf("MODELSCOPE_API_TOKEN=%s", ms.token))
	}
	cmd.Env = env

	// Set working directory
	cmd.Dir = outputDir

	logrus.Infof("Downloading model %s to %s using modelscope CLI", modelName, outputDir)

	// Execute the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to download model %s: %w, output: %s", modelName, err, string(output))
	}

	logrus.Infof("Successfully downloaded model %s to %s", modelName, outputDir)
	return nil
}

// Download downloads a single file or directory to a writer (not implemented for ModelScope)
func (ms *ModelScopeClient) Download(ctx context.Context, src string, rw io.Writer) error {
	return fmt.Errorf("Download method not implemented for ModelScope backend")
}

// Upload uploads a file to ModelScope (not implemented)
func (ms *ModelScopeClient) Upload(ctx context.Context, src, dst string) error {
	return fmt.Errorf("Upload method not implemented for ModelScope backend")
}

// UploadFromReader uploads data from a reader to ModelScope (not implemented)
func (ms *ModelScopeClient) UploadFromReader(ctx context.Context, reader io.Reader, dst string, size int64, contentType string) error {
	return fmt.Errorf("UploadFromReader method not implemented for ModelScope backend")
}

// Delete deletes a file from ModelScope (not implemented)
func (ms *ModelScopeClient) Delete(ctx context.Context, objectName string) error {
	return fmt.Errorf("Delete method not implemented for ModelScope backend")
}

// List lists files in a repository using ModelScope API
func (ms *ModelScopeClient) List(ctx context.Context, prefix string, recursive, skipItself bool) ([]backend.FileInfo, error) {
	if prefix == "" {
		return nil, fmt.Errorf("model path is required")
	}

	// Clean the model path (remove any leading/trailing slashes)
	modelName := strings.Trim(prefix, "/")
	
	// Validate model name format to prevent injection attacks
	if err := validateModelName(modelName); err != nil {
		return nil, fmt.Errorf("invalid model name: %w", err)
	}
	
	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// Construct the API URL with proper escaping
	apiURL := fmt.Sprintf("https://modelscope.cn/api/v1/models/%s/repo/files?Revision=master", url.PathEscape(modelName))
	
	// Create request
	req, err := http.NewRequestWithContext(ctx, "GET", apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("create request failed: %w", err)
	}

	// Add authorization header if token is available
	if ms.token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", ms.token))
	}

	// Make the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get model info from ModelScope API for %s: %w", modelName, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ModelScope API returned status %d for model %s", resp.StatusCode, modelName)
	}

	// Parse the response
	var modelInfo ModelScopeModelInfo
	if err := json.NewDecoder(resp.Body).Decode(&modelInfo); err != nil {
		return nil, fmt.Errorf("failed to parse model info response for %s: %w", modelName, err)
	}

	// Convert to FileInfo slice
	var files []backend.FileInfo
	for _, file := range modelInfo.Data.Files {
		files = append(files, backend.FileInfo{
			Name: file.Name,
			Size: file.Size,
			Path: file.Path,
			IsDir: false,
		})
	}

	return files, nil
}

// Copy copies a file (not implemented)
func (ms *ModelScopeClient) Copy(ctx context.Context, src, dst string) error {
	return fmt.Errorf("Copy method not implemented for ModelScope backend")
}

// CreateDirectory creates a directory (not implemented)
func (ms *ModelScopeClient) CreateDirectory(ctx context.Context, path string) error {
	return fmt.Errorf("CreateDirectory method not implemented for ModelScope backend")
}

// DeleteDirectory deletes a directory (not implemented)
func (ms *ModelScopeClient) DeleteDirectory(ctx context.Context, path string) error {
	return fmt.Errorf("DeleteDirectory method not implemented for ModelScope backend")
}

// GetObjectURL gets the URL of an object
func (ms *ModelScopeClient) GetObjectURL(objectName string) string {
	return fmt.Sprintf("https://modelscope.cn/models/%s", url.PathEscape(objectName))
}

// ModelScopeModelInfo represents the response from ModelScope model info API
type ModelScopeModelInfo struct {
	Data struct {
		Files []struct {
			Name string `json:"Name"`
			Path string `json:"Path"`
			Size int64  `json:"Size"`
			Type string `json:"Type"`
		} `json:"Files"`
	} `json:"Data"`
}

// GetSize gets the size of a model from ModelScope using the List method
func (ms *ModelScopeClient) GetSize(ctx context.Context, path string) (int64, error) {
	if path == "" {
		return 0, fmt.Errorf("model path is required")
	}

	// Use the List method to get file information
	files, err := ms.List(ctx, path, false, false)
	if err != nil {
		return 0, fmt.Errorf("failed to list files for model %s: %w", path, err)
	}

	// Calculate total size from all files
	var totalSize int64
	for _, file := range files {
		if file.Size > 0 {
			totalSize += file.Size
		}
	}

	if totalSize > 0 {
		logrus.Infof("Retrieved actual size %d bytes for ModelScope model %s", totalSize, path)
		return totalSize, nil
	}

	// Return error if no size information available
	return 0, fmt.Errorf("no size information available for model %s", path)
}

// validateModelName validates the model name format to prevent injection attacks
func validateModelName(modelName string) error {
	if modelName == "" {
		return fmt.Errorf("model name cannot be empty")
	}
	
	// Check for basic format: should contain only alphanumeric, hyphens, underscores, dots, and forward slashes
	// Typical format: organization/model-name or model-name
	validPattern := regexp.MustCompile(`^[a-zA-Z0-9._/-]+$`)
	if !validPattern.MatchString(modelName) {
		return fmt.Errorf("model name contains invalid characters")
	}
	
	// Prevent path traversal attempts
	if strings.Contains(modelName, "..") {
		return fmt.Errorf("model name cannot contain path traversal sequences")
	}
	
	// Prevent URLs or protocol schemes
	if strings.Contains(modelName, "://") {
		return fmt.Errorf("model name cannot contain URL schemes")
	}
	
	return nil
}