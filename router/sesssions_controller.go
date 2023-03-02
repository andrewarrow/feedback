package router

type SessionsController struct {
}

func NewSessionsController(c *Context) Controller {
	sc := SessionsController{}
	return &sc
}
