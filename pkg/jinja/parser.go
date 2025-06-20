package jinja

import (
	"github.com/nikolalohinski/gonja/v2"
	"github.com/nikolalohinski/gonja/v2/exec"
)

func Parse(templateStr string, dataMap map[string]any) (string, error) {
	// template, err := gonja.FromString("Hello {{ name | capitalize }}!")
	tpl, err := gonja.FromString(templateStr)
	if err != nil {
		panic(err)
	}

	// data := exec.NewContext(map[string]interface{}{
	// 	"name": "bob",
	// })

	data := exec.NewContext(dataMap)

	return tpl.ExecuteToString(data)
}
