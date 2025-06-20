//go:build js && wasm

package main

import (
	"encoding/json"
	"go-rest/pkg/textfsm"
	"syscall/js"

	"github.com/rs/zerolog/log"
)

func main() {
	js.Global().Set("foo", js.FuncOf(func(this js.Value, args []js.Value) any {
		str := args[0].String()
		log.Info().Msg(str)
		return nil
	}))

	js.Global().Set("parseTextFSM", js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 2 {
			return js.ValueOf(map[string]any{
				"error": "expected 2 arguments",
			})
		}

		template := args[0].String()
		text := args[1].String()

		log := log.With().
			Str("template", template).
			Str("text", text).
			Logger()

		log.Info().Msg("parsing textfsm")

		parsed, err := textfsm.Parse(template, text)
		if err != nil {
			return js.ValueOf(map[string]any{
				"error": err.Error(),
			})
		}

		jsonResult, err := json.Marshal(parsed)
		if err != nil {
			return js.ValueOf(map[string]any{
				"error": "failed to marshal result to JSON: " + err.Error(),
			})
		}

		return js.ValueOf(map[string]any{
			"data": string(jsonResult),
		})
	}))

	select {}
}
