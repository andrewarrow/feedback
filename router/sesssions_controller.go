package router

func handleSessions(c *Context, second, third string) {
}

/*
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

func (sc *SessionsController) CreateWithId(c *Context, id string) {
}

func (sc *SessionsController) Create(c *Context) {
	cookie := http.Cookie{}
	cookie.MaxAge = 86400 * 30
	cookie.Name = "user"
	cookie.Value = "123"
	cookie.Path = "/"
	http.SetCookie(c.writer, &cookie)
	http.Redirect(c.writer, c.request, "/", 302)
}

func (sc *SessionsController) Destroy(c *Context) {
	cookie := http.Cookie{}
	cookie.MaxAge = 0
	cookie.Name = "user"
	cookie.Value = ""
	cookie.Path = "/"
	http.SetCookie(c.writer, &cookie)
	http.Redirect(c.writer, c.request, "/", 302)
}

func (sc *SessionsController) CreateWithJson(context *Context, body string) {
}*/
