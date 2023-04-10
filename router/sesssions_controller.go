package router

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

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
	if checkPasswordHash(password, row["password"].(string)) {
		cookie.MaxAge = 86400 * 30
		cookie.Name = "user"
		cookie.Value = row["guid"].(string)
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
	cookie := http.Cookie{}
	cookie.MaxAge = 0
	cookie.Name = "user"
	cookie.Value = ""
	cookie.Path = "/"
	http.SetCookie(c.Writer, &cookie)
	http.Redirect(c.Writer, c.Request, "/", 302)
}
