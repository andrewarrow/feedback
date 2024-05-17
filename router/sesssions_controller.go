package router

import (
	"net/http"
	"os"

	"github.com/andrewarrow/feedback/util"
	"golang.org/x/crypto/bcrypt"
)

func NotLoggedIn(c *Context) bool {
	if len(c.User) == 0 {
		path := c.Request.URL.Path
		SetCookie(c, "desired_path", path)
		http.Redirect(c.Writer, c.Request, "/"+c.Router.NotLoggedInPath, 302)
		return true
	}
	return false
}

func HandleSessions(c *Context, second, third string) {
	if second == "" {
		handleSessionsIndex(c)
	} else if third != "" {
		c.NotFound = true
	} else {
		if second == "new" && c.Method == "GET" {
			m := map[string]any{}
			m["client_id"] = os.Getenv("GOOGLE_ID")
			c.SendContentInLayout("sessions_new.html", m, 200)
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
func HandleCreateSessionAutoForm(c *Context) {
	c.ReadJsonBodyIntoParams()
	email, _ := c.Params["email"].(string)
	password, _ := c.Params["password"].(string)
	row := c.One("user", "where email=$1", email)
	if len(row) > 0 && checkPasswordHash(password, row["password"].(string)) {

		guid := util.PseudoUuid()
		c.Params = map[string]any{"guid": guid, "user_id": row["id"]}
		c.Insert("cookie_token")
		SetUser(c, guid, os.Getenv("COOKIE_DOMAIN"))
		c.SendContentAsJson("ok", 200)
		return
	}
	c.SendContentAsJson("error", 422)
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
