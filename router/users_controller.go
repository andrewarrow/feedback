package router

import (
	"net/http"

	"github.com/andrewarrow/feedback/util"
)

func handleUsers(c *Context, second, third string) {
	if second == "" {
		handleUsersIndex(c)
	} else if third != "" {
		c.notFound = true
	} else {
		c.notFound = true
	}
}

func handleUsersIndex(c *Context) {
	if c.method == "POST" {
		username := c.request.FormValue("username")
		password := c.request.FormValue("password")
		guid := util.PseudoUuid()
		_, err := c.db.Exec("insert into users (username, password, guid) values ($1, $2, $3)", username, password, guid)
		returnPath := "/"
		if err != nil {
			setFlash(c, "username is taken.")
			returnPath = "/sessions/new/"
		} else {
			setUser(c, guid)
		}
		http.Redirect(c.writer, c.request, returnPath, 302)
		return
	}
	c.notFound = true
}
