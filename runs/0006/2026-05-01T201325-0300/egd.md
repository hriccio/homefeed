# EGD Run

## Change

`0006` Import schema source of truth

## Timestamp

`2026-05-01T201325-0300`

## Validation Evidence

- `go test ./...`
- `npm run build` in `frontend/`
- `go build -tags wails,production -o /tmp/homefeed-wails .`

## Summary

The import schema is now single-sourced in the canonical migration artifact.
Import behavior remains unchanged and deterministic validation passes.
