# GoreeCloud Documents

GoreeCloud Documents is a privacy-first, self-hosted intelligent document-management platform for capture, OCR, search, structured organization, automation, collaboration, archival preservation, and long-term recovery.

The product is being built natively from the ground up as original GoreeCloud-owned software. Paperless-ngx and other document-management systems may be used as migration, interoperability, research, or behavioral references, but they are not the primary application codebase.

## Product workflow

Capture → OCR → Extract → Classify → Organize → Search → Automate → Collaborate → Share → Integrate

## Current development state

Lifecycle: **Development — Milestone 1 native foundation**

The current source foundation includes:

- Native Go service and loopback-safe development entry point.
- `GET /healthz` and `GET /api/v1/status` development endpoints.
- First-party document domain model and validation invariants.
- Durable processing-job contracts for ingestion, OCR, indexing, and extraction.
- PostgreSQL foundation schema for documents, pages, and processing jobs.
- OCR provider abstraction so OCR engines remain supporting components rather than product authority.
- Search-index abstraction that keeps authorization decisions inside GoreeCloud Documents.
- Privacy-conscious HTTP response headers and no-store API responses.
- Automated formatting, test, vet, and build validation through GitHub Actions.

This checkpoint does **not** establish production authentication, persistent runtime database wiring, file ingestion, OCR execution, full-text indexing, AI classification, production deployment, Stable qualification, or replacement of any existing document-management service.

## Architecture direction

GoreeCloud Documents owns document lifecycle, metadata, processing state, authorization, organization, search policy, workflow behavior, collaboration, and preservation semantics. Supporting libraries and engines may provide OCR, PDF parsing, indexing, cryptography, database connectivity, and standards compatibility through explicit adapter boundaries.

The application integrates with GoreeCloud platform systems rather than treating them as decorative labels:

- **Glaze UI** — adaptive visual and interaction language.
- **Wardveil Security** — protection state, security contracts, and evidence-backed security experiences.
- **Privacy Shield** — privacy controls, minimization, and privacy-state contracts.
- **Everkeep** — resilience, recovery, preservation, portability, succession, and digital legacy.
- **GoreeCloud Mesh** — application/service coordination and governance plane.

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

The development service intentionally binds to loopback by default. Public or production publication requires separate authentication, authorization, deployment, security, privacy, recovery, observability, and exact-release acceptance work.

## License

AGPL-3.0-only.
