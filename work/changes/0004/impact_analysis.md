# Impact Analysis

## Change

`0004` SolidJS frontend for the desktop shell

## Areas Affected

- frontend build pipeline
- desktop UI runtime
- Wails asset bundle
- bridge invocation path

## Boundary Pressure

The main tension is between introducing the actual selected frontend runtime and
keeping the UI minimal enough that it does not turn into a product redesign.

Specific concerns:

- the current shell works with plain HTML/JavaScript
- the architecture explicitly calls for SolidJS
- the existing bridge and workspace behavior must not change

## Decisions Carried Into Build

- use a SolidJS frontend for the desktop shell
- keep the initialization bridge intact
- keep the frontend focused on proving initialization rather than adding
  product breadth

## Risks

- the frontend build could add more complexity than the slice needs
- generated assets could drift from source if the build step is not captured
  cleanly

## Build Guidance

The build should produce a minimal SolidJS app that exercises the existing
workspace initializer and can be rebuilt deterministically.

