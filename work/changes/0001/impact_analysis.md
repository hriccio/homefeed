# Impact Analysis

## Change

`0001` Homefeed foundation

## Areas Affected

- workspace initialization
- SQLite schema/bootstrap
- default feed records
- filesystem cleanliness rules
- future scaffold placement for Wails + SolidJS + Go

## Boundary Pressure

The main pressure is between making the first slice small enough to build
deterministically and making it broad enough to establish the repository's core
filesystem promise.

Specific tensions:

- the app must create a managed workspace without polluting visible folders
- the database must exist early enough to support feeds, but not drag in import
  or search implementation
- the first executable surface must prove initialization without forcing the
  entire UI to be implemented

## Decisions Carried Into Build

- use `~/Homefeed` as the managed root
- place operational metadata under `~/Homefeed/.homefeed`
- create the five canonical feeds on initialization
- keep the slice idempotent
- keep import, search, and AI out of scope for this build increment

## Risks

- the initializer could accidentally create visible metadata files beside feed
  folders
- the initial scaffold could become too large if UI, import, and search are
  pulled in early
- the default feed names could drift if the source handoff is not treated as the
  authoritative product evidence

## Build Guidance

The build should privilege deterministic tests around path creation and
idempotent initialization before expanding into richer UI or indexing work.

