package jinja_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/prnvbn/textfsm-playground/pkg/jinja"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type Testcase struct {
	Name     string
	Template string
	Data     map[string]any
	Rendered string
}

func TestJinjaParse(t *testing.T) {
	tests := []Testcase{}

	files, err := os.ReadDir("tests/")
	assert.NoError(t, err)

	for _, file := range files {
		if !file.IsDir() {
			continue
		}

		templatePath := fmt.Sprintf("tests/%s/template.jinja", file.Name())
		template, err := os.ReadFile(templatePath)
		require.NoError(t, err, "failed to read template file in %s", templatePath)

		dataPath := fmt.Sprintf("tests/%s/data.json", file.Name())
		dataBytes, err := os.ReadFile(dataPath)
		require.NoError(t, err, "failed to read data file in %s", dataPath)

		var data map[string]any
		err = json.Unmarshal(dataBytes, &data)
		require.NoError(t, err, "failed to unmarshal data in %s", dataPath)

		renderedPath := fmt.Sprintf("tests/%s/rendered.txt", file.Name())
		renderedBytes, err := os.ReadFile(renderedPath)
		require.NoError(t, err, "failed to read rendered file in %s", renderedPath)

		tests = append(tests, Testcase{
			Name:     file.Name(),
			Template: string(template),
			Data:     data,
			Rendered: string(renderedBytes),
		})

		if os.Getenv("UPDATE_SNAPSHOTS") == "true" {
			rendered, err := jinja.Parse(string(template), data)
			require.NoError(t, err)
			err = os.WriteFile(renderedPath, []byte(rendered), 0644)
			require.NoError(t, err)
		}
	}

	for _, test := range tests {
		rendered, err := jinja.Parse(test.Template, test.Data)
		require.NoError(t, err)

		if ok := assert.Equal(t, test.Rendered, rendered); !ok {
			t.Log(cmp.Diff(test.Rendered, rendered))
			t.Fail()
		}
	}
}
