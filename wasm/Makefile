.PHONY: build clean

build: wasm_exec.js main.wasm
	mkdir -p dist
	cp -r pages dist
	cp *.js dist
	cp *.html dist
	cp main.wasm dist
	cp ../version.txt dist


wasm_exec.js:
	cp $(shell go env GOROOT)/lib/wasm/wasm_exec.js .

# TODO: use wasm opt to optimize the wasm binary
main.wasm: main.go $(shell find .. -name '*.go')
	GOOS=js GOARCH=wasm go build -o main.wasm ./main.go

clean:
	rm -f wasm_exec.js main.wasm
	rm -rf dist


