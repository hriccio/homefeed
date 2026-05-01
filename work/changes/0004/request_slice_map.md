# Request Slice Map

- Change: `0004`
- Request: `work/changes/0004/request.md`
- Status: ready-for-build

## Request Boundary

Replace the static frontend assets with a SolidJS frontend for the desktop
shell, while preserving the existing Wails bridge and workspace initialization
behavior.

## Slice Mapping

| Slice | Status | Request coverage | Acceptance evidence |
| --- | --- | --- | --- |
| `docs/slices/0004-solidjs-frontend.md` | planned | frontend source scaffold, build pipeline, and a minimal SolidJS UI that calls `InitializeWorkspace` | frontend build output, preserved workspace tests, and a deterministic bridge or UI smoke check |

## Out Of Scope

- import flows
- search UI
- AI assistance
- workspace layout changes
- desktop shell redesign
- richer navigation or product polish

## Open Questions

- Should the SolidJS app be a single-screen initialization view or a tiny
  status-driven shell?
- Should the existing desktop entrypoint continue to trigger initialization on
  startup, or should the button remain the only invocation path?

## EGD Notes

Request-level review should confirm that the frontend runtime shape now matches
SolidJS while the workspace behavior remains unchanged.

