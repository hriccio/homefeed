# Request: Local Folder Import Workflow

## Intent

Add the first Homefeed import workflow so the desktop app can copy a selected
local folder into the managed workspace, preserve the folder structure, and
record the import in hidden metadata instead of polluting visible feed folders.

This request focuses on the first safe import strategy: copy into the
workspace. It does not ask for move-based import, search, or AI behavior.

## Source Evidence

- `docs/semantics/model_hypothesis.md`
- `docs/semantics/domain_background_knowledge.md`
- `architecture.md`
- `groundrules.md`
- `work/changes/0004/feedback.md`
- `work/changes/0004/release_decision.md`

## Acceptance Intent

The request should be considered satisfied when:

- a local source folder can be imported into a selected feed by copying it into
  `~/Homefeed`
- the import creates explicit hidden metadata under `~/Homefeed/.homefeed`
- visible feed folders remain free of sidecar metadata files
- the workflow is deterministic enough to prove with a repeatable local
  scenario
- the next slice still leaves search and AI out of scope
