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
	Params string `json:"params"`
}

func (r *Route) Generate(root string) string {

	buffer1 := []string{}
	buffer2 := []string{}
	names := []string{}
	params := []string{}
	logic := ""
	name := ""
	param := ""
	for _, path := range r.Paths {
		logic, name, param = handlePath(root, path.Verb, path.Second, path.Third)
		buffer1 = append(buffer1, logic)
		names = append(names, name)
		params = append(params, param)
	}
	for i, name := range names {
		logic := handleFuncs(name, params[i])
		buffer2 = append(buffer2, logic)
	}
	content := strings.Join(buffer1, "\n")
	top := handleWrapper(root, content)

	funcs := strings.Join(buffer2, "\n")

	return top + "\n\n" + funcs
}

func handleWrapper(root, templateContent string) string {
	c := `func Handle{{ index . "name" }}(c *router.Context, second, third string) {
  if len(c.User) == 0 {
    c.SendContentAsJsonMessage("auth not set", 401)
    return
  }
{{ index . "content" }}
  c.NotFound = true
}
`
	m := map[string]string{"name": util.ToCamelCase(root), "content": templateContent}
	t, _ := template.New("c").Parse(c)
	content := new(bytes.Buffer)
	t.Execute(content, m)
	logic := content.String()
	return logic
}

func handlePath(root, verb, second, third string) (string, string, string) {
	c := `  if second {{ index . "second_eq" }} && third {{ index . "third_eq" }} && c.Method == "{{ index . "method" }}" {
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

	return logic, name, params

}

func handleFuncs(name, params string) string {
	c := `func handle{{ index . "name" }}(c *router.Context{{ index . "params" }}) {
  c.ReadJsonBodyIntoParams()
  //c.SendContentAsJsonMessage("no guid", 422)
  c.SendRowAsJson("", nil)
}`

	paramList := ""
	if params != "" {
		paramList = params + " string"
	}
	m := map[string]string{"name": name, "params": paramList}
	t, _ := template.New("c").Parse(c)
	content := new(bytes.Buffer)
	t.Execute(content, m)
	logic := content.String()

	return logic

}
