package openapi

import (
	"strings"
)

type Param struct {
	Name   string
	Flavor string
}

func (oa *OpenAPI) lookForParams(name string, lines []string) {
	start := false
	lastFunc := ""
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "func ") {
			tokens := strings.Split(trimmed, " ")
			tokens = strings.Split(tokens[1], "(")
			lastFunc = tokens[0]
		}
		if strings.HasPrefix(trimmed, "// oa end") {
			start = false
		}
		if start {
			//  ].(float64)
			tokens := strings.Split(trimmed, ":=")
			item := tokens[1][11:]
			tokens = strings.Split(item, `"`)
			p := Param{}
			p.Name = tokens[0]
			flavor := tokens[1]
			p.Flavor = "number"
			if strings.Contains(flavor, "string") {
				p.Flavor = "string"
			} else if strings.Contains(flavor, "int") {
				p.Flavor = "integer"
			} else if strings.Contains(flavor, "bool") {
				p.Flavor = "boolean"
			}
			oa.ParamsByFunc[lastFunc] = append(oa.ParamsByFunc[lastFunc], p)
		}
		if strings.HasPrefix(trimmed, "// oa start") {
			start = true
		}
	}

}
