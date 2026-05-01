# Request: Homefeed Foundation

## Intent

Establish the smallest useful Homefeed foundation for the initial product:

- a local-first desktop app for Linux
- a managed workspace rooted at `~/Homefeed`
- hidden operational metadata under `~/Homefeed/.homefeed`
- an initial database and default feed records
- a first executable surface that can prove initialization works

The request is to create the workspace and database baseline without yet
implementing import, search, AI behavior, or the full feed UI.

## Source Evidence

- `work/sources/initial_handoff/homefeed-codex-handoff.md`
- `readme.md`
- `docs/semantics/model_hypothesis.md`
- `docs/semantics/domain_background_knowledge.md`
- `docs/slices/0001-workspace-and-database.md`

## Acceptance Intent

The request should be considered satisfied when the repository can prove:

- `~/Homefeed` is created when missing
- `~/Homefeed/.homefeed` is created
- `~/Homefeed/.homefeed/index.sqlite` exists
- default feed records exist for:
  - `professional`
  - `family`
  - `projects`
  - `personal`
  - `archive`
- rerunning initialization is idempotent
- visible feed folders remain free of sidecar metadata files

