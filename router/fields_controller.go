package router

func handleFields(c *Context, second, third string) {
	c.Layout = "models_layout.html"
	if c.User == nil {
		c.UserRequired = true
		return
	}
	if IsAdmin(c.User) == false {
		c.NotFound = true
		return
	}
	if second != "" && third != "" {
		handleFieldsShow(c)
		return
	}
	c.NotFound = true
}

func handleFieldsShow(c *Context) {
	c.SendContentInLayout("fields_show.html", nil, 200)
}
