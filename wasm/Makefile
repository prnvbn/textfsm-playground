all: wasm_exec.js main.wasm

wasm_exec.js:
	cp $(shell go env GOROOT)/lib/wasm/wasm_exec.js .

main.wasm: main.go $(shell find .. -name '*.go')
	GOOS=js GOARCH=wasm go build -o main.wasm ./main.go

clean:
	rm -f wasm_exec.js main.wasm