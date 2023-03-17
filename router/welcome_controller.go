package router

func handleWelcome(c *Context, second, third string) {
	if second == "" {
		handleWelcomeIndex(c)
	} else if second != "" && third == "" {
		c.NotFound = true
	} else {
		c.NotFound = true
	}
}

func handleWelcomeIndex(c *Context) {
	c.SendContentInLayout("welcome_index.html", nil, 200)
}
