# Request To Slice Map

## Change

`0005` Local folder import workflow

## Mapping

| Slice | Status | Coverage | Acceptance Evidence |
| --- | --- | --- | --- |
| `docs/slices/0005-local-folder-import.md` | built | the first copy-import workflow for a selected local folder, including import batch metadata, copied workspace files, and hidden metadata placement | a deterministic local scenario that imports a sample folder and proves the copy, metadata, and clean visible folder rules |

## Out Of Scope

- search indexing or search UI
- AI suggestion or mutation flows
- move-based import
- remote sync
- broad feed browsing polish

## Model Pressure

The current model wants Homefeed to keep ordinary filesystem usability intact
while adding an explicit social layer and import history. This slice should
preserve that boundary by copying data into the workspace and recording the
import under hidden metadata rather than treating imported files as ephemeral
UI-only state.
