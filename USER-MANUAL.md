# GoreeCloud Documents — Development User Manual

> GoreeCloud Documents is currently Development software. This manual covers the source-level development surface only; it is not a production deployment guide.

## Requirements

- Go toolchain compatible with the repository `go.mod`.
- A supported development host.

Runtime PostgreSQL wiring and production deployment configuration are not yet established by the current application entry point.

## Validate the source

```bash
gofmt -w ./cmd ./internal
go test ./...
go vet ./...
go build ./cmd/goreecloud-documents
```

CI is the authoritative automated validation for a pull-request candidate; local success alone is not acceptance evidence.

## Run the Development service

```bash
GOREECLOUD_DOCUMENTS_LISTEN=127.0.0.1:8790 go run ./cmd/goreecloud-documents
```

The documented Development service binds to loopback by default.

Available bounded diagnostics include:

- `GET /healthz`
- `GET /api/v1/status`

Do not expose this Development service publicly as a substitute for production authentication, authorization, TLS/publication, abuse controls, monitoring, privacy/security acceptance, backup/recovery, or deployment review.

## Current limitations

The current application does not provide an accepted end-user Glaze UI, production login/Identity integration, production document upload API, executable OCR/indexing pipeline, full-text search, AI assistant, production sharing, or Stable deployment.
