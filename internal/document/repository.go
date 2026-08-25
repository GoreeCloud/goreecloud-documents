package document

import (
	"context"
	"fmt"
	"sync"
)

// Repository is the metadata persistence boundary owned by GoreeCloud Documents.
type Repository interface {
	Create(context.Context, Record) error
	Get(context.Context, string, string) (Record, error)
}

var ErrNotFound = fmt.Errorf("document not found")

// MemoryRepository is a deterministic development implementation used until the
// PostgreSQL adapter is wired into the runtime.
type MemoryRepository struct {
	mu sync.RWMutex
	records map[string]Record
}

func NewMemoryRepository() *MemoryRepository { return &MemoryRepository{records: make(map[string]Record)} }

func (r *MemoryRepository) Create(_ context.Context, record Record) error {
	if err := record.Validate(); err != nil { return err }
	r.mu.Lock(); defer r.mu.Unlock()
	if _, exists := r.records[record.ID]; exists { return fmt.Errorf("document already exists") }
	r.records[record.ID] = record
	return nil
}

func (r *MemoryRepository) Get(_ context.Context, ownerID, documentID string) (Record, error) {
	r.mu.RLock(); defer r.mu.RUnlock()
	record, ok := r.records[documentID]
	if !ok || record.OwnerID != ownerID { return Record{}, ErrNotFound }
	return record, nil
}
