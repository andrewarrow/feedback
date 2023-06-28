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

func SetCookie(c *Context, name, value string) {
	cookie := http.Cookie{}
	cookie.Path = "/"
	cookie.MaxAge = 86400 * 30
	cookie.Name = name
	cookie.Value = value
	http.SetCookie(c.Writer, &cookie)
}

func GetCookie(c *Context, name string) string {
	cookie, err := c.Request.Cookie(name)
	if err != nil {
		return ""
	}
	return cookie.Value
}

func SetUser(c *Context, guid, domain string) {
	cookie := http.Cookie{}
	cookie.Path = "/"
	cookie.MaxAge = 86400 * 30
	cookie.Name = "user_v2"
	cookie.Value = guid
	if domain != "" && domain != "localhost" {
		cookie.Domain = domain
	}
	cookie.SameSite = http.SameSiteNoneMode
	//cookie.Secure=  true
	//HttpOnly= true
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
