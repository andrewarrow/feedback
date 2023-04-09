package router

func handleDashboard(c *Context, second, third string) {
	c.Layout = "admin_layout.html"
	if second == "" {
		handleTailwindIndex(c)
		return
	}
	c.NotFound = true
}

func handleTailwindIndex(c *Context) {
	c.SendContentInLayout("dashboard_index.html", nil, 200)
}
