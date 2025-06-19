package textfsm

import (
	"fmt"

	"github.com/sirikothe/gotextfsm"
)

func Parse(template, text string) ([]map[string]any, error) {
	fsm := gotextfsm.TextFSM{}

	err := fsm.ParseString(string(template))
	if err != nil {
		return nil, fmt.Errorf("failed to parse template: %w", err)
	}

	parser := gotextfsm.ParserOutput{}

	err = parser.ParseTextString(text, fsm, true)
	if err != nil {
		return nil, fmt.Errorf("failed to parse text: %w", err)
	}

	return parser.Dict, nil
}
