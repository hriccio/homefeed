const output = document.getElementById("output");
const initializeButton = document.getElementById("initialize");

function setOutput(message) {
  output.textContent = message;
}

async function initializeWorkspace() {
  const api = window.go && window.go.main && window.go.main.App;
  if (!api || typeof api.InitializeWorkspace !== "function") {
    setOutput("Wails bridge is unavailable in this runtime.");
    return;
  }

  try {
    const result = await api.InitializeWorkspace();
    setOutput(JSON.stringify(result, null, 2));
  } catch (error) {
    setOutput(`Initialization failed: ${error}`);
  }
}

initializeButton.addEventListener("click", initializeWorkspace);
initializeWorkspace();

