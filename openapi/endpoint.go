package openapi

import "strings"

type Endpoint struct {
	Method string
}

var verbs = []string{"GET", "POST", "DELETE", "PATCH", "PUT"}

func NewEndpoint(comment, line string) Endpoint {
	ep := Endpoint{}

	for _, verb := range verbs {
		if strings.Contains(line, verb) {
			ep.Method = verb
			break
		}
	}

	return ep
}
