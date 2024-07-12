package app

import (
	"github.com/andrewarrow/feedback/router"
)

func Core(c *router.Context, second, third string) {
	if second == "start" && third == "" && c.Method == "GET" {
		handleStart(c)
		return
	}
	if second == "about-us" && third == "" && c.Method == "GET" {
		handleAboutUs(c)
		return
	}
	if second == "privacy" && third == "" && c.Method == "GET" {
		handlePrivacy(c)
		return
	}
	if second == "terms" && third == "" && c.Method == "GET" {
		handleTerms(c)
		return
	}
	if second == "register" && third == "" && c.Method == "GET" {
		handleRegister(c)
		return
	}
	if second == "login" && third == "" && c.Method == "GET" {
		handleLogin(c)
		return
	}
	if second == "register" && third == "" && c.Method == "POST" {
		router.HandleCreateUserAutoForm(c, "")
		return
	}
	if second == "login" && third == "" && c.Method == "POST" {
		router.HandleCreateSessionAutoForm(c)
		return
	}
	if second == "logout" && third == "" && c.Method == "DELETE" {
		router.DestroySession(c)
		return
	}
	c.NotFound = true
}

func handleIndex(c *router.Context) {
	send := map[string]any{}
	c.SendContentInLayout("welcome.html", send, 200)
}

func handleRegister(c *router.Context) {
	send := map[string]any{}
	c.SendContentInLayout("register.html", send, 200)
}
func handleLogin(c *router.Context) {
	send := map[string]any{}
	c.SendContentInLayout("login.html", send, 200)
}

func handlePrivacy(c *router.Context) {
	send := map[string]any{}
	c.SendContentInLayout("privacy.html", send, 200)
}
func handleTerms(c *router.Context) {
	send := map[string]any{}
	c.SendContentInLayout("terms.html", send, 200)
}
func handleAboutUs(c *router.Context) {
	send := map[string]any{}
	c.SendContentInLayout("about_us.html", send, 200)
}
func handleStart(c *router.Context) {
	send := map[string]any{}
	c.SendContentInLayout("start.html", send, 200)
}
