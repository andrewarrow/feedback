package router

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

func (sc *SessionsController) Create(context *Context, body string) {
}

func (sc *SessionsController) CreateWithJson(context *Context, body string) {
}
