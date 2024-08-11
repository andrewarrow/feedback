package openapi

import (
	"os"
	"strings"
)

func Parse(path, dir string) {
	b, _ := os.ReadFile(path)
	s := string(b)
	lines := strings.Split(s, "\n")
	m := map[string][]Endpoint{}
	for i, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "// oa ") == false {
			continue
		}
		target := lines[i+1]
		ep := NewEndpoint(trimmed, target)
		m[ep.Path] = append(m[ep.Path], ep)
	}

	//fmt.Println(items)
	entries, _ := os.ReadDir(dir)
	for _, entry := range entries {
		//fmt.Println(entry.Name())
		lookForParams(dir + "/" + entry.Name())
	}

	MakeYaml(m)

}

func lookForParams(path string) {
	b, _ := os.ReadFile(path)
	s := string(b)
	lines := strings.Split(s, "\n")
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
			//fmt.Println(lastFunc, trimmed)
			_ = lastFunc
		}
		if strings.HasPrefix(trimmed, "// oa start") {
			start = true
		}
	}

}
