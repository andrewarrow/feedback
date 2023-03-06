package router

func handleFresh(c *Context, second, third string) {
	if second == "" {
		handleFreshIndex(c)
	} else if third != "" {
		c.notFound = true
	} else {
		c.notFound = true
	}
}
func handleFreshIndex(c *Context) {
	c.SendContentInLayout("welcome.html", WelcomeIndexVars(c.db, "created_at desc", ""), 200)
}
