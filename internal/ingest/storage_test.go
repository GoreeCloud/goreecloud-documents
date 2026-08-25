package ingest

import (
	"context"
	"io"
	"strings"
	"testing"
)

func TestLocalOriginalStorePreservesOriginal(t *testing.T) {
	store, err := NewLocalOriginalStore(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	key := "11111111-1111-1111-1111-111111111111"
	size, checksum, err := store.Put(context.Background(), key, strings.NewReader("original bytes"), 64)
	if err != nil {
		t.Fatal(err)
	}
	if size != 14 || len(checksum) != 64 {
		t.Fatalf("size=%d checksum=%q", size, checksum)
	}
	r, err := store.Open(context.Background(), key)
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	body, _ := io.ReadAll(r)
	if string(body) != "original bytes" {
		t.Fatalf("body=%q", body)
	}
	if _, _, err := store.Put(context.Background(), key, strings.NewReader("replacement"), 64); err == nil {
		t.Fatal("expected immutable original rejection")
	}
}

func TestLocalOriginalStoreRejectsOversizedOriginal(t *testing.T) {
	store, _ := NewLocalOriginalStore(t.TempDir())
	if _, _, err := store.Put(context.Background(), "11111111-1111-1111-1111-111111111111", strings.NewReader("abcd"), 3); err == nil {
		t.Fatal("expected size error")
	}
}
