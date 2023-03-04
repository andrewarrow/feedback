package router

import "net/http"

func handleSessions(c *Context, second, third string) {
	if second == "" {
		handleSessionsIndex(c)
	} else if third != "" {
		c.notFound = true
	} else {
		if second == "new" && c.method == "GET" {
			c.SendContentInLayout("sessions_new.html", nil, 200)
			return
		}
		c.notFound = true
	}
}

func handleSessionsIndex(c *Context) {
	if c.request.Method == "POST" {
		c.ReadFormPost()
		if c.method == "DELETE" {
			DestroySession(c)
		} else if c.method == "POST" {
			CreateSession(c)
		}
		return
	}
	c.notFound = true
}

func DestroySession(c *Context) {
	cookie := http.Cookie{}
	cookie.MaxAge = 0
	cookie.Name = "user"
	cookie.Value = ""
	cookie.Path = "/"
	http.SetCookie(c.writer, &cookie)
	http.Redirect(c.writer, c.request, "/", 302)
}

func CreateSession(c *Context) {
	cookie := http.Cookie{}
	cookie.MaxAge = 86400 * 30
	cookie.Name = "user"
	cookie.Value = "123"
	cookie.Path = "/"
	http.SetCookie(c.writer, &cookie)
	http.Redirect(c.writer, c.request, "/", 302)
}

/*
type SessionsController struct {
}

func NewSessionsController() Controller {
	sc := SessionsController{}
	return &sc
}

func (sc *SessionsController) New(c *Context) {
}

func (sc *SessionsController) Index(context *Context) {
}

func (sc *SessionsController) Show(c *Context, id string) {
}

func (sc *SessionsController) CreateWithId(c *Context, id string) {
}

func (sc *SessionsController) CreateWithJson(context *Context, body string) {
}*/
