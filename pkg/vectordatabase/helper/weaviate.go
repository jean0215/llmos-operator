package helper

import (
	"strings"

	vd "github.com/llmos-ai/llmos-operator/pkg/vectordatabase"
	"github.com/llmos-ai/llmos-operator/pkg/vectordatabase/vectorizer"
	"github.com/llmos-ai/llmos-operator/pkg/vectordatabase/weaviate"
)

const (
	weaviateHost = "weaviate.llmos-agents.svc.cluster.local:80"
	httpScheme   = "http"
)

// webhook will prove the embedding model is valid
func newVectorizer(embeddingModel string) *vectorizer.CustomVectorizer {
	tmp := strings.Split(embeddingModel, "/")
	namespace, name := tmp[0], tmp[1]
	serviceName := "modelservice-" + name
	host := serviceName + "." + namespace + ".svc.cluster.local:8000"

	return vectorizer.NewCustomVectorizer(host, httpScheme)
}

func NewVectorDatabaseClient(embeddingModel string) (vd.Client, error) {
	vectorizer := newVectorizer(embeddingModel)
	return weaviate.NewClient(weaviateHost, "http", vectorizer)
}
