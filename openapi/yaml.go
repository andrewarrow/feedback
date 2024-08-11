package openapi

import (
	"fmt"
	"strings"
)

func MakeYaml(eps []Endpoint) {

	buffer := []string{}

	for _, ep := range eps {
		buffer = append(buffer, "  "+ep.Path+":")
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
