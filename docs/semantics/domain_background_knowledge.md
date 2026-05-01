# Domain Background Knowledge

## Product Context

Homefeed is a local-first social file manager for Linux desktop.

The user has many files and does not navigate well with traditional folder
browsers. The desired interaction model borrows from social media feeds because
feeds make browsing, filtering, posting, commenting, following context, and
returning over time feel more natural.

The product is not a file explorer replacement. It is a social layer over the
local filesystem.

## Platform Context

Initial target:

- Linux Mint 21.3
- desktop app
- Wails + SolidJS + Go + SQLite
- root workspace at `~/Homefeed`

## Filesystem Context

The app should create and manage:

```text
~/Homefeed/
  professional/
  family/
  projects/
  personal/
  archive/
  .homefeed/
    index.sqlite
    meta/
    profiles/
    imports/
    agents/
    logs/
    cache/
```

The visible feed folders should remain clean. Metadata belongs under the hidden
`.homefeed` folder.

## Social Model Context

Folders become feeds. Files and non-file artifacts become posts. Owners become
profiles. Comments become local memory and interpretation rather than networked
social comments.

The unified feed is the main surface. Filters should include feed, owner,
profile, tag, file type, date, import batch, and AI classification.

## Import Context

Importing existing files or folders is required. The safest first strategy is to
copy into the workspace while preserving folder structure. Move and reference
strategies can come later.

Imports should produce file references, posts, and import batch records.

## Search Context

MVP search should cover title, body, comments, tags, owner/profile, feed,
filename, path, MIME type, and import source when relevant.

SQLite FTS5 is the likely search foundation, but implementation should verify
availability in the chosen SQLite setup.

## AI Context

AI support is required as an integration boundary, not as a hard provider
dependency.

Possible AI operations:

- summarize a file
- suggest title, feed, owner, and tags
- detect duplicates
- build timelines
- answer questions over indexed content
- identify stale or orphan metadata
- propose cleanup plans

AI suggestions should remain separate from user-confirmed metadata.

## Evaluation Risks

Watch for these expectation gaps:

- accidentally locking files into an app-only data model
- polluting visible folders with metadata files
- building a generic file explorer instead of a feed-first experience
- letting AI mutate filesystem organization without confirmation
- treating raw files as search results without the post context
- making the first slice too large by including full import, search, and AI behavior at once
