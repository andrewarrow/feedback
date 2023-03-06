package router

import "net/http"

func DestroyFlash(writer http.ResponseWriter) {
	cookie := http.Cookie{}
	cookie.MaxAge = 0
	cookie.Name = "flash"
	cookie.Value = ""
	cookie.Path = "/"
	http.SetCookie(writer, &cookie)
}
