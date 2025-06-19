package textfsm_test

import (
	"encoding/json"
	"fmt"
	"go-rest/pkg/textfsm"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type Testcase struct {
	Name     string
	Template string
	Text     string
	Parsed   []map[string]any
}

func TestTextFSMParse(t *testing.T) {

	tests := []Testcase{}

	files, err := os.ReadDir("tests/")
	assert.NoError(t, err)

	for _, file := range files {
		if !file.IsDir() {
			continue
		}

		templatePath := fmt.Sprintf("tests/%s/template.textfsm", file.Name())
		template, err := os.ReadFile(templatePath)
		require.NoError(t, err, "failed to read template file in %s", templatePath)

		textPath := fmt.Sprintf("tests/%s/text.txt", file.Name())
		text, err := os.ReadFile(textPath)
		require.NoError(t, err, "failed to read text file in %s", textPath)

		parsedPath := fmt.Sprintf("tests/%s/parsed.json", file.Name())
		parsedBytes, err := os.ReadFile(parsedPath)
		require.NoError(t, err, "failed to read parsed file in %s", parsedPath)

		var parsed []map[string]any
		err = json.Unmarshal(parsedBytes, &parsed)
		require.NoError(t, err, "failed to unmarshal parsed data in %s", parsedPath)

		tests = append(tests, Testcase{
			Name:     file.Name(),
			Template: string(template),
			Text:     string(text),
			Parsed:   parsed,
		})

	}

	for _, test := range tests {
		parsed, err := textfsm.Parse(test.Template, test.Text)
		require.NoError(t, err)

		if ok := assert.Equal(t, test.Parsed, parsed); !ok {
			t.Log(cmp.Diff(test.Parsed, parsed))
			t.Fail()
		}
	}
}
