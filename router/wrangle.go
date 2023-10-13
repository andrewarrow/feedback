package router

func (c *Context) Wrangle() *Context {
	newContext := c.Router.ToContext()
	newContext.Db = c.Router.WrangleDb
	return newContext
}
