package ingest

import (
	"context"
	"crypto/rand"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/GoreeCloud/goreecloud-documents/internal/document"
)

type Service struct {
	repository document.Repository
	originals  OriginalStore
	maxBytes   int64
	now        func() time.Time
}

func NewService(repository document.Repository, originals OriginalStore, maxBytes int64) Service {
	return Service{repository: repository, originals: originals, maxBytes: maxBytes, now: time.Now}
}

// Ingest preserves the original bytes first, then records authoritative metadata.
// OCR and search indexing remain later processing stages and never replace the original.
func (s Service) Ingest(ctx context.Context, ownerID, title, originalName, mediaType string, src io.Reader) (document.Record, error) {
	if s.repository == nil || s.originals == nil || s.maxBytes <= 0 {
		return document.Record{}, fmt.Errorf("ingestion service unavailable")
	}
	if strings.TrimSpace(title) == "" || strings.TrimSpace(originalName) == "" || strings.TrimSpace(mediaType) == "" {
		return document.Record{}, fmt.Errorf("document metadata is incomplete")
	}
	id, err := newUUID()
	if err != nil {
		return document.Record{}, err
	}
	size, checksum, err := s.originals.Put(ctx, id, src, s.maxBytes)
	if err != nil {
		return document.Record{}, fmt.Errorf("preserve original: %w", err)
	}
	now := s.now().UTC()
	record := document.Record{
		ID:              id,
		OwnerID:         ownerID,
		Title:           title,
		OriginalName:    originalName,
		MediaType:       mediaType,
		StorageKey:      id,
		ChecksumSHA256:  checksum,
		SizeBytes:       size,
		ProcessingState: document.StatePending,
		CreatedAt:       now,
		UpdatedAt:       now,
	}
	if err := s.repository.Create(ctx, record); err != nil {
		return document.Record{}, fmt.Errorf("persist document metadata: %w", err)
	}
	return record, nil
}

func newUUID() (string, error) {
	var b [16]byte
	if _, err := rand.Read(b[:]); err != nil {
		return "", fmt.Errorf("generate document ID: %w", err)
	}
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:16]), nil
}
