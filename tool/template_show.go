package main

func showTemplate() string {

	t := `package app

{{$name := index . "name"}}
{{$lower := index . "lower"}}
{{$withS := index . "with_s"}}

import (
	"net/http"

	"github.com/andrewarrow/feedback/router"
)

func handle{{$name}}ShowPost(c *router.Context, guid string) {
	c.ReadFormValuesIntoParams("file")
	returnPath := "/"

	//c.ValidateCreate("{{$lower}}")
	message := c.Update("{{$lower}}", "where guid=", guid)
	if message != "" {
		router.SetFlash(c, message)
		http.Redirect(c.Writer, c.Request, returnPath+"/"+guid, 302)
		return
	}
	http.Redirect(c.Writer, c.Request, returnPath, 302)
}

func handle{{$name}}Show(c *router.Context, guid string) {
	item := c.One("{{$lower}}", "where guid=$1", guid)
	send := map[string]any{}
	send["item"] = item
	c.SendContentInLayout(".html", send, 200)
}`

	return t
}
