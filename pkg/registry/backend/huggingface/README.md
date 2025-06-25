# Hugging Face Backend for LLMOS Registry

This package implements a Hugging Face backend for the LLMOS registry system, enabling efficient model downloads from Hugging Face Hub.

## Prerequisites

- Install `huggingface-hub`: `pip install huggingface-hub` or `pipx install huggingface-hub`
- Ensure `huggingface-cli` is available in your PATH

## Usage

### Registry Configuration

Create Registry resources for different use cases:

#### Public Models Registry

```yaml
apiVersion: ml.llmos.ai/v1
kind: Registry
metadata:
  name: huggingface-public
spec:
  backendType: HuggingFace
  huggingFaceConfig: {}
    # No token needed for public models
```

#### Private Models Registry

```yaml
apiVersion: ml.llmos.ai/v1
kind: Registry
metadata:
  name: huggingface-private
spec:
  backendType: HuggingFace
  huggingFaceConfig:
    tokenSecretName: "hf-token-secret"
```

For private models, create a secret with your token:

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: hf-token-secret
  namespace: llmos-system
type: Opaque
data:
  # Get your token from: https://huggingface.co/settings/tokens
  # echo -n "your-hf-token" | base64
  token: <base64-encoded-huggingface-token>
```

### Downloading Models

Use the downloader with registry name:

```bash
# Download public model
./downloader download --registry=huggingface-public --name=microsoft/DialoGPT-small --output-dir=/path/to/output

# Download private model
./downloader download --registry=huggingface-private --name=private-org/private-model --output-dir=/path/to/output
```

### Programmatic Usage

The backend integrates with the registry manager:

```go
manager := registry.NewManager(secretGetter, registryGetter)
backend, err := manager.NewBackendFromRegistry(ctx, "huggingface-public")
if err != nil {
    return err
}

err = backend.IncrementalDownload(ctx, "microsoft/DialoGPT-small", "/path/to/output", 3)
```

## Features

- **Registry Integration**: Full integration with LLMOS registry CRD system
- **Token Management**: Secure token storage via Kubernetes secrets
- **Public/Private Support**: Supports both public and private models
- **Incremental Downloads**: Uses `huggingface-cli download` for efficient downloading
- **Concurrent Downloads**: Configurable thread count for parallel downloads

## Limitations

- Only `IncrementalDownload` method is implemented
- Requires `huggingface-cli` to be installed and available in PATH
- No true incremental downloads (downloads entire model each time)
- Other backend methods (Upload, Delete, List, etc.) are not supported

## Implementation Details

The backend uses the `huggingface-cli download` command under the hood:

```bash
huggingface-cli download model-name --local-dir /output/path
```

For authenticated requests, the `HF_TOKEN` environment variable is set before executing the command.

## Configuration Examples

See `sample/registry/huggingface-registry.yaml` for complete configuration examples.