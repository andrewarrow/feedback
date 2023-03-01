package router

func (r *Router) SessionsResource(c *Context) {
	r.SendContentInLayout(c.writer, "sessions_index.html", nil, 200)
}
