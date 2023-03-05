package router

func handleAbout(c *Context, second, third string) {
	if second == "" {
		handleAboutIndex(c)
	} else if third != "" {
		c.notFound = true
	} else {
		c.notFound = true
	}
}
func handleAboutIndex(c *Context) {
	c.SendContentInLayout("about_index.html", nil, 200)
}
