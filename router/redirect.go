package router

import "net/http"

func Redirect(c *Context, path string) {
	http.Redirect(c.Writer, c.Request, path, 302)
}
