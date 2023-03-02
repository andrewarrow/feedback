package router

import (
	"net/http"
)

type SessionsController struct {
}

func NewSessionsController() Controller {
	sc := SessionsController{}
	return &sc
}

func (sc *SessionsController) New(c *Context) {
	c.SendContentInLayout("sessions_new.html", nil, 200)
}

func (sc *SessionsController) Index(context *Context) {
}

func (sc *SessionsController) Show(c *Context, id string) {
}

func (sc *SessionsController) Create(c *Context, body string) {
	cookie := http.Cookie{}
	cookie.MaxAge = 86400 * 30
	cookie.Name = "user"
	cookie.Value = "123"
	http.SetCookie(c.writer, &cookie)
	http.Redirect(c.writer, c.request, "/", 301)
}

func (sc *SessionsController) CreateWithJson(context *Context, body string) {
}
