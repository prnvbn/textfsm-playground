//go:build js && wasm

package main

import (
	"encoding/json"
	"fmt"
	"go-rest/pkg/textfsm"
	"syscall/js"
	"unsafe"

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

		log.Info().Str("parsed", fmt.Sprintf("%v", parsed)).Msg("parsed")

		if err != nil {
			return js.ValueOf(map[string]any{
				"error": err.Error(),
			})
		}

		log.Info().Msg("parsing textfsm")

		return js.ValueOf(map[string]any{
			"data": parsed,
		})
	}))

	select {}
}

//export add
func add(x int, y int) int {
	return x + y
}

func parseTextFSM(template string, text string) (string, error) {
	parsed, err := textfsm.Parse(template, text)
	if err != nil {
		return "", fmt.Errorf("failed to parse textfsm: %w", err)
	}

	json, err := json.Marshal(parsed)
	if err != nil {
		return "", fmt.Errorf("failed to marshal parsed: %w", err)
	}

	return string(json), nil
}

//export parse
func parse(templatePtr, templateLen, textPtr, textLen int) {
	templateBytes := unsafe.Slice((*byte)(unsafe.Pointer(uintptr(templatePtr))), templateLen)
	textBytes := unsafe.Slice((*byte)(unsafe.Pointer(uintptr(textPtr))), textLen)

	templateStr := string(templateBytes)
	textStr := string(textBytes)

	log := log.With().
		Str("template", templateStr).
		Str("text", textStr).
		Logger()

	log.Info().Msg("parse called")

	// 	parsed, err := textfsm.Parse(templateStr, textStr)
	// 	if err != nil {
	// 		log.Error().Err(err).Msg("failed to parse textfsm")
	// 		return "", err
	// 	}

	// 	json, err := json.Marshal(parsed)
	// 	if err != nil {
	// 		log.Error().Err(err).Msg("failed to marshal parsed")
	// 		return "", err
	// 	}

	// return string(json), nil
}
