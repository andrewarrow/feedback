package router

import (
	"fmt"
	"net/http"
)

func HandleGoogle(c *Context, second, third string) {
	if second == "login" && third == "" && c.Method == "POST" {
		handleGoogleLogin(c)
		return
	}
	c.NotFound = true
}

func handleGoogleLogin(c *Context) {
	c.ReadJsonBodyIntoParams()
	fmt.Println(c.Params)
	returnPath := "/"
	SetFlash(c, "hi there")
	http.Redirect(c.Writer, c.Request, returnPath, 302)
}
