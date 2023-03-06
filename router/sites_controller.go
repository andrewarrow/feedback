package router

func handleSites(c *Context, second, third string) {
	if second == "" {
		c.notFound = true
	} else if third != "" {
		c.notFound = true
	} else {
		c.SendContentInLayout("welcome.html", WelcomeIndexVars(c.db, "created_at desc", second), 200)
	}
}
