package router

import (
	"fmt"
	"net/http"
)

func handleStories(c *Context, second, third string) {
	if second == "" {
		handleStoriesIndex(c)
	} else if third != "" {
		c.notFound = true
	} else {
		if second == "new" && c.method == "GET" {
			c.SendContentInLayout("stories_new.html", nil, 200)
			return
		}
		c.notFound = true
	}
}

func handleStoriesIndex(c *Context) {
	if c.method == "POST" {
		title := c.request.FormValue("title")
		url := c.request.FormValue("url")
		text := c.request.FormValue("text")
		fmt.Println(title, url, text)
		http.Redirect(c.writer, c.request, "/", 302)
		return
	}
	c.notFound = true
}
