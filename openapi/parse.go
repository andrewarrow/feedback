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
	for _, line := range lines {
		fmt.Println(len(line))
	}
}
