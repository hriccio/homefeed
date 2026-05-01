# Slice 0003 - Minimal Wails Desktop Shell

## Status

Ready for build.

## Pack

- `polyglot_client_server`

## Runtime Targets

- `client/typescript`
- `server/go`

## Architecture Mode

Local-first desktop app with a Wails shell, a SolidJS frontend, and a Go
backend bridge.

## Discovery Scope

This slice only covers the first desktop shell surface and the bridge to the
existing workspace initializer.

## Intent

Expose the proven workspace initializer through a minimal Wails desktop shell
so the repository begins using its selected runtime shape.

The shell should not add product breadth beyond initialization proof.

## Use-Case Contract

- `InitializeWorkspace`

The shell should invoke the existing initializer through a thin Go bridge and
surface the result in a minimal desktop context.

## Main Business Rules

- the workspace bootstrap behavior must remain unchanged
- hidden metadata stays under `~/Homefeed/.homefeed`
- visible feed folders stay clean
- import, search, and AI remain out of scope
- the shell should be minimal and not pretend to be the full product UI

## In Scope

- create the first Wails app entrypoint
- connect the frontend to the existing initialization use case through a thin
  bridge
- provide a minimal desktop surface that can prove the app starts and can
  initialize the workspace
- add tests or scenario checks around the bridge and the preserved workspace
  behavior

## Required Ports

- Wails runtime bridge
- filesystem access
- SQLite persistence
- existing migration loader

## Initial Test Plan

- tests for the bridge or command used to call `InitializeWorkspace`
- smoke validation that the desktop entrypoint starts
- workspace initialization tests continue to pass unchanged

## Scenario Definition

1. launch the desktop shell
2. trigger workspace initialization through the shell or its startup flow
3. confirm the workspace and database are created exactly as before
4. confirm a repeated launch does not break idempotency

## Out Of Scope

- import flows
- search UI
- AI assistance
- feed management polish
- broader frontend navigation

## Done Criteria

- the repository has a working Wails desktop shell entrypoint
- the existing initializer is reachable through the shell
- the workspace foundation still behaves as proven in earlier slices
- the shell remains minimal and does not absorb later product work

