const go = new Go();

const loadingModal = document.getElementById("loadingModal");
const loadingState = document.getElementById("loadingState");
const errorState = document.getElementById("errorState");
const errorMessage = document.getElementById("errorMessage");

const showError = (error) => {
  loadingState.classList.add("hidden");
  errorState.classList.remove("hidden");
  errorMessage.textContent = error.toString();
};

// Load and initialize WASM
const loadWasm = async () => {
  const result = await WebAssembly.instantiateStreaming(
    fetch("main.wasm"),
    go.importObject
  ).catch((error) => {
    console.error("Failed to load WASM:", error);
    showError(new Error("Failed to load the WASM module"));
    throw error;
  });

  try {
    go.run(result.instance);
    loadingModal.classList.add("hidden");
    return result;
  } catch (error) {
    console.error("Failed to execute WASM:", error);
    showError(new Error("Failed to execute the WASM module"));
    throw error;
  }
};

// Initialize WebAssembly and export the add function
const result = await loadWasm();
const wasmAdd = result.instance.exports.add;

// Example usage
const sum = wasmAdd(5, 3);
console.log("Result of add(5, 3):", sum);
