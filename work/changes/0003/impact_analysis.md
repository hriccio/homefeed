# Impact Analysis

## Change

`0003` Minimal Wails desktop shell

## Areas Affected

- desktop entrypoint
- Go-to-frontend bridge
- client runtime scaffold
- build/run workflow

## Boundary Pressure

The main pressure is between introducing the real desktop runtime and keeping
the slice small enough that it does not become a UI product in disguise.

Specific concerns:

- the repository currently proves behavior with a Go CLI scenario
- the selected product shape expects a Wails + SolidJS shell
- the first shell should expose the existing initializer, not redefine the
  workspace model

## Decisions Carried Into Build

- preserve the current workspace bootstrap behavior exactly
- add the first desktop shell as a thin wrapper around the proven initializer
- keep import, search, AI, and rich feed UI out of scope

## Risks

- a Wails scaffold could expand into a larger UI implementation than the slice
  needs
- if startup initialization is chosen, the shell could hide an important action
  that ought to be explicit in the UI

## Build Guidance

The build should prioritize a minimal, testable desktop entrypoint that exposes
the existing initialization use case without changing the workspace semantics.

