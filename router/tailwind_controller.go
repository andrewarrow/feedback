package router

func handleTailwind(c *Context, second, third string) {
	c.Layout = "tailwind_layout.html"
	if second == "" {
		handleTailwindIndex(c)
		return
	}
	c.NotFound = true
}

func handleTailwindIndex(c *Context) {
	c.SendContentInLayout("tailwind_index.html", nil, 200)
}
