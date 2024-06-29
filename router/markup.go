package router

import (
	"io/ioutil"
)

func Markup(c *Context, second, third string) {
	if second != "" && third == "" && c.Method == "GET" {
		handleMarkupShow(c, second)
		return
	}
	c.NotFound = true
}

func handleMarkupShow(c *Context, name string) {
	c.Router.GetLiveOrCachedTemplate("form")
	asBytes, _ := ioutil.ReadFile("views/" + name)
	contentType := "text/plain"
	c.Writer.Header().Set("Content-Type", contentType)
	c.Writer.Write(asBytes)
}
