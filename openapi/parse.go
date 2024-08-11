package openapi

import (
	"fmt"
	"os"
	"strings"
)

func Parse(path, dir string) {
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
	entries, _ := os.ReadDir(dir)
	for _, entry := range entries {
		fmt.Println(entry.Name())
		lookForParams(dir + "/" + entry.Name())
	}

}

func lookForParams(path string) {
	b, _ := os.ReadFile(path)
	s := string(b)
	lines := strings.Split(s, "\n")
	start := false
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "// oa end") {
			start = false
		}
		if start {
			fmt.Println(trimmed)
		}
		if strings.HasPrefix(trimmed, "// oa start") {
			start = true
		}
	}

}
