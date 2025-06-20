//go:build js && wasm

package main

import (
	"encoding/json"
	"fmt"
	"go-rest/pkg/textfsm"
	"syscall/js"

	"github.com/rs/zerolog/log"
)

// fn defines a standard signature for our Go functions.
type fn func(args []js.Value) (any, error)

// promisify takes a function and wraps it in a JavaScript Promise.
// This allows for clean async/await syntax on the JS side.

// for example, say this was used to wrap foo(s string) (string, error)
// then, this can be called from JavaScript like this:
//
//	foo("hello").then(result => {
//		console.log(result);
//	}).catch(error => {
//		console.log(error);
//	});
func promisify(fn fn) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		promiseConstructor := js.Global().Get("Promise")
		return promiseConstructor.New(js.FuncOf(func(this js.Value, promiseArgs []js.Value) any {
			resolve := promiseArgs[0]
			reject := promiseArgs[1]

			go func() {
				result, err := fn(args)
				if err != nil {
					errorConstructor := js.Global().Get("Error")
					jsError := errorConstructor.New(err.Error())
					reject.Invoke(jsError)
					return
				}

				jsonResult, err := json.Marshal(result)
				if err != nil {
					errorConstructor := js.Global().Get("Error")
					jsError := errorConstructor.New("failed to marshal result to JSON: " + err.Error())
					reject.Invoke(jsError)
					return
				}

				resolve.Invoke(string(jsonResult))
			}()

			return nil
		}))
	})
}

func parseTextFSMHandler(args []js.Value) (any, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("expected 2 arguments, got %d", len(args))
	}
	template := args[0].String()
	text := args[1].String()

	log := log.With().
		Str("template", template).
		Str("text", text).
		Logger()
	log.Info().Msg("parsing textfsm")

	return textfsm.Parse(template, text)
}

func main() {
	// Register our Go function, now wrapped in a promise.
	js.Global().Set("parseTextFSM", promisify(parseTextFSMHandler))

	// Keep the Go program alive.
	select {}
}
