import { wasmPromise } from "../../loader.js";
import { registerJinjaLanguage } from "./jinja-language.js";

// Function to initialize editors
function createEditors() {
  const templateEditor = document.getElementById("template-editor");
  const dataEditor = document.getElementById("data-editor");
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
    onEditorReady(dataEditor),
    onEditorReady(resultsEditor),
  ]).then(() => {
    // Register Jinja language after Monaco is available
    registerJinjaLanguage();

    // Set default Jinja template
    templateEditor.value = [
      "Hello {{ name }}!",
      "",
      "{%- if age %}",
      "You are {{ age }} years old.",
      "{% endif %}",
      "",
      "{%- for item in items %}",
      "- {{ item }}",
      "{%- endfor %}",
    ].join("\n");

    // Set default JSON data
    dataEditor.value = JSON.stringify(
      {
        name: "World",
        age: 25,
        items: ["apple", "banana", "cherry"],
      },
      null,
      2
    );

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
    .getElementById("render-button")
    .addEventListener("click", async function () {
      const template = templateEditor.value;
      const dataString = dataEditor.value;
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

        resultsEditor.value = "Rendering...";

        let data;
        try {
          data = JSON.parse(dataString);
        } catch (jsonError) {
          throw new Error("Invalid JSON data: " + jsonError.message);
        }

        const renderedOutput = await window.renderJinja(template, dataString);

        resultsEditor.value = renderedOutput;
        flashBorder(true);
      } catch (error) {
        console.error("An error occurred:", error);
        resultsEditor.value = "Error: " + error.message;
        flashBorder(false);
      }
    });

  document.addEventListener("keydown", function (event) {
    if (event.shiftKey && event.key === "Enter") {
      event.preventDefault();
      parseFunction();
    }
  });
}

export function init() {
  createEditors();
}
