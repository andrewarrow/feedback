package main

func controllerTemplate() string {
	t := `package app

import (
	"github.com/andrewarrow/feedback/router"
)

{{$name := index . "name"}}
{{$lower := index . "lower"}}
{{$withS := index . "with_s"}}
func Handle{{$name}}(c *router.Context, second, third string) {
	if router.NotLoggedIn(c) {
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
	//list := c.All("{{$lower}}", "where user_id=$1 order by created_at desc", "", c.User["id"])

	send := map[string]any{}
	c.SendContentInLayout(".html", send, 200)
}`

	return t
}
