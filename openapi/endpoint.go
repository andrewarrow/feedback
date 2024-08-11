package openapi

import (
	"strings"
)

type Endpoint struct {
	LowerVerb string
	Method    string
	Returns   string
	Path      string
	HasId     bool
	Params    []Param
	LastFunc  string
}

var verbs = []string{"GET", "POST", "DELETE", "PATCH", "PUT"}

func NewEndpoint(comment, line, lastFunc string) Endpoint {
	ep := Endpoint{}
	ep.LastFunc = lastFunc

	for _, verb := range verbs {
		if strings.Contains(line, verb) {
			ep.Method = verb
			ep.LowerVerb = strings.ToLower(verb)
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
			if thing != "" {
				buffer = append(buffer, thing)
			}
		} else if strings.Contains(item, "!=") {
			buffer = append(buffer, "{id}")
			ep.HasId = true
		}
	}
	ep.Path = "/" + strings.Join(buffer, "/")
	if ep.Path == "/" {
		ep.Path = ""
	}

	if ep.Method == "POST" {
		p1 := Param{"email", "string"}
		p2 := Param{"latitude", "number"}
		ep.Params = []Param{p1, p2}
	}

	return ep
}
