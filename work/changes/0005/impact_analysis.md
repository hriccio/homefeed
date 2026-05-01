# Impact Analysis

## Change

`0005` Local folder import workflow

## Likely Impacted Areas

- `internal/imports/`
- `internal/files/`
- `internal/posts/`
- `internal/workspace/`
- `migrations/`
- `internal/desktop/`
- `frontend/`
- `tests/`

## Tensions To Watch

- the import workflow must preserve ordinary filesystem access
- visible feed folders must stay free of sidecar metadata files
- the first import strategy should stay copy-based, not move-based
- the import feature should not drag in search or AI behavior
- the slice should stay small enough to prove deterministically

## Expected Behavioral Boundaries

- source content is copied into the managed workspace
- import metadata is written under `~/Homefeed/.homefeed`
- the workspace layout remains stable for the existing initialization flow
- the visible feed folders still look like normal folders to shell tools

## Verification Pressure

- deterministic path and copy behavior
- visible-folder cleanliness
- repeatable local scenario
- bridge or CLI exposure for the import use case

## Notes

The existing startup shell and workspace bootstrap should remain unchanged.
This slice should layer import behavior on top of the current foundation rather
than revisiting the initialization contract.
