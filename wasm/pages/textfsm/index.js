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

  // Handle parse button click
  document
    .getElementById("parse-button")
    .addEventListener("click", function () {
      const template = templateEditor.value;
      const input = inputEditor.value;

      wasmPromise
        .then((result) => {
          console.log(result);
          // TODO: Add parsing logic here, using result.instance.exports
          document.getElementById("results").textContent =
            "Parsing not implemented yet";
        })
        .catch((error) => {
          console.error("Error accessing WebAssembly module:", error);
          const results = document.getElementById("results");
          results.textContent =
            "Error: The WebAssembly module failed to load. Please check the console for details.";
          results.classList.add("text-red-500");
        });
    });
}

export function init() {
  createEditors();
}
