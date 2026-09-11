# GoreeCloud Documents

GoreeCloud Documents is a privacy-first, self-hosted intelligent document-management platform for capture, OCR, search, structured organization, automation, collaboration, archival preservation, and long-term recovery.

The product is being built natively from the ground up as original GoreeCloud-owned software. Paperless-ngx and other document-management systems may be used as migration, interoperability, research, or behavioral references, but they are not the primary application codebase.

## Product workflow

Capture → OCR → Extract → Classify → Organize → Search → Automate → Collaborate → Share → Integrate

## Current development state

Lifecycle: **Development / nonconformant**

The current source includes:

- Native Go service and loopback-safe development entry point.
- `GET /healthz` and `GET /api/v1/status` bounded Development diagnostics.
- First-party document and processing domain contracts.
- PostgreSQL foundation schema and PostgreSQL document-metadata repository adapter.
- Immutable original-file storage abstraction with a private local filesystem adapter and SHA-256 integrity recording.
- Native ingestion orchestration that preserves original bytes before authoritative metadata creation.
- OCR-provider and search-index adapter boundaries.
- Bounded processing queue with atomic single-worker claim, attempts, delayed retry, completion, and failure transitions.
- PostgreSQL-backed implementation of the processing queue contract.
- Transport-neutral bounded processing-worker orchestration.
- Privacy-conscious API response headers.
- Automated formatting, test, vet, build, and Platform Contract validation through GitHub Actions.

This checkpoint does **not** establish runtime PostgreSQL connection/driver wiring, crash-safe worker leasing, executable OCR or indexing workers, production authentication/authorization, full-text search, AI classification, Glaze UI application surfaces, accepted Integral Platform System runtime integration, production deployment, Release Candidate, Stable qualification, or replacement of an existing production document-management service.

## Architecture direction

GoreeCloud Documents owns document lifecycle, metadata, processing state, authorization, organization, search policy, workflow behavior, collaboration, and preservation semantics. Supporting libraries and engines may provide OCR, PDF parsing, indexing, cryptography, database connectivity, and standards compatibility through explicit adapter boundaries.

All seven Integral Platform Systems are evaluated independently:

- **GoreeCloud Manager** — application administration and management integration.
- **Privacy Shield** — privacy authorization, minimization, processing, retention, and sharing boundaries.
- **Wardveil Security** — security contracts and evidence-backed protection state.
- **Everkeep** — resilience, recovery, preservation, portability, succession, and digital legacy.
- **Glaze UI** — adaptive visual and interaction language; current required Stable consumer target is V1.3 / 1.3.0.
- **GoreeCloud Mesh** — application/service coordination and governed cross-product capability exchange.
- **GoreeCloud Identity** — authenticated identity and claims; Documents remains responsible for application authorization.

Planned first-party integrations include GoreeCloud AI, Search, Drive, Mail, Backup, Identity, Notify, Monitor, and Mesh through explicit, versioned, least-privilege interfaces.

## Milestones

1. **Native Foundation** — repository architecture, domain model, database schema, API, configuration, authentication boundary, platform contracts, tests, CI, documentation, and Glaze UI shell.
2. **Capture and Document Engine** — ingestion, preserved originals, pages, OCR abstraction, extracted text, PDF handling, and processing history.
3. **Search and Organization** — full-text indexing, filters, tags, correspondents, document types, custom fields, relationships, saved searches, and reusable views.
4. **Intelligence** — extraction assistance, AI classification, document assistant, grounded Q&A, semantic retrieval, and similar-document discovery.
5. **Automation and Communications** — ingestion locations, email import/rules, workflows, retry history, and notifications.
6. **Collaboration and Access** — multi-user permissions, groups, secure sharing, expiration, activity, and administration.
7. **Preservation and Advanced Operations** — versioning, PDF operations, PDF/A, archival integrity, portable export, recovery validation, and Everkeep integration.
8. **Platform Integration and Stabilization** — Mesh, first-party integrations, observability, security/privacy/accessibility validation, migration, recovery exercises, release engineering, and production-readiness gates.

## Development

```bash
go test ./...
go vet ./...
go build ./cmd/goreecloud-documents
GOREECLOUD_DOCUMENTS_LISTEN=127.0.0.1:8790 go run ./cmd/goreecloud-documents
```

The Development service intentionally binds to loopback by default. Public or production publication requires separate authentication, authorization, deployment, security, privacy, recovery, observability, and exact-release acceptance work.

See `SPECIFICATIONS.md`, `FEATURES.md`, `FEATURE-ROADMAP.md`, `SECURITY.md`, `PRIVACY POLICY.md`, and `goreecloud.platform.yaml` for the repository-side governed boundaries.

## License

AGPL-3.0-only.
