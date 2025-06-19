all: wasm_exec.js main.wasm

wasm_exec.js:
	cp $(shell tinygo env TINYGOROOT)/targets/wasm_exec.js .

main.wasm:
	tinygo build -o main.wasm -target wasm ./main.go

clean:
	rm -f wasm_exec.js main.wasm