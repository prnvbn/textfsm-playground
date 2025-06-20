package jinja

import (
	"github.com/nikolalohinski/gonja/v2"
	"github.com/nikolalohinski/gonja/v2/exec"
)

func Render(templateStr string, dataMap map[string]any) (string, error) {
	tpl, err := gonja.FromString(templateStr)
	if err != nil {
		return "", err
	}

	data := exec.NewContext(dataMap)
	return tpl.ExecuteToString(data)
}
