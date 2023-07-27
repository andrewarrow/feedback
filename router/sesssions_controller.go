package router

import (
	"net/http"
	"os"

	"github.com/andrewarrow/feedback/util"
	"golang.org/x/crypto/bcrypt"
)

func NotLoggedIn(c *Context) bool {
	if len(c.User) == 0 {
		http.Redirect(c.Writer, c.Request, "/"+c.Router.Prefix, 302)
		return true
	}
	return false
}

func handleSessions(c *Context, second, third string) {
	if second == "" {
		handleSessionsIndex(c)
	} else if third != "" {
		c.NotFound = true
	} else {
		if second == "new" && c.Method == "GET" {
			c.SendContentInLayout("sessions_new.html", nil, 200)
			return
		}
		c.NotFound = true
	}
}

func handleSessionsIndex(c *Context) {
	if c.Method == "DELETE" {
		DestroySession(c)
	} else if c.Method == "POST" {
		CreateSession(c)
	} else {
		c.NotFound = true
	}
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateSession(c *Context) {
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")
	row := c.SelectOne("user", "where username=$1", []any{username})

	returnPath := "/"
	cookie := http.Cookie{}
	cookie.Path = "/"
	if len(row) > 0 && checkPasswordHash(password, row["password"].(string)) {

		guid := util.PseudoUuid()
		c.Params = map[string]any{"guid": guid, "user_id": row["id"].(int64)}
		c.Insert("cookie_token")
		SetUser(c, guid, os.Getenv("COOKIE_DOMAIN"))
	} else {
		cookie.MaxAge = 86400 * 30
		cookie.Name = "flash"
		cookie.Value = "username not found."
		returnPath = "/sessions/new"
	}
	http.SetCookie(c.Writer, &cookie)
	http.Redirect(c.Writer, c.Request, returnPath, 302)
}

func DestroySession(c *Context) {
	id := c.User["id"].(int64)
	c.Delete("cookie_token", "user_id", id)

	cookie := http.Cookie{}
	cookie.MaxAge = 0
	cookie.Name = "user"
	cookie.Value = ""
	cookie.Path = "/"
	http.SetCookie(c.Writer, &cookie)
	http.Redirect(c.Writer, c.Request, "/", 302)
}
