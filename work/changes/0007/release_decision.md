# Release Decision

## Change

`0007` Create note post workflow

## Decision

Accepted as the intended internal version.

## Basis

- the request asked for the first note-post workflow so the desktop app can
  create a note inside a selected feed as a normal file while recording the
  post in hidden workspace metadata.
- the implemented slice now exposes a `CreateNotePost` bridge in the desktop
  shell.
- the SolidJS frontend includes a minimal note-post form.
- the note workflow creates a normal markdown file inside the selected feed.
- the note workflow records a durable post row in SQLite under the hidden
  workspace database.
- visible feed folders remain free of sidecar metadata files.
- `go test ./...` passes.
- `npm run build` in `frontend/` passes.
- `go build -tags wails,production -o /tmp/homefeed-wails .` passes.
- the EGD result judged the request satisfied.

## Residual Risk

- No manual GUI click-through of the note form was run in this environment, so
  the exact interactive path remains unverified here.

## Notes

The release boundary is the request, not just the note-post slice. The request
is considered satisfied with the current internal state.
