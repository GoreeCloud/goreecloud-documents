// Package document defines GoreeCloud Documents' core domain model.
package document

import (
	"fmt"
	"strings"
	"time"
)

// ProcessingState records the durable document-processing lifecycle.
type ProcessingState string

const (
	StatePending    ProcessingState = "pending"
	StateProcessing ProcessingState = "processing"
	StateReady      ProcessingState = "ready"
	StateFailed     ProcessingState = "failed"
)

// Record is the stable metadata authority for one ingested document. Original
// file bytes are stored separately and referenced through StorageKey.
type Record struct {
	ID              string          `json:"id"`
	OwnerID         string          `json:"owner_id"`
	Title           string          `json:"title"`
	OriginalName    string          `json:"original_name"`
	MediaType       string          `json:"media_type"`
	StorageKey      string          `json:"storage_key"`
	ChecksumSHA256  string          `json:"checksum_sha256"`
	SizeBytes       int64           `json:"size_bytes"`
	ProcessingState ProcessingState `json:"processing_state"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
}

// Validate applies minimum domain invariants before persistence.
func (r Record) Validate() error {
	if !canonicalID(r.ID) || !canonicalID(r.OwnerID) {
		return fmt.Errorf("document and owner IDs must be canonical lowercase UUIDs")
	}
	if strings.TrimSpace(r.Title) == "" {
		return fmt.Errorf("document title must not be empty")
	}
	if strings.TrimSpace(r.OriginalName) == "" {
		return fmt.Errorf("original filename must not be empty")
	}
	if strings.TrimSpace(r.MediaType) == "" {
		return fmt.Errorf("media type must not be empty")
	}
	if strings.TrimSpace(r.StorageKey) == "" {
		return fmt.Errorf("storage key must not be empty")
	}
	if len(r.ChecksumSHA256) != 64 {
		return fmt.Errorf("SHA-256 checksum must contain 64 hexadecimal characters")
	}
	for _, ch := range r.ChecksumSHA256 {
		if !((ch >= '0' && ch <= '9') || (ch >= 'a' && ch <= 'f')) {
			return fmt.Errorf("SHA-256 checksum must be lowercase hexadecimal")
		}
	}
	if r.SizeBytes < 0 {
		return fmt.Errorf("document size must not be negative")
	}
	switch r.ProcessingState {
	case StatePending, StateProcessing, StateReady, StateFailed:
	default:
		return fmt.Errorf("unsupported processing state %q", r.ProcessingState)
	}
	if r.CreatedAt.IsZero() || r.UpdatedAt.IsZero() || r.UpdatedAt.Before(r.CreatedAt) {
		return fmt.Errorf("document timestamps are invalid")
	}
	return nil
}

func canonicalID(value string) bool {
	if len(value) != 36 {
		return false
	}
	for i, r := range value {
		switch i {
		case 8, 13, 18, 23:
			if r != '-' {
				return false
			}
		default:
			if !((r >= '0' && r <= '9') || (r >= 'a' && r <= 'f')) {
				return false
			}
		}
	}
	return true
}
