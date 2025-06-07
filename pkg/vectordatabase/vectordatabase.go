package vectordatabase

import "context"

type Client interface {
	CreateCollection(ctx context.Context, collectionName string) error
	CollectionExists(ctx context.Context, collectionName string) (bool, error)
	DeleteCollection(ctx context.Context, collectionName string) error
	InsertObjects(ctx context.Context, collectionName string, objects []Document) error
	ListObjects(ctx context.Context, collectionName, uid string, offset, limit int) (*ObjectList, error)
	DeleteObjects(ctx context.Context, collectionName string, uid string) (*DeleteResult, error)
	Search(ctx context.Context, collectionName, keyword string, threshold float64, limit int) (*SearchResults, error)
}

// BaseDocument represents the core document properties
type BaseDocument struct {
	UID       string `json:"uid"`       // Document UID
	Document  string `json:"document"`  // Document name
	Index     int    `json:"index"`     // Chunk index
	Keywords  string `json:"keywords"`  // Keywords
	Content   string `json:"content"`   // Chunk content
	Timestamp string `json:"timestamp"` // Creation time
}

// Document represents a document to be inserted
type Document struct {
	BaseDocument
}

// SearchResult represents a single query result
type SearchResult struct {
	BaseDocument
	ID       string    `json:"id"`       // Object ID
	Distance float64   `json:"distance"` // Vector distance
	Vector   []float32 `json:"vector"`   // Vector representation
}

// SearchResults represents the complete query results
type SearchResults struct {
	Results []SearchResult
	Total   int
}

// ObjectInfo represents information about an object in the collection
type ObjectInfo struct {
	BaseDocument
	ID     string    `json:"id"`     // Object ID
	Vector []float32 `json:"vector"` // Vector representation
}

// ObjectList represents a list of objects with pagination support
type ObjectList struct {
	Objects []ObjectInfo `json:"objects"` // List of objects
	Total   int          `json:"total"`   // Total count
	Offset  int          `json:"offset"`  // Current offset
	Limit   int          `json:"limit"`   // Current limit
}

// DeleteResult represents the result of a delete operation
type DeleteResult struct {
	DeletedIDs []string `json:"deleted_ids"` // IDs of deleted objects
	Total      int      `json:"total"`       // Total number of deleted objects
}
