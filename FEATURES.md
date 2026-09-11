# GoreeCloud Documents — Features

This file distinguishes verified Development source from planned product capability. The canonical **Project Specification — Documents** remains authoritative.

## Present in current Development source

- Native Go service and loopback-oriented development runtime.
- First-party document and processing domain contracts.
- PostgreSQL foundation schema.
- Document repository abstraction and PostgreSQL metadata adapter.
- Immutable original-file storage abstraction with a private local filesystem adapter and SHA-256 integrity recording.
- Native ingestion orchestration that preserves original bytes before authoritative metadata creation.
- OCR-provider and search-index adapter boundaries.
- Bounded processing queue with claim, attempts, delayed retry, completion, and failure transitions.
- PostgreSQL-backed processing queue adapter using atomic eligible-job claim semantics.
- Transport-neutral bounded processing worker orchestration.
- Development health/status endpoints and privacy-conscious API response headers.
- Go formatting, test, vet, and build validation.

## Planned / not yet accepted

- Runtime PostgreSQL connection/driver configuration and crash-safe worker leasing/recovery.
- Executable OCR, extraction, and full-text indexing workers.
- 100+ OCR language support and production OCR acceptance.
- Full-text search, highlighting, relevance ranking, autocomplete, saved searches, and similar-document discovery.
- Tags, correspondents, document types, custom fields, relationships, and configurable filing/storage paths.
- AI classification, assistant, grounded document Q&A, summarization, semantic retrieval, and provenance controls.
- Workflow automation, email ingestion/rules, notifications, and processing-history administration.
- Multi-user access, GoreeCloud Identity, groups/roles, granular permissions, and secure expiring sharing.
- Document versioning, PDF editing, PDF/A archival workflows, portable export, and recovery acceptance.
- Glaze UI document inbox, library, viewer, search, metadata, workflow, sharing, administration, and settings surfaces.
- Substantive Manager, Privacy Shield, Wardveil Security, Everkeep, Mesh, and Identity runtime integration and acceptance.
- Production deployment, Release Candidate, Stable, and production acceptance.
