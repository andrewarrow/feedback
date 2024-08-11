package openapi

import (
	"fmt"
	"strings"
)

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

	tokens = strings.Split(line, "&&")
	tokens = tokens[0 : len(tokens)-1]

	for _, item := range tokens {
		if strings.Contains(item, "==") {
			tokens := strings.Split(item, "==")
			fmt.Println(tokens[1])
		} else if strings.Contains(item, "!=") {
			tokens := strings.Split(item, "!=")
			fmt.Println(tokens[1])
		}
	}

	// second == "user" && third == "charge-history" && foo
	// second == "user" && third != "" && bar

	return ep
}
