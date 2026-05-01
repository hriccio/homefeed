# Slice 0004 - SolidJS Frontend

## Status

Built.

## Pack

- `polyglot_client_server`

## Runtime Targets

- `client/typescript`
- `server/go`

## Architecture Mode

Local-first desktop app with a Wails shell, a SolidJS frontend, and a Go
backend bridge.

## Discovery Scope

This slice only covers the frontend runtime and build pipeline for the minimal
desktop shell.

## Intent

Replace the static frontend assets with a real SolidJS frontend while keeping
the existing Wails bridge and workspace bootstrap behavior unchanged.

## Use-Case Contract

- `InitializeWorkspace`

The frontend should call the existing Go bridge so the desktop shell can
initialize the workspace and render the result.

## Main Business Rules

- the workspace bootstrap behavior must remain unchanged
- the frontend stays minimal and purpose-built for initialization proof
- import, search, and AI remain out of scope
- the frontend build output must still be embeddable by the Wails entrypoint

## In Scope

- create the SolidJS frontend source scaffold
- add the frontend build configuration needed to produce embeddable assets
- render a minimal UI that can invoke `InitializeWorkspace`
- preserve the current workspace and database behavior
- add deterministic tests or smoke checks around the bridge and build output

## Required Ports

- Wails runtime bridge
- filesystem access
- SQLite persistence
- frontend build toolchain

## Initial Test Plan

- frontend build produces the embeddable assets
- Go tests for the bridge and workspace behavior continue to pass
- the Wails entrypoint still compiles with the `wails` tag

## Scenario Definition

1. build the SolidJS frontend
2. launch the desktop shell
3. invoke workspace initialization from the SolidJS UI
4. confirm the workspace and database are created exactly as before
5. confirm repeat invocation remains idempotent

## Out Of Scope

- import flows
- search UI
- AI assistance
- workspace layout changes
- broader desktop product UI

## Done Criteria

- the frontend is SolidJS rather than plain static HTML/JS
- the desktop shell still invokes the proven workspace initializer
- the workspace foundation behavior remains unchanged
- the frontend build output can be embedded by the Wails app

## Implementation Notes

- the frontend source lives under `frontend/src`
- the Vite build output is embedded from `frontend/dist`
- the default build remains testable through `go test ./...` and
  `go build -tags wails .`
