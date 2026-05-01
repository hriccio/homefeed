# Request Slice Map

- Change: `0001`
- Request: `work/changes/0001/request.md`
- Status: ready-for-build

## Request Boundary

Create the smallest Homefeed foundation that proves the workspace and database
can be initialized locally without polluting visible folders or expanding into
import, search, AI, or the full UI.

## Slice Mapping

| Slice | Status | Request coverage | Acceptance evidence |
| --- | --- | --- | --- |
| `docs/slices/0001-workspace-and-database.md` | planned | workspace creation, hidden metadata root creation, SQLite bootstrap, default feed records, and idempotent re-initialization | deterministic tests for path conventions and idempotency, plus a local scenario or command that proves initialization |

## Out Of Scope

- file or folder import
- search indexing or query UX
- AI suggestion or mutation flows
- full feed browsing UI
- visible sidecar metadata beside feed folders
- cloud sync or collaboration

## Open Questions

- Should the first executable surface be a Wails app command, a Go CLI scenario, or both?
- Should migrations be embedded Go files or SQL files under `migrations/`?
- Should the default workspace path be configurable before the settings UI exists?

## EGD Notes

Request-level expectation-gap detection should verify that the built slice proves
the foundation only, not the later product surface. It should check for
idempotent workspace initialization, hidden metadata placement, presence of the
default feed records, and absence of visible sidecar metadata in the feed
folders.

