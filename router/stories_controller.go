package router

func handleStories(c *Context, second, third string) {
	if second == "" {
		c.notFound = true
	} else if third != "" {
		c.notFound = true
	} else {
		if second == "new" && c.method == "GET" {
			c.SendContentInLayout("stories_new.html", nil, 200)
			return
		}
		c.notFound = true
	}
}
