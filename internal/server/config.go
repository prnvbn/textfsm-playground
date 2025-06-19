package server

type Config struct {
	Port    int    `yaml:"port"`
	WasmDir string `yaml:"wasm_dir"`
}
