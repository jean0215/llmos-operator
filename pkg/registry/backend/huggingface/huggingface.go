package huggingface

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

// HuggingFaceClient represents a Hugging Face client
type HuggingFaceClient struct {
	token string
}

var _ backend.Backend = (*HuggingFaceClient)(nil)

// NewHuggingFaceClient initializes a new Hugging Face client
func NewHuggingFaceClient(token string) (backend.Backend, error) {
	return &HuggingFaceClient{
		token: token,
	}, nil
}

// IncrementalDownload downloads a model from Hugging Face using huggingface-hub CLI
func (hf *HuggingFaceClient) IncrementalDownload(ctx context.Context, modelName, outputDir string, concurrency int) error {
	// Ensure output directory exists
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("create output directory %s failed: %w", outputDir, err)
	}

	// Check if huggingface-hub CLI is available
	if _, err := exec.LookPath("huggingface-cli"); err != nil {
		return fmt.Errorf("huggingface-cli not found in PATH, please install it: %w", err)
	}

	// Prepare the download command
	cmd := exec.CommandContext(ctx, "huggingface-cli", "download", modelName, "--local-dir", outputDir)
	
	// Set environment variables
	env := os.Environ()
	if hf.token != "" {
		env = append(env, fmt.Sprintf("HF_TOKEN=%s", hf.token))
	}
	cmd.Env = env

	// Set working directory
	cmd.Dir = outputDir

	logrus.Infof("Downloading model %s to %s using huggingface-cli", modelName, outputDir)

	// Execute the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to download model %s: %w, output: %s", modelName, err, string(output))
	}

	logrus.Infof("Successfully downloaded model %s to %s", modelName, outputDir)
	return nil
}

// Download downloads a single file or directory to a writer (not implemented for HF)
func (hf *HuggingFaceClient) Download(ctx context.Context, src string, rw io.Writer) error {
	return fmt.Errorf("Download method not implemented for Hugging Face backend")
}

// Upload uploads a file to Hugging Face (not implemented)
func (hf *HuggingFaceClient) Upload(ctx context.Context, src, dst string) error {
	return fmt.Errorf("Upload method not implemented for Hugging Face backend")
}

// UploadFromReader uploads data from a reader to Hugging Face (not implemented)
func (hf *HuggingFaceClient) UploadFromReader(ctx context.Context, reader io.Reader, dst string, size int64, contentType string) error {
	return fmt.Errorf("UploadFromReader method not implemented for Hugging Face backend")
}

// Delete deletes a file from Hugging Face (not implemented)
func (hf *HuggingFaceClient) Delete(ctx context.Context, objectName string) error {
	return fmt.Errorf("Delete method not implemented for Hugging Face backend")
}

// List lists files in a repository using HuggingFace API
func (hf *HuggingFaceClient) List(ctx context.Context, prefix string, recursive, skipItself bool) ([]backend.FileInfo, error) {
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

	// Construct the API URL - use the correct HuggingFace API endpoint with proper escaping
	apiURL := fmt.Sprintf("https://huggingface.co/api/models/%s", url.PathEscape(modelName))
	
	// Create request
	req, err := http.NewRequestWithContext(ctx, "GET", apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("create request failed: %w", err)
	}

	// Add authorization header if token is available
	if hf.token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", hf.token))
	}

	// Make the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get model info from HuggingFace API for %s: %w", modelName, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HuggingFace API returned status %d for model %s", resp.StatusCode, modelName)
	}

	// Parse the response
	var modelInfo ModelInfo
	if err := json.NewDecoder(resp.Body).Decode(&modelInfo); err != nil {
		return nil, fmt.Errorf("failed to parse model info response for %s: %w", modelName, err)
	}

	// Convert to FileInfo slice and get file sizes using HEAD requests
	var files []backend.FileInfo
	for _, sibling := range modelInfo.Siblings {
		// Get file size using HEAD request
		fileSize, err := hf.getFileSize(ctx, modelName, sibling.Filename, client)
		if err != nil {
			logrus.Warnf("Failed to get size for file %s: %v", sibling.Filename, err)
			// Continue with size 0 if we can't get the actual size
			fileSize = 0
		}
		
		files = append(files, backend.FileInfo{
			Name: sibling.Filename,
			Size: fileSize,
			Path: sibling.Filename,
			IsDir: false,
		})
	}

	return files, nil
}

// Copy copies a file (not implemented)
func (hf *HuggingFaceClient) Copy(ctx context.Context, src, dst string) error {
	return fmt.Errorf("Copy method not implemented for Hugging Face backend")
}

// CreateDirectory creates a directory (not implemented)
func (hf *HuggingFaceClient) CreateDirectory(ctx context.Context, path string) error {
	return fmt.Errorf("CreateDirectory method not implemented for Hugging Face backend")
}

// DeleteDirectory deletes a directory (not implemented)
func (hf *HuggingFaceClient) DeleteDirectory(ctx context.Context, path string) error {
	return fmt.Errorf("DeleteDirectory method not implemented for Hugging Face backend")
}

// GetObjectURL gets the URL of an object
func (hf *HuggingFaceClient) GetObjectURL(objectName string) string {
	return fmt.Sprintf("https://huggingface.co/%s", url.PathEscape(objectName))
}

// ModelInfo represents the response from HuggingFace model info API
type ModelInfo struct {
	Siblings []RepoFile `json:"siblings"`
}

// RepoFile represents a file in a HuggingFace repository
type RepoFile struct {
	Filename string `json:"rfilename"`
	Size     int64  `json:"size,omitempty"`
	Oid      string `json:"oid,omitempty"`
}

// getFileSize gets the size of a specific file using HEAD request
func (hf *HuggingFaceClient) getFileSize(ctx context.Context, modelName, filename string, client *http.Client) (int64, error) {
	// Construct the file URL with proper escaping
	fileURL := fmt.Sprintf("https://huggingface.co/%s/resolve/main/%s", url.PathEscape(modelName), url.PathEscape(filename))
	
	// Create HEAD request
	req, err := http.NewRequestWithContext(ctx, "HEAD", fileURL, nil)
	if err != nil {
		return 0, fmt.Errorf("create HEAD request failed: %w", err)
	}

	// Add authorization header if token is available
	if hf.token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", hf.token))
	}

	// Make the HEAD request
	resp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("HEAD request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("HEAD request returned status %d for file %s", resp.StatusCode, filename)
	}

	// Get content length from header
	contentLength := resp.Header.Get("Content-Length")
	if contentLength == "" {
		return 0, fmt.Errorf("no Content-Length header found for file %s", filename)
	}

	// Parse content length
	var size int64
	if _, err := fmt.Sscanf(contentLength, "%d", &size); err != nil {
		return 0, fmt.Errorf("failed to parse Content-Length %s: %w", contentLength, err)
	}

	return size, nil
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

// GetSize gets the size of a model from Hugging Face using the List method
func (hf *HuggingFaceClient) GetSize(ctx context.Context, path string) (int64, error) {
	if path == "" {
		return 0, fmt.Errorf("model path is required")
	}

	// Use the List method to get file information
	files, err := hf.List(ctx, path, false, false)
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
		logrus.Infof("Retrieved actual size %d bytes for HuggingFace model %s", totalSize, path)
		return totalSize, nil
	}

	// Return error if no size information available
	return 0, fmt.Errorf("no size information available for model %s", path)
}