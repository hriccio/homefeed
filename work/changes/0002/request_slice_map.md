# Request Slice Map

- Change: `0002`
- Request: `work/changes/0002/request.md`
- Status: ready-for-build

## Request Boundary

Make the workspace schema bootstrap single-sourced so the runtime behavior and
the repository migration artifact do not drift apart.

## Slice Mapping

| Slice | Status | Request coverage | Acceptance evidence |
| --- | --- | --- | --- |
| `docs/slices/0002-workspace-migration-source-of-truth.md` | planned | migration loading and schema execution for the workspace initializer, with no behavior change to the existing workspace foundation | deterministic tests showing workspace initialization still works and the migration definition is not duplicated as executable source |

## Out Of Scope

- workspace shape changes
- feed model changes
- import, search, AI, or UI work
- broader schema redesign

## Open Questions

- Should the migration runner execute the SQL artifact directly or via a small
  loader abstraction?
- Should future migrations be numbered SQL files under `migrations/` with a
  shared executor?

## EGD Notes

Request-level review should check that the runtime behavior stays identical for
the foundation slice while the schema definition moves to one authoritative
source.

