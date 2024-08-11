package openapi

import (
	"os"
	"strings"
)

func Parse(path, dir string) {

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
