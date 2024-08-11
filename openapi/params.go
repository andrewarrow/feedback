package openapi

import (
	"fmt"
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
			//  lat, _ := c.Params["latitude"].(float64)
			tokens := strings.Split(trimmed, ":=")
			item := tokens[1][10:]
			fmt.Println(item)
			p := Param{}
			oa.ParamsByFunc[lastFunc] = append(oa.ParamsByFunc[lastFunc], p)
		}
		if strings.HasPrefix(trimmed, "// oa start") {
			start = true
		}
	}

}
