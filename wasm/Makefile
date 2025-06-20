.PHONY: all build clean

all: wasm_exec.js main.wasm

wasm_exec.js:
	cp $(shell go env GOROOT)/lib/wasm/wasm_exec.js .

# TODO: use wasm opt to optimize the wasm binary
main.wasm: main.go $(shell find .. -name '*.go')
	GOOS=js GOARCH=wasm go build -o main.wasm ./main.go

clean:
	rm -f wasm_exec.js main.wasm dist/*


build: wasm_exec.js main.wasm
	mkdir -p dist
	cp wasm_exec.js main.wasm dist
	cp -r pages dist
	cp *.js dist
	cp *.html dist

