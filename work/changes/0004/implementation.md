# Implementation Notes

## Change

`0004` SolidJS frontend for the desktop shell

## Intended Path

- replace the static frontend with a SolidJS app
- keep the Wails bridge and workspace initializer unchanged
- keep the UI minimal and initialization-focused
- build the frontend with Vite into `frontend/dist`

## Validation

- frontend build
- `go test ./...`
- `go build -tags wails .`
