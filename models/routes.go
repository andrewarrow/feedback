package models

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/andrewarrow/feedback/util"
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

func (r *Route) Generate(root string) string {

	buffer := []string{}
	for _, path := range r.Paths {
		logic := handlePath(root, path.Verb, path.Second, path.Third)
		buffer = append(buffer, logic)
	}
	for _, path := range r.Paths {
		logic := handleFuncs(path.Verb)
		buffer = append(buffer, logic)
	}

	return strings.Join(buffer, "\n")
}

func handlePath(root, verb, second, third string) string {
	c := `if second {{ index . "second_eq" }} && third {{ index . "third_eq" }} && c.Method == "{{ index . "method" }}" {
    handle{{ index . "name" }}(c{{ index . "params" }})
    return
  }
`

	flavor := ""
	if third == "" && second == "" {
		flavor = "root"
	} else if third == "" && second == "*" {
		flavor = "second"
	} else if third == "*" {
		flavor = "third"
	}

	second_eq := ""
	third_eq := ""
	empty := `""`
	q := `"`
	params := ""
	if flavor == "root" {
		second_eq = "== " + empty
		third_eq = "== " + empty
	} else if flavor == "second" {
		second_eq = "!= " + empty
		third_eq = "== " + empty
		params = ", second"
	} else if flavor == "third" {
		if second == "*" {
			second_eq = "!= " + empty
			params = ", second, third"
		} else {
			second_eq = "== " + fmt.Sprintf("%s%s%s", q, second, q)
			params = ", third"
		}
		third_eq = "!= " + empty
	}

	name := fmt.Sprintf("%s%s%s", util.ToCamelCase(root),
		util.ToCamelCase(strings.ToLower(verb)),
		util.ToCamelCase(flavor))
	m := map[string]string{"name": name, "params": params,
		"second_eq": second_eq, "third_eq": third_eq, "method": verb}
	t, _ := template.New("c").Parse(c)
	content := new(bytes.Buffer)
	t.Execute(content, m)
	logic := content.String()

	return logic

}

func handleFuncs(verb string) string {
	c := `func {{ index . "name" }}(c *router.Context{{ index . "params" }}) {
  c.ReadJsonBodyIntoParams()
  //c.SendContentAsJsonMessage("no guid", 422)
  c.SendRowAsJson("", nil)
}`

	params := ""
	m := map[string]string{"name": "foo", "params": params}
	t, _ := template.New("c").Parse(c)
	content := new(bytes.Buffer)
	t.Execute(content, m)
	logic := content.String()

	return logic

}
