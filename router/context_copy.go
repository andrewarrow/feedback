package router

func (c *Context) CopyContext() *Context {
	cc := Context{}
	cc.Params = map[string]any{}
	cc.Router = c.Router
	cc.Db = c.Router.Db
	return &cc
}
