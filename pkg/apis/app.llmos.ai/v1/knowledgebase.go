package v1

import (
	"github.com/llmos-ai/llmos-operator/pkg/apis/common"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=kb;kbs
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// KnownledgeBase is a definition for the LLM KnownledgeBase
type KnownledgeBase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KnownledgeBaseSpec   `json:"spec,omitempty"`
	Status KnownledgeBaseStatus `json:"status,omitempty"`
}

type KnownledgeBaseSpec struct {
	// EmbeddingModel is from model service including namespace, such as defalt/text-embedding-3-small
	EmbeddingModel string          `json:"embeddingModel"`
	Files          []ImportingFile `json:"files,omitempty"`
}

type ImportingFile struct {
	// Category is the application data name
	Category string `json:"category"`
	UID      string `json:"uid"`
}

type KnownledgeBaseStatus struct {
	ParsedFiles []ParsedFile `json:"parsedFiles,omitempty"`
}

type ParsedFile struct {
	UID          string           `json:"uid"`
	File         FileInfo         `json:"file"`
	ImportedTime metav1.Time      `json:"importedTime,omitempty"`
	Condition    common.Condition `json:"condition,omitzero"`
}
