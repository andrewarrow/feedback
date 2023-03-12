package router

import (
	"net/http"
)

func SetFlash(c *Context, flash string) {
	cookie := http.Cookie{}
	cookie.Path = "/"
	cookie.MaxAge = 86400 * 30
	cookie.Name = "flash"
	cookie.Value = flash
	http.SetCookie(c.Writer, &cookie)
}

func setUser(c *Context, guid string) {
	cookie := http.Cookie{}
	cookie.Path = "/"
	cookie.MaxAge = 86400 * 30
	cookie.Name = "user"
	cookie.Value = guid
	http.SetCookie(c.Writer, &cookie)
}

func removeFlash(writer http.ResponseWriter) {
	cookie := http.Cookie{}
	cookie.MaxAge = 0
	cookie.Name = "flash"
	cookie.Value = ""
	cookie.Path = "/"
	http.SetCookie(writer, &cookie)
}
