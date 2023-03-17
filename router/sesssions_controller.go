package router

import (
	"fmt"
	"net/http"

	"github.com/andrewarrow/feedback/prefix"
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

func CreateSession(c *Context) {
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")
	sql := fmt.Sprintf("SELECT * FROM %s where username=$1 and password=$2", prefix.Tablename("users"))
	rows, err := c.Db.Queryx(sql, username, password)
	if err != nil {
		return
	}
	defer rows.Close()

	m := make(map[string]any)
	rows.Next()
	rows.MapScan(m)

	returnPath := "/"
	cookie := http.Cookie{}
	cookie.Path = "/"
	if len(m) > 0 {
		cookie.MaxAge = 86400 * 30
		cookie.Name = "user"
		cookie.Value = fmt.Sprintf("%s", m["guid"])
	} else {
		cookie.MaxAge = 86400 * 30
		cookie.Name = "flash"
		cookie.Value = "username not found."
		returnPath = "/sessions/new/"
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
