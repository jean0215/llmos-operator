# ModelScope Backend

This package provides a ModelScope backend implementation for the LLMOS registry system. It allows downloading models from ModelScope using the ModelScope CLI.

## Prerequisites

1. **ModelScope CLI**: The ModelScope CLI must be installed and available in the system PATH.
   ```bash
   pip install modelscope
   ```

2. **ModelScope Token** (for private models): If you need to access private models, you'll need a ModelScope API token.

## Configuration

### Public Models

For public models, you can create a registry without specifying a token:

```yaml
apiVersion: ml.llmos.ai/v1
kind: Registry
metadata:
  name: modelscope-public
  namespace: llmos-system
spec:
  backendType: ModelScope
  modelScopeConfig: {}
```

### Private Models

For private models, you need to create a secret with your ModelScope token and reference it in the registry:

1. Create a secret with your ModelScope token:
   ```yaml
   apiVersion: v1
   kind: Secret
   metadata:
     name: modelscope-token
     namespace: llmos-system
   type: Opaque
   data:
     token: <base64-encoded-modelscope-token>
   ```

2. Create a registry that references the secret:
   ```yaml
   apiVersion: ml.llmos.ai/v1
   kind: Registry
   metadata:
     name: modelscope-private
     namespace: llmos-system
   spec:
     backendType: ModelScope
     modelScopeConfig:
       tokenSecretName: modelscope-token
   ```

## Usage

### Using the Downloader CLI

```bash
# Download a public model
./downloader --registry modelscope-public --model "damo/nlp_structbert_backbone_base_std"

# Download a private model (requires token)
./downloader --registry modelscope-private --model "your-org/your-private-model"
```

### Programmatic Usage

```go
package main

import (
    "context"
    "github.com/llmos-ai/llmos-operator/pkg/registry/backend/modelscope"
)

func main() {
    // Create a ModelScope client
    client, err := modelscope.NewModelScopeClient("your-token-here")
    if err != nil {
        panic(err)
    }

    // Download a model
    err = client.IncrementalDownload(context.Background(), "damo/nlp_structbert_backbone_base_std", "/path/to/output", 4)
    if err != nil {
        panic(err)
    }
}
```

## Environment Variables

The ModelScope backend uses the following environment variable:

- `MODELSCOPE_API_TOKEN`: The ModelScope API token for authentication (set automatically when using secrets)

## Model Naming

ModelScope models are typically named in the format `organization/model-name`, for example:
- `damo/nlp_structbert_backbone_base_std`
- `modelscope/chatglm2-6b`
- `your-org/your-model`

## Limitations

Currently, the ModelScope backend only supports:
- `IncrementalDownload`: Download complete models to a local directory
- `GetObjectURL`: Get the ModelScope URL for a model

Other operations like upload, delete, list, etc., are not implemented as they are not commonly needed for model downloading scenarios.