import { wasmPromise } from "../../loader.js";

// Function to initialize editors
function createEditors() {
  const templateEditor = document.getElementById("template-editor");
  const inputEditor = document.getElementById("input-editor");

  const onEditorReady = (editor) => {
    return new Promise((resolve) => {
      if (editor.editor) {
        resolve();
      } else {
        editor.addEventListener("ready", resolve, { once: true });
      }
    });
  };

  Promise.all([onEditorReady(templateEditor), onEditorReady(inputEditor)]).then(
    () => {
      templateEditor.value = [
        "Value UPTIME (\\d+:\\d+:\\d+)",
        "Value VERSION (\\d+\\.\\d+\\.[A-Z]+)",
        "",
        "Start",
        "  ^Time since last reboot: ${UPTIME}",
        "  ^Version: ${VERSION} -> Record",
      ].join("\n");

      inputEditor.value = [
        "System Status:",
        "Time since last reboot: 5:10:23",
        "Version: 2.1.GA",
        "Memory Usage: 65%",
      ].join("\n");
    }
  );

  // Update editor themes when the page theme changes
  const observer = new MutationObserver(function (mutations) {
    mutations.forEach(function (mutation) {
      if (mutation.attributeName === "class") {
        const isDark = document.documentElement.classList.contains("dark");
        const theme = isDark ? "vs-dark" : "vs-light";
        if (window.monaco) {
          monaco.editor.setTheme(theme);
        }
      }
    });
  });

  observer.observe(document.documentElement, {
    attributes: true,
    attributeFilter: ["class"],
  });

  document
    .getElementById("parse-button")
    .addEventListener("click", async function () {
      const template = templateEditor.value;
      const input = inputEditor.value;
      const resultsEl = document.getElementById("results");

      try {
        await wasmPromise;

        resultsEl.textContent = "Parsing...";
        resultsEl.classList.remove("text-red-500");

        const jsonString = await window.parseTextFSM(template, input);

        const formattedJson = JSON.stringify(JSON.parse(jsonString), null, 2);
        resultsEl.textContent = formattedJson;
      } catch (error) {
        console.error("An error occurred:", error);
        resultsEl.textContent = "Error: " + error.message;
        resultsEl.classList.add("text-red-500");
      }
    });
}

export function init() {
  createEditors();
}
