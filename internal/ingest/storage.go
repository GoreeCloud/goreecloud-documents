package ingest

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// OriginalStore preserves immutable source bytes independently of OCR and index output.
type OriginalStore interface {
	Put(context.Context, string, io.Reader, int64) (size int64, checksum string, err error)
	Open(context.Context, string) (io.ReadCloser, error)
}

// LocalOriginalStore is a private development filesystem adapter.
type LocalOriginalStore struct { root string }

func NewLocalOriginalStore(root string) (*LocalOriginalStore, error) {
	if root == "" { return nil, fmt.Errorf("original store root is required") }
	abs, err := filepath.Abs(root); if err != nil { return nil, err }
	return &LocalOriginalStore{root: filepath.Clean(abs)}, nil
}

func (s *LocalOriginalStore) Put(_ context.Context, key string, src io.Reader, maxBytes int64) (int64, string, error) {
	if src == nil || maxBytes <= 0 { return 0, "", fmt.Errorf("source and positive size limit are required") }
	path, err := s.path(key); if err != nil { return 0, "", err }
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil { return 0, "", err }
	if _, err := os.Stat(path); err == nil { return 0, "", fmt.Errorf("original already exists") } else if !os.IsNotExist(err) { return 0, "", err }
	tmp, err := os.CreateTemp(filepath.Dir(path), ".original-*"); if err != nil { return 0, "", err }
	name := tmp.Name(); keep := true
	defer func(){ _ = tmp.Close(); if keep { _ = os.Remove(name) } }()
	_ = tmp.Chmod(0o600)
	h := sha256.New(); n, err := io.Copy(io.MultiWriter(tmp, h), io.LimitReader(src, maxBytes+1))
	if err != nil { return 0, "", err }
	if n > maxBytes { return 0, "", fmt.Errorf("original exceeds configured maximum") }
	if err := tmp.Sync(); err != nil { return 0, "", err }; if err := tmp.Close(); err != nil { return 0, "", err }
	if err := os.Link(name, path); err != nil { return 0, "", fmt.Errorf("publish original: %w", err) }
	if err := os.Remove(name); err != nil { return 0, "", err }; keep = false
	return n, hex.EncodeToString(h.Sum(nil)), nil
}

func (s *LocalOriginalStore) Open(_ context.Context, key string) (io.ReadCloser, error) {
	path, err := s.path(key); if err != nil { return nil, err }
	return os.Open(path)
}

func (s *LocalOriginalStore) path(key string) (string, error) {
	if len(key) != 36 { return "", fmt.Errorf("storage key must be a canonical UUID") }
	for i, r := range key { if i==8 || i==13 || i==18 || i==23 { if r!='-' { return "", fmt.Errorf("invalid storage key") }; continue }; if !((r>='0'&&r<='9')||(r>='a'&&r<='f')) { return "", fmt.Errorf("invalid storage key") } }
	return filepath.Join(s.root, key), nil
}
