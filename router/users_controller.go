package router

import (
	"net/http"

	"github.com/andrewarrow/feedback/util"
	"golang.org/x/crypto/bcrypt"
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

func hashPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(hashedPassword)
}

func handleCreateUser(c *Context) {
	c.ReadFormValuesIntoParams("username", "password")
	message := c.ValidateCreate("user")
	returnPath := "/sessions/new"
	if message != "" {
		SetFlash(c, message)
		http.Redirect(c.Writer, c.Request, returnPath, 302)
		return
	}

	c.Params["password"] = hashPassword(c.Params["password"].(string))
	message = c.Insert("user")
	if message != "" {
		SetFlash(c, message)
		http.Redirect(c.Writer, c.Request, returnPath, 302)
		return
	}

	row := c.SelectOne("user", "where username=$1", []any{c.Params["username"]})
	guid := util.PseudoUuid()
	c.Params = map[string]any{"guid": guid, "user_id": row["id"].(int64)}
	c.Insert("cookie_token")
	setUser(c, guid)
	returnPath = "/"

	funcToRun := c.router.afterFuncToRun("user")

	if funcToRun != nil {
		funcToRun(c, guid)
	}

	http.Redirect(c.Writer, c.Request, returnPath, 302)
	return
}
