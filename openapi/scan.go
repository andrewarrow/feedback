package openapi

import (
	"os"
	"strings"
)

type OpenAPI struct {
	Endpoints  map[string][]Endpoint
	FuncToPath map[string]string
}

func NewOpenAPI() *OpenAPI {
	oa := OpenAPI{}
	oa.Endpoints = map[string][]Endpoint{}
	oa.FuncToPath = map[string]string{}
	return &oa
}

func (oa *OpenAPI) ScanDir(dir string) {
	entries, _ := os.ReadDir(dir)
	for _, entry := range entries {
		name := entry.Name()
		if strings.HasSuffix(name, ".go") == false {
			continue
		}
		b, _ := os.ReadFile(dir + "/" + name)
		s := string(b)
		lines := strings.Split(s, "\n")
		lastFunc := ""
		for i, line := range lines {
			trimmed := strings.TrimSpace(line)
			if strings.HasPrefix(trimmed, "func ") {
				tokens := strings.Split(trimmed, " ")
				tokens = strings.Split(tokens[1], "(")
				lastFunc = tokens[0]
			}
			if strings.HasPrefix(trimmed, "// oa ") == false {
				continue
			}
			if strings.HasPrefix(trimmed, "// oa start") == true {
				continue
			}
			if strings.HasPrefix(trimmed, "// oa end") == true {
				continue
			}
			target := lines[i+1]
			ep := NewEndpoint(trimmed, target, lastFunc)
			prefix := oa.FuncToPath[lastFunc]
			oa.Endpoints[prefix+ep.Path] = append(oa.Endpoints[ep.Path], ep)
		}
	}
}
