import { createSignal, onMount, Show } from "solid-js";

type WorkspaceResult = {
  layout?: {
    root?: string;
    databasePath?: string;
  };
  feeds?: Array<{
    slug?: string;
    name?: string;
  }>;
};

type ImportResult = {
  batchId?: number;
  sourcePath?: string;
  feedSlug?: string;
  feedPath?: string;
  destinationPath?: string;
};

type NoteResult = {
  postId?: number;
  kind?: string;
  feedSlug?: string;
  feedPath?: string;
  title?: string;
  body?: string;
  path?: string;
};

function hasBridge() {
  return Boolean(window.go?.main?.App?.InitializeWorkspace);
}

export function App() {
  const [message, setMessage] = createSignal("Ready to initialize Homefeed.");
  const [busy, setBusy] = createSignal(false);
  const [result, setResult] = createSignal<WorkspaceResult | null>(null);
  const [sourcePath, setSourcePath] = createSignal("");
  const [feedSlug, setFeedSlug] = createSignal("projects");
  const [importMessage, setImportMessage] = createSignal(
    "Choose a local folder to copy into the workspace.",
  );
  const [importBusy, setImportBusy] = createSignal(false);
  const [importResult, setImportResult] = createSignal<ImportResult | null>(
    null,
  );
  const [noteFeedSlug, setNoteFeedSlug] = createSignal("projects");
  const [noteTitle, setNoteTitle] = createSignal("");
  const [noteBody, setNoteBody] = createSignal("");
  const [noteMessage, setNoteMessage] = createSignal(
    "Write a note title and body to create a post.",
  );
  const [noteBusy, setNoteBusy] = createSignal(false);
  const [noteResult, setNoteResult] = createSignal<NoteResult | null>(null);
  const [bridgeReady, setBridgeReady] = createSignal(false);

  const refreshBridgeState = () => {
    setBridgeReady(hasBridge());
  };

  const initializeWorkspace = async () => {
    const bridge = window.go?.main?.App?.InitializeWorkspace;
    if (!bridge) {
      setMessage("Wails bridge is unavailable in this runtime.");
      return;
    }

    setBusy(true);
    setMessage("Initializing workspace...");

    try {
      const response = (await bridge()) as WorkspaceResult;
      setResult(response);
      setMessage("Workspace initialized.");
    } catch (error) {
      setMessage(`Initialization failed: ${String(error)}`);
    } finally {
      setBusy(false);
    }
  };

  const importFolder = async () => {
    const bridge = window.go?.main?.App?.ImportFolder;
    if (!bridge) {
      setImportMessage("Import bridge is unavailable in this runtime.");
      return;
    }

    if (!sourcePath().trim()) {
      setImportMessage("Enter a source folder path first.");
      return;
    }

    if (!feedSlug().trim()) {
      setImportMessage("Choose a target feed slug first.");
      return;
    }

    setImportBusy(true);
    setImportMessage("Importing folder...");

    try {
      const response = (await bridge(
        sourcePath().trim(),
        feedSlug().trim(),
      )) as ImportResult;
      setImportResult(response);
      setImportMessage("Folder imported.");
    } catch (error) {
      setImportMessage(`Import failed: ${String(error)}`);
    } finally {
      setImportBusy(false);
    }
  };

  const createNotePost = async () => {
    const bridge = window.go?.main?.App?.CreateNotePost;
    if (!bridge) {
      setNoteMessage("Note bridge is unavailable in this runtime.");
      return;
    }

    if (!noteFeedSlug().trim()) {
      setNoteMessage("Choose a target feed slug first.");
      return;
    }

    if (!noteTitle().trim()) {
      setNoteMessage("Enter a note title first.");
      return;
    }

    setNoteBusy(true);
    setNoteMessage("Creating note...");

    try {
      const response = (await bridge(
        noteFeedSlug().trim(),
        noteTitle().trim(),
        noteBody().trim(),
      )) as NoteResult;
      setNoteResult(response);
      setNoteMessage("Note created.");
    } catch (error) {
      setNoteMessage(`Note creation failed: ${String(error)}`);
    } finally {
      setNoteBusy(false);
    }
  };

  onMount(() => {
    refreshBridgeState();
    const timer = window.setInterval(refreshBridgeState, 100);
    window.setTimeout(() => {
      clearInterval(timer);
    }, 5000);
  });

  return (
    <main class="shell">
      <section class="panel hero">
        <p class="eyebrow">Homefeed</p>
        <h1>Minimal desktop shell</h1>
        <p class="lede">
          The first SolidJS screen proves workspace initialization, folder
          import, and note-post creation through the Wails bridge without
          adding search or AI behavior.
        </p>
      </section>

      <section class="panel control">
        <button
          type="button"
          onClick={initializeWorkspace}
          disabled={busy() || !bridgeReady()}
        >
          {busy() ? "Initializing..." : "Initialize workspace"}
        </button>

        <p class="status">{message()}</p>

        <Show when={result()}>
          {(value) => (
            <pre>{JSON.stringify(value(), null, 2)}</pre>
          )}
        </Show>

        <Show when={!bridgeReady()}>
          <p class="status">Waiting for the Wails bridge...</p>
        </Show>
      </section>

      <section class="panel control">
        <h2>Import a local folder</h2>
        <label>
          <span>Source folder path</span>
          <input
            type="text"
            value={sourcePath()}
            onInput={(event) => setSourcePath(event.currentTarget.value)}
            placeholder="/home/henrique/Downloads/sample"
          />
        </label>

        <label>
          <span>Target feed slug</span>
          <input
            type="text"
            value={feedSlug()}
            onInput={(event) => setFeedSlug(event.currentTarget.value)}
            placeholder="projects"
          />
        </label>

        <button
          type="button"
          onClick={importFolder}
          disabled={importBusy() || !bridgeReady()}
        >
          {importBusy() ? "Importing..." : "Import folder"}
        </button>

        <p class="status">{importMessage()}</p>

        <Show when={importResult()}>
          {(value) => (
            <pre>{JSON.stringify(value(), null, 2)}</pre>
          )}
        </Show>
      </section>

      <section class="panel control">
        <h2>Create a note post</h2>
        <label>
          <span>Target feed slug</span>
          <input
            type="text"
            value={noteFeedSlug()}
            onInput={(event) => setNoteFeedSlug(event.currentTarget.value)}
            placeholder="projects"
          />
        </label>

        <label>
          <span>Note title</span>
          <input
            type="text"
            value={noteTitle()}
            onInput={(event) => setNoteTitle(event.currentTarget.value)}
            placeholder="Daily log"
          />
        </label>

        <label>
          <span>Note body</span>
          <textarea
            rows={6}
            value={noteBody()}
            onInput={(event) => setNoteBody(event.currentTarget.value)}
            placeholder="Write the note body here."
          />
        </label>

        <button
          type="button"
          onClick={createNotePost}
          disabled={noteBusy() || !bridgeReady()}
        >
          {noteBusy() ? "Creating..." : "Create note"}
        </button>

        <p class="status">{noteMessage()}</p>

        <Show when={noteResult()}>
          {(value) => (
            <pre>{JSON.stringify(value(), null, 2)}</pre>
          )}
        </Show>
      </section>
    </main>
  );
}
