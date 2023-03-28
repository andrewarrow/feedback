package router

import (
	"net/http"
)

func handleUsers(c *Context, second, third string) {
	if second != "" && third == "" && c.Method == "GET" {
		handleUsersShow(c, second)
		return
	}
	if second == "" && third == "" && c.Method == "POST" {
		handleCreateUser(c)
		return
	}
	c.NotFound = true
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

func handleCreateUser(c *Context) {
	c.ReadFormValuesIntoParams("username", "password")
	message := c.Validate("user")
	returnPath := "/sessions/new"
	if message != "" {
		SetFlash(c, message)
		http.Redirect(c.Writer, c.Request, returnPath, 302)
		return
	}

	message = c.Insert("user")
	if message != "" {
		SetFlash(c, message)
		http.Redirect(c.Writer, c.Request, returnPath, 302)
		return
	}

	guid := c.Params["guid"].(string)
	setUser(c, guid)
	returnPath = "/"

	funcToRun := c.router.afterFuncToRun("user")

	if funcToRun != nil {
		funcToRun(c, guid)
	}

	http.Redirect(c.Writer, c.Request, returnPath, 302)
	return
}
