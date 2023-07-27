package main

func controllerTemplate() string {
	t := `package app

import (
	"github.com/andrewarrow/feedback/router"
	"github.com/andrewarrow/feedback/util"
)

{{$name := index . "name"}}
{{$lower := index . "lower"}}
{{$withS := index . "with_s"}}
func Handle{{$name}}(c *router.Context, second, third string) {
	if NotLoggedIn(c) {
		return
	}
	if second == "" && third == "" && c.Method == "GET" {
		handle{{$name}}Index(c)
		return
	}
	if second == "" && third == "" && c.Method == "POST" {
		handle{{$name}}Create(c)
		return
	}
	if second != "" && third == "" && c.Method == "GET" {
		handle{{$name}}Show(c, second)
		return
	}
	if second != "" && third == "" && c.Method == "POST" {
		handle{{$name}}ShowPost(c, second)
		return
	}
	c.NotFound = true
}

func handle{{$name}}Index(c *router.Context) {
	list := c.All("{{$lower}}", "where user_id=$1 order by created_at desc", "", c.User["id"])

	colAttributes := map[int]string{}
	//colAttributes[0] = "w-1/2"

	m := map[string]any{}
	headers := []string{"name", "foo", "created"}

	params := map[string]any{}
	m["headers"] = headers
	m["cells"] = c.MakeCells(util.ToAnyArray(list), headers, params, "_{{$lower}}")
	m["col_attributes"] = colAttributes

	topVars := map[string]any{}
	send := map[string]any{}
	send["bottom"] = c.Template("table_show.html", m)
	send["top"] = c.Template("{{$withS}}_list_top.html", topVars)
	c.SendContentInLayout("generic_top_bottom.html", send, 200)
}`

	return t
}
