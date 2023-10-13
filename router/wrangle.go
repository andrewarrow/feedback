package router

func (c *Context) Wrangle(s *FeedbackSite) *Context {
	r := Router{}
	r.Site = s
	r.Db = c.Router.WrangleDb
	newContext := r.ToContext()
	newContext.Db = c.Router.WrangleDb
	return newContext
}
