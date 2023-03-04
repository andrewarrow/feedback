package router

import (
	"net/http"

	"github.com/andrewarrow/feedback/util"
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
		body := c.request.FormValue("body")
		guid := util.PseudoUuid()
		if url != "" {
			c.db.Exec("insert into stories (title, url, guid, username) values ($1, $2, $3, $4)", title, url, guid, c.user.Username)
		} else {
			c.db.Exec("insert into stories (title, body, guid, username) values ($1, $2, $3, $4)", title, body, guid, c.user.Username)
		}
		http.Redirect(c.writer, c.request, "/", 302)
		return
	}
	c.notFound = true
}
