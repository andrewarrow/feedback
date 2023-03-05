package router

func handleVote(c *Context, second, third string) {
	if second == "" {
		c.notFound = true
	} else if third != "" {
		c.notFound = true
	} else {
		c.writer.WriteHeader(200)
	}
}
