package openapi

import (
	"fmt"
	"os"
	"strings"
)

func Parse(path string) {
	b, _ := os.ReadFile(path)
	s := string(b)
	lines := strings.Split(s, "\n")
	items := []Endpoint{}
	for i, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "// oa ") == false {
			continue
		}
		fmt.Println(line)
		target := lines[i+1]
		ep := NewEndpoint(trimmed, target)
		items = append(items, ep)
	}

	fmt.Println(items)
}
