# Homefeed — Codex Project Handoff

## 1. Project Intent

Build a **local-first social file manager** for Linux desktop.

The user has many files on their computer and does not navigate well with traditional folder browsers. The desired interaction model is inspired by social media feeds, because that metaphor has become natural for browsing, filtering, posting, commenting, and following context over time.

The product is not merely a file explorer. It is a **social layer over the local filesystem**.

Core idea:

> Files are not only “inside folders”. Files are “posts in feeds”, with owners, profiles, comments, search, import history, and AI-assisted organization.

The app should preserve normal filesystem usability. It must not lock files into a proprietary database. The filesystem remains usable without the app.

---

## 2. Target Platform

Initial target:

- OS: **Linux Mint 21.3**
- App type: **Desktop app**
- Storage: **Local-first**
- Root data folder: user home directory

Default workspace path:

```txt
~/Homefeed
```

The app should create and manage this folder.

---

## 3. Chosen Stack

Recommended stack:

```txt
Wails + SolidJS + Go + SQLite
```

### Rationale

The app is primarily about:

- filesystem access
- file importing
- indexing
- search
- metadata management
- background jobs
- AI-agent orchestration

Go is a strong fit for this backend-heavy local application, and it aligns with the user’s preferred ecosystem.

### Alternative Considered

```txt
Tauri + SolidJS + Rust + SQLite
```

Tauri/Rust has stronger native security and desktop integration, but Rust adds cognitive overhead and ties the domain engine to Rust unless extra effort is spent decoupling it. Wails + Go is preferred because the core logic can later become a CLI, daemon, or standalone service more naturally.

---

## 4. Product Metaphor

### Feeds

A folder can behave as a feed.

Examples:

```txt
~/Homefeed/professional
~/Homefeed/family
~/Homefeed/projects
~/Homefeed/personal
~/Homefeed/archive
```

Each top-level folder is a feed by default.

The main screen is a unified feed containing posts from all feeds.

Users can filter by:

- feed
- owner
- profile
- tag
- file type
- date
- import batch
- AI classification

### Posts

A post can represent:

- a file
- a folder import
- a note
- a link
- an AI-generated summary
- a decision log
- a reference
- a comment thread around an artifact

Files are treated as posts, but not every post must have a file.

### Owners and Profiles

Posts have an owner.

The owner may be:

- the user
- another person
- a company
- a group
- a project
- an unknown source

When the user assigns an owner that does not exist, the system should create a profile automatically.

Examples:

```txt
Me
Family
Client X
Company Y
Project Living Circles
```

### Comments

Posts can have comments.

Comments are local metadata, not social-network comments. They are meant for context, memory, interpretation, and collaboration-like self-annotation.

---

## 5. Core Requirements

### Required

1. Create and manage a workspace under `~/Homefeed`.
2. Treat folders as feeds.
3. Show a main unified feed.
4. Allow filtering the main feed by feed/profile/tag/type/date.
5. Allow creating posts.
6. Allow posts to reference files.
7. Allow importing existing folders/files.
8. Preserve imported files as normal filesystem files.
9. Allow assigning owners to posts.
10. Auto-create profiles for unknown owners.
11. Allow comments on posts.
12. Provide search.
13. Provide an AI-agent integration boundary.
14. Avoid polluting normal folders with visible sidecar metadata files.

### Non-goals for MVP

- Multi-user cloud sync
- Real social networking
- Remote collaboration
- Mobile app
- Full permissions/ACL model
- Replacing the native file manager completely
- Complex file versioning

---

## 6. Filesystem Layout

Use a hidden metadata/system folder inside the workspace.

```txt
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

### Important Principle

Do not create visible `.meta.yaml` files beside each file.

The user prefers to avoid folder pollution. Metadata should live under:

```txt
~/Homefeed/.homefeed/meta
```

The visible feed folders should remain clean.

---

## 7. Metadata Strategy

SQLite is the primary operational index.

Markdown/YAML/JSON files under `.homefeed/` may be used for portability, debugability, and recovery.

Recommended approach:

- SQLite stores indexed entities and relationships.
- `.homefeed/meta/` stores canonical or exportable metadata records.
- Metadata records should use stable IDs rather than raw filenames only.
- File records should include path, content hash, size, mtime, and import batch.

Potential file metadata path:

```txt
~/Homefeed/.homefeed/meta/posts/<post_id>.json
~/Homefeed/.homefeed/meta/profiles/<profile_id>.json
~/Homefeed/.homefeed/meta/imports/<import_id>.json
```

Avoid metadata filenames that depend only on mutable file paths.

---

## 8. Data Model Draft

### Feed

```go
type Feed struct {
    ID          string
    Name        string
    Slug        string
    Path        string
    Description string
    CreatedAt   time.Time
    UpdatedAt   time.Time
}
```

### Profile

```go
type Profile struct {
    ID          string
    DisplayName string
    Slug        string
    Kind        string // person | group | company | project | system | unknown
    Notes       string
    CreatedAt   time.Time
    UpdatedAt   time.Time
}
```

### Post

```go
type Post struct {
    ID          string
    FeedID      string
    OwnerID     string
    Title       string
    Kind        string // file | note | link | folder_import | ai_summary | decision_log | reference
    Body        string
    FileID      *string
    Source      string
    CreatedAt   time.Time
    UpdatedAt   time.Time
}
```

### FileRef

```go
type FileRef struct {
    ID          string
    WorkspacePath string
    AbsolutePath  string
    OriginalPath  string
    MimeType      string
    SizeBytes     int64
    SHA256        string
    ModifiedAt    time.Time
    ImportedAt    time.Time
    ImportID      *string
}
```

### Comment

```go
type Comment struct {
    ID        string
    PostID    string
    AuthorID  string
    Body      string
    CreatedAt time.Time
    UpdatedAt time.Time
}
```

### Tag

```go
type Tag struct {
    ID   string
    Name string
    Slug string
}
```

### ImportBatch

```go
type ImportBatch struct {
    ID          string
    SourcePath  string
    TargetFeedID string
    Strategy    string // copy | move | reference
    Status      string // pending | running | completed | failed
    CreatedAt   time.Time
    CompletedAt *time.Time
}
```

---

## 9. Import Behavior

The import feature is required.

A user should be able to select an existing folder or files and import them into Homefeed.

### Import Strategies

Support at least one initially:

1. **Copy into workspace** — safest default.

Later possible strategies:

2. Move into workspace.
3. Reference external path without copying.

### Import Flow

1. User selects source folder/file.
2. User chooses target feed.
3. System scans files.
4. System creates an import batch.
5. System copies files into the feed folder.
6. System creates FileRef records.
7. System creates Post records.
8. System optionally runs AI classification/summarization.
9. System indexes content and metadata.

### Import Preservation

When importing a folder, preserve folder structure unless the user chooses flattening.

Example:

```txt
/source/documents/contracts/client-a/contract.pdf
```

Could become:

```txt
~/Homefeed/professional/imported/documents/contracts/client-a/contract.pdf
```

---

## 10. Search Requirements

MVP search should support:

- title
- body
- comments
- tags
- owner/profile
- feed
- filename
- path
- MIME type

Use SQLite FTS5 if available.

Search should return posts, not raw files only.

A search result should show:

- title
- feed
- owner
- file type
- snippet
- path
- tags
- import source if relevant

---

## 11. AI Agent Boundary

AI support is required, but the MVP should not depend on any specific provider.

Design an agent boundary/interface that can support:

- local models later
- OpenAI-compatible APIs later
- manual/no-AI mode
- background classification jobs

### Agent Responsibilities

Potential operations:

- summarize a file
- suggest title
- suggest feed
- suggest owner/profile
- suggest tags
- detect duplicates
- build a timeline from related posts
- answer natural language questions over indexed content
- identify stale/orphan metadata
- propose cleanup plans

### Agent Safety

The agent should not move/delete files automatically in MVP.

It may propose actions. User confirms before destructive or large organizing changes.

### Agent Interface Sketch

```go
type Agent interface {
    SummarizeFile(ctx context.Context, file FileRef) (SummaryResult, error)
    SuggestMetadata(ctx context.Context, post Post, file *FileRef) (MetadataSuggestion, error)
    Answer(ctx context.Context, query string) (AnswerResult, error)
}
```

---

## 12. UI Concept

### Main Views

1. **Main Feed**
   - unified feed
   - filter chips
   - search bar
   - post cards

2. **Feed View**
   - one feed/folder
   - feed-specific posts
   - import button
   - create post button

3. **Post Detail**
   - file preview or file metadata
   - owner/profile
   - comments
   - tags
   - AI summary
   - related posts

4. **Profiles**
   - list profiles
   - profile detail
   - posts owned by profile

5. **Imports**
   - import history
   - import status
   - failed items

6. **Settings**
   - workspace path
   - AI provider config placeholder
   - indexing options

### Post Card Fields

Each card should display:

- title
- owner
- feed
- file type/icon
- date
- tags
- short summary/snippet
- comment count

---

## 13. Backend Modules

Suggested Go module layout:

```txt
homefeed/
  app/                    # Wails app shell
  frontend/               # SolidJS UI
  internal/
    domain/               # entities and domain rules
    workspace/            # workspace creation, path conventions
    feeds/                # feed management
    profiles/             # profile management
    posts/                # post management
    files/                # file refs, hashing, mime detection
    imports/              # import scanning/copying
    index/                # sqlite, FTS, migrations
    search/               # query handling
    comments/             # comments
    tags/                 # tags
    agent/                # AI abstraction
    jobs/                 # background jobs
    config/               # app settings
    logging/              # logs
  migrations/
  docs/
```

---

## 14. MVP Milestones

### Milestone 1 — Workspace + Database

- Create Wails app.
- Create SolidJS frontend.
- On first run, initialize `~/Homefeed`.
- Create `.homefeed/index.sqlite`.
- Run migrations.
- Create default feeds:
  - professional
  - family
  - projects
  - personal
  - archive

### Milestone 2 — Feed and Post CRUD

- List feeds.
- List posts from all feeds.
- Create note post.
- Assign owner.
- Auto-create profile if owner does not exist.
- Add tags.
- Add comments.

### Milestone 3 — File Import

- Select file/folder.
- Choose target feed.
- Copy files into workspace.
- Create FileRef and Post records.
- Show import batch status.

### Milestone 4 — Search

- Add SQLite FTS5.
- Index posts/comments/file metadata.
- Add search UI.
- Add filters by feed/profile/tag/type/date.

### Milestone 5 — AI Boundary

- Add agent interface.
- Add no-op/mock agent.
- Add job queue for “suggest metadata”.
- Store AI suggestions separately from user-confirmed metadata.
- User can accept/reject suggestions.

### Milestone 6 — Polish

- File preview/open-in-system-file-manager.
- Duplicate detection by hash.
- Orphan metadata checks.
- Basic settings.
- Import error recovery.

---

## 15. UX Principles

1. The app should feel calmer than a traditional file browser.
2. The feed metaphor should reduce navigation anxiety.
3. Search should be prominent.
4. The user should not need to remember exact paths.
5. Files must remain normal files.
6. Metadata should not pollute visible folders.
7. AI should assist, not silently reorganize.
8. Destructive actions require confirmation.
9. The app should tolerate low-energy usage.
10. The system should remain understandable if the app breaks.

---

## 16. Technical Principles

1. Filesystem remains source of physical truth.
2. SQLite is the operational index.
3. `.homefeed/` contains system metadata.
4. Derived indexes can be rebuilt.
5. Use stable IDs for entities.
6. Track both original path and workspace path for imports.
7. Keep domain logic independent from Wails when possible.
8. Make backend functions reusable by a future CLI.
9. Keep AI provider replaceable.
10. Avoid hidden magic for file movements.

---

## 17. Open Questions

Codex should not block on these, but should make implementation choices explicit.

1. Should imported folders land under each feed root directly or under `feed/imported/`?
2. Should the first MVP support external references, or only copy imports?
3. Should metadata JSON files under `.homefeed/meta/` be canonical or exports from SQLite?
4. Should file watching be included in MVP or delayed?
5. Should the app support markdown note posts from the start?
6. Should profiles support avatars/colors in MVP?
7. Should search index file contents or only metadata initially?

Recommended MVP assumptions:

- Copy imports only.
- Store imported files under `~/Homefeed/<feed>/imported/<import_id>/...`.
- SQLite is canonical for MVP.
- `.homefeed/meta/` export can come later.
- File watching can come after import/search works.
- Search metadata first, file content later.

---

## 18. First Codex Task

Implement the initial skeleton:

1. Create Wails + SolidJS project.
2. Add Go backend workspace initialization.
3. Create `~/Homefeed` and `.homefeed/index.sqlite`.
4. Add SQLite migrations for feeds, profiles, posts, file_refs, comments, tags, post_tags, import_batches.
5. Seed default feeds.
6. Show main feed UI with empty state.
7. Show feeds list.
8. Add a simple “Create note post” flow.

Keep code modular so later tasks can add import/search/agent support cleanly.

---

## 19. Suggested Initial SQLite Schema

```sql
CREATE TABLE IF NOT EXISTS feeds (
  id TEXT PRIMARY KEY,
  name TEXT NOT NULL,
  slug TEXT NOT NULL UNIQUE,
  path TEXT NOT NULL,
  description TEXT,
  created_at TEXT NOT NULL,
  updated_at TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS profiles (
  id TEXT PRIMARY KEY,
  display_name TEXT NOT NULL,
  slug TEXT NOT NULL UNIQUE,
  kind TEXT NOT NULL,
  notes TEXT,
  created_at TEXT NOT NULL,
  updated_at TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS file_refs (
  id TEXT PRIMARY KEY,
  workspace_path TEXT NOT NULL,
  absolute_path TEXT NOT NULL,
  original_path TEXT,
  mime_type TEXT,
  size_bytes INTEGER,
  sha256 TEXT,
  modified_at TEXT,
  imported_at TEXT,
  import_id TEXT
);

CREATE TABLE IF NOT EXISTS posts (
  id TEXT PRIMARY KEY,
  feed_id TEXT NOT NULL REFERENCES feeds(id),
  owner_id TEXT NOT NULL REFERENCES profiles(id),
  title TEXT NOT NULL,
  kind TEXT NOT NULL,
  body TEXT,
  file_id TEXT REFERENCES file_refs(id),
  source TEXT,
  created_at TEXT NOT NULL,
  updated_at TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS comments (
  id TEXT PRIMARY KEY,
  post_id TEXT NOT NULL REFERENCES posts(id),
  author_id TEXT NOT NULL REFERENCES profiles(id),
  body TEXT NOT NULL,
  created_at TEXT NOT NULL,
  updated_at TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS tags (
  id TEXT PRIMARY KEY,
  name TEXT NOT NULL,
  slug TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS post_tags (
  post_id TEXT NOT NULL REFERENCES posts(id),
  tag_id TEXT NOT NULL REFERENCES tags(id),
  PRIMARY KEY (post_id, tag_id)
);

CREATE TABLE IF NOT EXISTS import_batches (
  id TEXT PRIMARY KEY,
  source_path TEXT NOT NULL,
  target_feed_id TEXT NOT NULL REFERENCES feeds(id),
  strategy TEXT NOT NULL,
  status TEXT NOT NULL,
  created_at TEXT NOT NULL,
  completed_at TEXT
);
```

---

## 20. Product Summary

Homefeed is a desktop, local-first, social-style file manager.

It creates a workspace under `~/Homefeed`, where folders behave as feeds. Files and notes become posts. Posts have owners, profiles, comments, tags, searchability, and AI-assisted organization.

The app should make local files feel more like a living knowledge stream than a static tree of folders, while still preserving normal filesystem access and long-term durability.
