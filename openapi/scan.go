package openapi

import (
	"os"
	"strings"
)

type OpenAPI struct {
	Endpoints map[string][]Endpoint
}

func ScanDir(dir string) *OpenAPI {
	oa := OpenAPI{}
	entries, _ := os.ReadDir(dir)
	for _, entry := range entries {
		name := entry.Name()
		if strings.HasSuffix(name, ".go") == false {
			continue
		}
		b, _ := os.ReadFile(dir + "/" + name)
		s := string(b)
		lines := strings.Split(s, "\n")
		for i, line := range lines {
			trimmed := strings.TrimSpace(line)
			if strings.HasPrefix(trimmed, "// oa ") == false {
				continue
			}
			target := lines[i+1]
			ep := NewEndpoint(trimmed, target)
			oa.Endpoints[ep.Path] = append(oa.Endpoints[ep.Path], ep)
		}
	}
	return &oa
}
