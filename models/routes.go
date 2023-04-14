package models

import (
	"bytes"
	"strings"
	"text/template"
)

type Route struct {
	Root  string  `json:"root"`
	Paths []*Path `json:"paths"`
}

type Path struct {
	Verb   string `json:"verb"`
	Second string `json:"second"`
	Third  string `json:"third"`
}

func (r *Route) Generate() string {

	buffer := []string{}
	for _, path := range r.Paths {
		logic := handlePath(path.Verb, path.Second, path.Third)
		buffer = append(buffer, logic)
	}

	return strings.Join(buffer, "\n")
}

func handlePath(verb, second, third string) string {
	c := `if second {{ index . "second_eq" }} "" && third {{ index . "third_eq" }} "" && c.Method == "{{ index . "method" }}" {
    handle{{ index . "name" }}(c, {{ index . "params" }})
    return
  }
`
	m := map[string]string{"name": "Foo", "params": "second",
		"second_eq": "!=", "third_eq": "=="}
	t, _ := template.New("c").Parse(c)
	content := new(bytes.Buffer)
	t.Execute(content, m)
	logic := content.String()

	return logic

}
