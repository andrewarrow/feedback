package router

func handleVote(c *Context, second, third string) {
	if second == "" {
		c.notFound = true
	} else if third != "" {
		c.notFound = true
	} else {
		story := FetchStory(c.db, second)
		if story == nil {
			c.writer.WriteHeader(404)
			return
		}
		c.db.Exec("update stories set points=points+1 where guid=$1", second)
		c.writer.WriteHeader(200)
	}
}
