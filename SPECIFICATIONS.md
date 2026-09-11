# GoreeCloud Documents — Specifications

## Authority

The controlling product record is **Project Specification — Documents** in the canonical GoreeCloud Google Drive hierarchy. This repository document is a source-side summary and must not replace that authoritative record.

## Product identity

- Product: GoreeCloud Documents
- Repository: `GoreeCloud/goreecloud-documents`
- Lifecycle: Development
- Implementation model: original GoreeCloud-owned native application
- Design language: Glaze UI
- Security: Wardveil Security
- Privacy: Privacy Shield
- Resilience/preservation: Everkeep
- Coordination: GoreeCloud Mesh

## Product purpose

GoreeCloud Documents is the native privacy-first, self-hosted intelligent document-management platform for document capture, OCR, extraction, organization, search, automation, collaboration, controlled sharing, AI-assisted understanding, archival preservation, portability, and recovery.

The core product workflow is:

`Capture → OCR → Extract → Classify → Organize → Search → Automate → Collaborate → Share → Integrate`

## Current verified source boundary

Current `main` contains a native Go development service, first-party document and processing contracts, PostgreSQL schema and repository/queue adapters, immutable local original-file storage primitives, ingestion orchestration, bounded processing-queue semantics, worker orchestration, development health/status HTTP surfaces, tests, and CI.

Runtime PostgreSQL driver/configuration wiring, executable OCR and indexing workers, production authentication/authorization, full-text search, AI capabilities, Glaze UI application surfaces, substantive Integral Platform System runtime adapters, deployment acceptance, Release Candidate, Stable, and production acceptance remain outstanding.

## Integral Platform Systems

All seven GoreeCloud Integral Platform Systems must be evaluated independently: GoreeCloud Manager, Privacy Shield, Wardveil Security, Everkeep, Glaze UI, GoreeCloud Mesh, and GoreeCloud Identity. Repository declarations or passing schema validation do not prove runtime acceptance.

## Implementation-state rule

Planned capabilities must not be represented as implemented, secure, accepted, production-ready, or Stable until the exact source, tests, runtime evidence, platform-system acceptance, and release evidence support that claim.
