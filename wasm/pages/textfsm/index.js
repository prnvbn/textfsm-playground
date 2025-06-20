import { wasmPromise } from "../../loader.js";

// Function to initialize editors
function createEditors() {
  const templateEditor = document.getElementById("template-editor");
  const inputEditor = document.getElementById("input-editor");
  const resultsEditor = document.getElementById("results-editor");

  const onEditorReady = (editor) => {
    return new Promise((resolve) => {
      if (editor.editor) {
        resolve();
      } else {
        editor.addEventListener("ready", resolve, { once: true });
      }
    });
  };

  Promise.all([
    onEditorReady(templateEditor),
    onEditorReady(inputEditor),
    onEditorReady(resultsEditor),
  ]).then(() => {
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

    resultsEditor.editor.updateOptions({ readOnly: true });
  });

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
      const resultsEditor = document.getElementById("results-editor");
      const resultsContainer = resultsEditor.parentElement;

      const flashBorder = (success) => {
        const originalDarkBorder = "dark:border-github-dark-border";
        const successClasses = ["border-green-500", "dark:border-green-500"];
        const errorClasses = ["border-red-500", "dark:border-red-500"];

        const classesToAdd = success ? successClasses : errorClasses;

        resultsContainer.classList.remove(originalDarkBorder);
        resultsContainer.classList.add(...classesToAdd);

        setTimeout(() => {
          resultsContainer.classList.remove(...classesToAdd);
          resultsContainer.classList.add(originalDarkBorder);
        }, 1000);
      };

      try {
        await wasmPromise;

        resultsEditor.value = "Parsing...";

        const jsonString = await window.parseTextFSM(template, input);

        const formattedJson = JSON.stringify(JSON.parse(jsonString), null, 2);
        resultsEditor.value = formattedJson;
        flashBorder(true);
      } catch (error) {
        console.error("An error occurred:", error);
        resultsEditor.value = "Error: " + error.message;
        flashBorder(false);
      }
    });
}

export function init() {
  createEditors();
}
