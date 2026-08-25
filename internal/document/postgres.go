package document

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

// PostgreSQLRepository persists GoreeCloud Documents metadata in the authoritative
// relational schema. Driver selection remains a runtime concern.
type PostgreSQLRepository struct {
	db *sql.DB
}

func NewPostgreSQLRepository(db *sql.DB) (*PostgreSQLRepository, error) {
	if db == nil {
		return nil, fmt.Errorf("database must not be nil")
	}
	return &PostgreSQLRepository{db: db}, nil
}

func (r *PostgreSQLRepository) Create(ctx context.Context, record Record) error {
	if err := record.Validate(); err != nil {
		return err
	}
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO documents (
			id, owner_id, title, original_name, media_type, storage_key,
			checksum_sha256, size_bytes, processing_state, created_at, updated_at
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`,
		record.ID, record.OwnerID, record.Title, record.OriginalName, record.MediaType,
		record.StorageKey, record.ChecksumSHA256, record.SizeBytes, record.ProcessingState,
		record.CreatedAt, record.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("create document: %w", err)
	}
	return nil
}

func (r *PostgreSQLRepository) Get(ctx context.Context, ownerID, documentID string) (Record, error) {
	var record Record
	err := r.db.QueryRowContext(ctx, `
		SELECT id, owner_id, title, original_name, media_type, storage_key,
		       checksum_sha256, size_bytes, processing_state, created_at, updated_at
		FROM documents
		WHERE owner_id = $1 AND id = $2`, ownerID, documentID).Scan(
		&record.ID, &record.OwnerID, &record.Title, &record.OriginalName, &record.MediaType,
		&record.StorageKey, &record.ChecksumSHA256, &record.SizeBytes, &record.ProcessingState,
		&record.CreatedAt, &record.UpdatedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return Record{}, ErrNotFound
	}
	if err != nil {
		return Record{}, fmt.Errorf("get document: %w", err)
	}
	return record, nil
}
