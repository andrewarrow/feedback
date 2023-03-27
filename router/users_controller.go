package router

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/andrewarrow/feedback/models"
	"github.com/andrewarrow/feedback/prefix"
	"github.com/andrewarrow/feedback/util"
)

func handleUsers(c *Context, second, third string) {
	if second == "" {
		handleUsersIndex(c)
	} else if third != "" {
		c.NotFound = true
	} else {
		if c.Method == "GET" {
			handleUsersShow(c, second)
			return
		}
		c.NotFound = true
	}
}

func handleUsersShow(c *Context, username string) {
	c.Title = username
	u := c.router.LookupUsername(username)
	if u == nil {
		c.NotFound = true
		return
	}
	c.SendContentInLayout("users_show.html", u, 200)
}

func handleUsersIndex(c *Context) {
	if c.Method == "POST" {
		username := strings.TrimSpace(c.Request.FormValue("username"))
		password := strings.TrimSpace(c.Request.FormValue("password"))
		returnPath := "/sessions/new/"

		if len(password) < 8 {
			SetFlash(c, "password too short.")
			http.Redirect(c.Writer, c.Request, returnPath, 302)
			return
		}
		if len(password) > 255 {
			SetFlash(c, "password too long.")
			http.Redirect(c.Writer, c.Request, returnPath, 302)
			return
		}
		username = models.RemoveNonAlphanumeric(username)
		if len(username) < 2 {
			SetFlash(c, "username too short.")
			http.Redirect(c.Writer, c.Request, returnPath, 302)
			return
		}
		if len(username) > 20 {
			SetFlash(c, "username too long.")
			http.Redirect(c.Writer, c.Request, returnPath, 302)
			return
		}

		guid := util.PseudoUuid()
		sql := fmt.Sprintf("insert into %s (username, password, guid) values ($1, $2, $3)",
			prefix.Tablename("users"))
		_, err := c.Db.Exec(sql, username, password, guid)
		if err != nil {
			SetFlash(c, "username is taken.")
		} else {
			setUser(c, guid)
			returnPath = "/"
		}
		funcToRun := c.router.afterFuncToRun("user")

		if funcToRun != nil {
			funcToRun(c, guid)
		}

		http.Redirect(c.Writer, c.Request, returnPath, 302)
		return
	}
	c.NotFound = true
}
