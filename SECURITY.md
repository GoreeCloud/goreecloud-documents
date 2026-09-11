# GoreeCloud Documents — Security

## Status

GoreeCloud Documents is Development software and is not approved for public production deployment.

## Security boundary

Documents may contain highly sensitive content. Security design must protect original files, extracted text, metadata, indexes, processing state, API operations, collaboration/sharing state, credentials, and recovery material according to least privilege.

Current source uses a loopback-oriented Development service and does not claim production authentication or authorization. That boundary must not be weakened to make an unfinished service publicly reachable.

## Required controls before production acceptance

- GoreeCloud Identity-backed authentication with application-owned authorization decisions.
- Wardveil Security integration and evidence-backed protection state.
- Privacy Shield authorization and privacy-policy enforcement for document processing and sharing.
- Protected credential/secret handling outside source control and ordinary status output.
- Explicit authorization for original files, derived content, search results, AI operations, bulk actions, and sharing.
- Audit-relevant events without leaking document contents or secrets.
- Dependency, parser, upload, PDF/OCR, database, and API abuse review.
- Backup/restore, integrity, rollback, incident, and recovery validation.
- Production publication, TLS/proxy, monitoring, alerting, resource-bound, and failure-mode acceptance.

## Reporting

Security defects should be handled through authorized GoreeCloud security processes and project records. Do not place credentials, private documents, exploit secrets, or production-sensitive evidence in public issues or source files.

Passing CI or the Platform Contract validator does not establish security acceptance or Stable qualification.
