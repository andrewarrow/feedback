package router

import (
	"net/http"
)

func setFlash(c *Context, flash string) {
	cookie := http.Cookie{}
	cookie.Path = "/"
	cookie.MaxAge = 86400 * 30
	cookie.Name = "flash"
	cookie.Value = flash
	http.SetCookie(c.writer, &cookie)
}

func setUser(c *Context, guid string) {
	cookie := http.Cookie{}
	cookie.Path = "/"
	cookie.MaxAge = 86400 * 30
	cookie.Name = "user"
	cookie.Value = guid
	http.SetCookie(c.writer, &cookie)
}
