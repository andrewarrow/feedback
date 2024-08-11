package openapi

import (
	"strings"
)

type Endpoint struct {
	Method  string
	Returns string
	Path    string
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

	tokens = strings.Split(line, "&&")
	tokens = tokens[0 : len(tokens)-1]

	buffer := []string{}
	for _, item := range tokens {
		if strings.Contains(item, "==") {
			tokens := strings.Split(item, "==")
			thing := strings.TrimSpace(tokens[1])
			thing = thing[1 : len(thing)-1]
			buffer = append(buffer, thing)
		} else if strings.Contains(item, "!=") {
			buffer = append(buffer, "{id}")
		}
	}
	ep.Path = "/" + strings.Join(buffer, "/")

	return ep
}
