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

  onMount(() => {
    if (!hasBridge()) {
      setMessage("Desktop bridge not yet attached.");
    }
  });

  return (
    <main class="shell">
      <section class="panel hero">
        <p class="eyebrow">Homefeed</p>
        <h1>Minimal desktop shell</h1>
        <p class="lede">
          The first SolidJS screen proves the workspace initializer and a
          minimal folder import workflow through the Wails bridge without
          adding search or AI behavior.
        </p>
      </section>

      <section class="panel control">
        <button type="button" onClick={initializeWorkspace} disabled={busy()}>
          {busy() ? "Initializing..." : "Initialize workspace"}
        </button>

        <p class="status">{message()}</p>

        <Show when={result()}>
          {(value) => (
            <pre>{JSON.stringify(value(), null, 2)}</pre>
          )}
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

        <button type="button" onClick={importFolder} disabled={importBusy()}>
          {importBusy() ? "Importing..." : "Import folder"}
        </button>

        <p class="status">{importMessage()}</p>

        <Show when={importResult()}>
          {(value) => (
            <pre>{JSON.stringify(value(), null, 2)}</pre>
          )}
        </Show>
      </section>
    </main>
  );
}
