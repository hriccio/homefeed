# Request Slice Map

- Change: `0003`
- Request: `work/changes/0003/request.md`
- Status: ready-for-build

## Request Boundary

Add the first Wails desktop shell around the existing workspace initializer so
the repository begins to match its selected runtime shape without expanding into
import, search, AI, or a full product UI.

## Slice Mapping

| Slice | Status | Request coverage | Acceptance evidence |
| --- | --- | --- | --- |
| `docs/slices/0003-minimal-wails-shell.md` | planned | Wails app entrypoint, thin Go bridge to `InitializeWorkspace`, and a minimal frontend surface that can prove the desktop runtime starts and invokes initialization | desktop launch smoke test or documented scenario plus deterministic tests for the bridge and preserved workspace behavior |

## Out Of Scope

- import flows
- search UI
- AI assistance UI
- feed browsing polish
- metadata model changes
- workspace layout changes

## Open Questions

- Should initialization happen on app startup or through an explicit UI action?
- Should the first shell be a status page, a single action screen, or a hidden
  bridge-only desktop stub?

## EGD Notes

Request-level review should check that the desktop shell exists without
disturbing the already-proven workspace bootstrap behavior.

