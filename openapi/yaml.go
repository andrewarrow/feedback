package openapi

import (
	"fmt"
	"strings"
)

func MakeYaml(m map[string][]Endpoint) {

	buffer := []string{}

	for k, _ := range m {
		buffer = append(buffer, "  "+k+":")
	}

	final := yaml + "\n" + strings.Join(buffer, "\n")

	fmt.Println(final)
}

var yaml = `openapi: 3.1.0
info:
  title: Feedback API
  description: Feedback API 
  version: 1.0.0
paths:`
