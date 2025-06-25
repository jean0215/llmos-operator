package registry

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"

	mlv1 "github.com/llmos-ai/llmos-operator/pkg/apis/ml.llmos.ai/v1"
	"github.com/llmos-ai/llmos-operator/pkg/registry/backend"
	"github.com/llmos-ai/llmos-operator/pkg/registry/backend/huggingface"
	"github.com/llmos-ai/llmos-operator/pkg/registry/backend/modelscope"
	"github.com/llmos-ai/llmos-operator/pkg/registry/backend/s3"
)

const (
	defaultSecretNamespace = "llmos-system"

	accessKeyIDName     = "accessKeyID"
	accessKeySecretName = "accessKeySecret"
	tokenName           = "token"

	// Backend types
	BackendTypeS3          = "S3"
	BackendTypeHuggingFace = "HuggingFace"
	BackendTypeModelScope  = "ModelScope"
)

type RegistryGetter func(name string) (*mlv1.Registry, error)
type SecretGetter func(namespace, name string) (*corev1.Secret, error)

type Manager struct {
	RegistryGetter
	SecretGetter
}

func NewManager(sg SecretGetter, rg RegistryGetter) *Manager {
	return &Manager{
		rg,
		sg,
	}
}

func (r *Manager) NewBackend(ctx context.Context, registry *mlv1.Registry) (backend.Backend, error) {
	switch registry.Spec.BackendType {
	case BackendTypeS3:
		// Get the secret containing access credentials from llmos-system namespace
		id, secret, err := getAccessKey(r.SecretGetter, registry.Spec.S3Config.AccessCredentialSecretName)
		if err != nil {
			return nil, fmt.Errorf("get access key failed: %w", err)
		}

		return s3.NewMinioClient(ctx, registry.Spec.S3Config.Endpoint, id, secret,
			registry.Spec.S3Config.Bucket, registry.Spec.S3Config.UseSSL)

	case BackendTypeHuggingFace:
		// Get the token from secret if specified
		token := ""
		if registry.Spec.HuggingFaceConfig.TokenSecretName != "" {
			var err error
			token, err = getHuggingFaceToken(r.SecretGetter, registry.Spec.HuggingFaceConfig.TokenSecretName)
			if err != nil {
				return nil, fmt.Errorf("get hugging face token failed: %w", err)
			}
		}

		return huggingface.NewHuggingFaceClient(token)

	case BackendTypeModelScope:
		// Get the token from secret if specified
		token := ""
		if registry.Spec.ModelScopeConfig.TokenSecretName != "" {
			var err error
			token, err = getModelScopeToken(r.SecretGetter, registry.Spec.ModelScopeConfig.TokenSecretName)
			if err != nil {
				return nil, fmt.Errorf("get modelscope token failed: %w", err)
			}
		}

		return modelscope.NewModelScopeClient(token)

	default:
		return nil, fmt.Errorf("unsupported backend type: %s", registry.Spec.BackendType)
	}
}

func getAccessKey(sg SecretGetter, accessCredentialSecretName string) (string, string, error) {
	secret, err := sg(defaultSecretNamespace, accessCredentialSecretName)
	if err != nil {
		if errors.IsNotFound(err) {
			return "", "", fmt.Errorf("secret %s not found in llmos-system namespace", accessCredentialSecretName)
		}
		return "", "", fmt.Errorf("get secret failed: %w", err)
	}

	// Extract credentials from the secret
	accessKeyID, ok := secret.Data[accessKeyIDName]
	if !ok {
		return "", "", fmt.Errorf("secret %s does not contain %s key", accessCredentialSecretName, accessKeyIDName)
	}

	accessKeySecret, ok := secret.Data[accessKeySecretName]
	if !ok {
		return "", "", fmt.Errorf("secret %s does not contain %s key", accessCredentialSecretName, accessKeySecretName)
	}

	return string(accessKeyID), string(accessKeySecret), nil
}

// getToken is a generic function to extract token from secret
func getToken(sg SecretGetter, tokenSecretName string) (string, error) {
	secret, err := sg(defaultSecretNamespace, tokenSecretName)
	if err != nil {
		if errors.IsNotFound(err) {
			return "", fmt.Errorf("secret %s not found in llmos-system namespace", tokenSecretName)
		}
		return "", fmt.Errorf("get secret failed: %w", err)
	}

	// Extract token from the secret
	token, ok := secret.Data[tokenName]
	if !ok {
		return "", fmt.Errorf("secret %s does not contain %s key", tokenSecretName, tokenName)
	}

	return string(token), nil
}

func getHuggingFaceToken(sg SecretGetter, tokenSecretName string) (string, error) {
	return getToken(sg, tokenSecretName)
}

func getModelScopeToken(sg SecretGetter, tokenSecretName string) (string, error) {
	return getToken(sg, tokenSecretName)
}

func (r *Manager) NewBackendFromRegistry(ctx context.Context, registryName string) (backend.Backend, error) {
	registry, err := r.RegistryGetter(registryName)
	if err != nil {
		return nil, fmt.Errorf("get registry %s failed: %w", registryName, err)
	}

	return r.NewBackend(ctx, registry)
}
