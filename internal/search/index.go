// Package search defines GoreeCloud Documents' search-engine abstraction.
package search

import "context"

// Entry is the authorization-scoped representation sent to an approved search
// implementation. The native application remains authoritative for access.
type Entry struct {
	DocumentID string
	OwnerID    string
	Title      string
	Text       string
	Tags       []string
}

// Query captures the first native search contract.
type Query struct {
	RequesterID string
	Text        string
	Limit       int
}

// Hit is one ranked search result.
type Hit struct {
	DocumentID string  `json:"document_id"`
	Score      float64 `json:"score"`
	Highlight  string  `json:"highlight,omitempty"`
}

// Index is implemented by approved full-text or semantic search supporting
// components without delegating authorization decisions to those components.
type Index interface {
	Upsert(context.Context, Entry) error
	Delete(context.Context, string) error
	Search(context.Context, Query) ([]Hit, error)
}
