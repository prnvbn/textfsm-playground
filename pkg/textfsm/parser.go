package textfsm

import (
	"github.com/sirikothe/gotextfsm"
)

func Parse(template, text string) ([]map[string]any, error) {
	fsm := gotextfsm.TextFSM{}

	err := fsm.ParseString(string(template))
	if err != nil {
		return nil, err
	}

	parser := gotextfsm.ParserOutput{}

	err = parser.ParseTextString(text, fsm, true)
	if err != nil {
		return nil, err
	}

	return parser.Dict, nil
}
