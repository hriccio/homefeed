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

function hasBridge() {
  return Boolean(window.go?.main?.App?.InitializeWorkspace);
}

export function App() {
  const [message, setMessage] = createSignal("Ready to initialize Homefeed.");
  const [busy, setBusy] = createSignal(false);
  const [result, setResult] = createSignal<WorkspaceResult | null>(null);

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
          The first SolidJS screen proves the workspace initializer through the
          Wails bridge without adding import, search, or AI behavior.
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
    </main>
  );
}

