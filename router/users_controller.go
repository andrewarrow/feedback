package router

import (
	"net/http"
	"os"

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
	u := c.Router.LookupUsername(username)
	if u == nil {
		c.NotFound = true
		return
	}
	c.SendContentInLayout("users_show.html", u, 200)
}

func HashPassword(password string) string {
	// between 4 and 31, Each increment of the cost parameter doubles the amount of time needed to compute the hash, making it more computationally expensive to brute-force the hash.
	// With a cost of 4, the bcrypt hash would take roughly 2 milliseconds to generate on a modern CPU. Assuming an attacker can generate hashes at a similar rate, they would be able to try around 500 passwords per second on an average laptop. At a rate of 500 passwords per second, it would take approximately 12 days to try all possible passwords.
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

	c.Params["password"] = HashPassword(c.Params["password"].(string))
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
	SetUser(c, guid, os.Getenv("COOKIE_DOMAIN"))
	returnPath = "/"

	funcToRun := c.Router.afterFuncToRun("user")

	if funcToRun != nil {
		funcToRun(c, guid)
	}

	http.Redirect(c.Writer, c.Request, returnPath, 302)
	return
}

func HandleCreateUserAutoForm(c *Context, username string) string {
	c.ReadJsonBodyIntoParams()
	c.Params["username"] = username
	if username == "" {
		c.Params["username"] = c.Params["email"]
	}
	password, _ := c.Params["password"].(string)
	message := c.ValidateCreate("user")
	send := map[string]any{}
	if message != "" {
		send["error"] = message
		c.SendContentAsJson(send, 422)
		return ""
	}
	c.Params["password"] = HashPassword(password)
	message = c.Insert("user")
	if message != "" {
		send["error"] = message
		c.SendContentAsJson(send, 422)
		return ""
	}
	row := c.One("user", "where username=$1", c.Params["email"])
	guid := util.PseudoUuid()
	c.Params = map[string]any{"guid": guid, "user_id": row["id"]}
	c.Insert("cookie_token")
	SetUser(c, guid, os.Getenv("COOKIE_DOMAIN"))
	c.SendContentAsJson("ok", 200)
	return guid
}
