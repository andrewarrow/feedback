package router

import (
	"net/http"
	"strings"

	"github.com/andrewarrow/feedback/models"
	"github.com/andrewarrow/feedback/util"
)

func handleUsers(c *Context, second, third string) {
	if second == "" {
		handleUsersIndex(c)
	} else if third != "" {
		c.notFound = true
	} else {
		if c.method == "GET" {
			handleUsersShow(c, second)
			return
		}
		c.notFound = true
	}
}

func handleUsersShow(c *Context, username string) {
	c.title = username
	u := c.router.LookupUsername(username)
	if u == nil {
		c.notFound = true
		return
	}
	c.SendContentInLayout("users_show.html", u, 200)
}

func handleUsersIndex(c *Context) {
	if c.method == "POST" {
		username := strings.TrimSpace(c.request.FormValue("username"))
		password := strings.TrimSpace(c.request.FormValue("password"))
		returnPath := "/sessions/new/"

		if len(password) < 8 {
			setFlash(c, "password too short.")
			http.Redirect(c.writer, c.request, returnPath, 302)
			return
		}
		if len(password) > 255 {
			setFlash(c, "password too long.")
			http.Redirect(c.writer, c.request, returnPath, 302)
			return
		}
		username = models.RemoveNonAlphanumeric(username)
		if len(username) < 2 {
			setFlash(c, "username too short.")
			http.Redirect(c.writer, c.request, returnPath, 302)
			return
		}
		if len(username) > 20 {
			setFlash(c, "username too long.")
			http.Redirect(c.writer, c.request, returnPath, 302)
			return
		}

		guid := util.PseudoUuid()
		_, err := c.db.Exec("insert into users (username, password, guid) values ($1, $2, $3)", username, password, guid)
		if err != nil {
			setFlash(c, "username is taken.")
		} else {
			setUser(c, guid)
			returnPath = "/"
		}
		http.Redirect(c.writer, c.request, returnPath, 302)
		return
	}
	c.notFound = true
}
