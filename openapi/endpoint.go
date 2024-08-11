package openapi

import "strings"

type Endpoint struct {
	Method  string
	Returns string
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

	tokens := strings.Split(comment, " ")
	ep.Returns = tokens[len(tokens)-1]

	return ep
}
