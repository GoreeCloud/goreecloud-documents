package ingest

import (
	"context"
	"strings"
	"testing"

	"github.com/GoreeCloud/goreecloud-documents/internal/document"
)

func TestServiceIngestPreservesAndPersistsOriginal(t *testing.T) {
	repository := document.NewMemoryRepository()
	originals, err := NewLocalOriginalStore(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	service := NewService(repository, originals, 1024)
	ownerID := "11111111-1111-1111-1111-111111111111"
	record, err := service.Ingest(context.Background(), ownerID, "Receipt", "receipt.pdf", "application/pdf", strings.NewReader("source bytes"))
	if err != nil {
		t.Fatal(err)
	}
	if record.ProcessingState != document.StatePending || record.SizeBytes != 12 || len(record.ChecksumSHA256) != 64 {
		t.Fatalf("unexpected record: %#v", record)
	}
	persisted, err := repository.Get(context.Background(), ownerID, record.ID)
	if err != nil {
		t.Fatal(err)
	}
	if persisted.StorageKey != record.ID {
		t.Fatalf("storage key=%q", persisted.StorageKey)
	}
	reader, err := originals.Open(context.Background(), record.StorageKey)
	if err != nil {
		t.Fatal(err)
	}
	_ = reader.Close()
}
