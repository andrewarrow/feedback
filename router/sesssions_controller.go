package router

type SessionsController struct {
}

func NewSessionsController() Controller {
	sc := SessionsController{}
	return &sc
}

func (mc *SessionsController) Index(r *Router, context *Context) {
}

func (mc *SessionsController) Create(r *Router, context *Context) {
}
