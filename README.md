# textfsm playgorund

This playground is inspired by https://textfsm.nornir.tech/. This is something I have occasionally used to sanity check [TextFSM](https://github.com/google/textfsm/wiki/TextFSM) or [Jinja](https://jinja.palletsprojects.com/en/stable/) templates.

## Why?

- I wanted a nicer UI
- I wanted to ensure that anything I input was never leaving the browser and was always present on the client side.

## Run locally

Run `make -C wasm` to generate all the relevant HTML, JS and WASM files at `wasm/dist/*` which can be served by whatever you choose

There is a simple go file server which you can used by running `go run cmd/server/main.go`.

## Supported Features

- [x] TextFSM parsing
- [] JSON Schema Validation
