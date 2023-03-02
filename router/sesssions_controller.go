package router

type SessionsController struct {
}

func NewSessionsController() Controller {
	sc := SessionsController{}
	return &sc
}

func (sc *SessionsController) Index(context *Context) {
}

func (sc *SessionsController) Create(context *Context) {
}

func (sc *SessionsController) Show(c *Context, id string) {
}
