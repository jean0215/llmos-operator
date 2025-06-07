package v1

import (
	"github.com/llmos-ai/llmos-operator/pkg/apis/common"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=ad;ads
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="Registry",type="string",JSONPath=`.spec.registry`

// ApplicationData is a definition for the application data
type ApplicationData struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApplicationDataSpec   `json:"spec,omitempty"`
	Status ApplicationDataStatus `json:"status,omitempty"`
}

type ApplicationDataSpec struct {
	Registry    string       `json:"registry"`
	SourceFiles []SourceFile `json:"files,omitempty"`
}

type ApplicationDataStatus struct {
	RootPath          string             `json:"rootPath"`
	PreprocessedFiles []PreprocessedFile `json:"preprocessedFiles,omitempty"`
	Conditions        []common.Condition `json:"conditions,omitempty"`
}

type SourceFile struct {
	Path string `json:"path"`
	Etag string `json:"etag"`
}

type PreprocessedFile struct {
	UID           string           `json:"uid"`
	SourceFile    FileInfo         `json:"sourceFile"`
	ProcessedFile FileInfo         `json:"processedFile"`
	Condition     common.Condition `json:"condition,omitzero"`
}
