package router

type SessionsController struct {
}

func NewSessionsController() Controller {
	sc := SessionsController{}
	return &sc
}

func (mc *SessionsController) Index(context *Context) {
}

func (mc *SessionsController) Create(context *Context) {
}
