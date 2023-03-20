package router

import "github.com/andrewarrow/feedback/stats"

func handleStats(c *Context, second, third string) {
	c.Layout = "models_layout.html"
	if c.User == nil {
		c.UserRequired = true
		return
	}
	if IsAdmin(c.User) == false {
		c.NotFound = true
		return
	}
	if second == "" {
		handleStatsIndex(c)
		return
	}
	c.NotFound = true
}

func handleStatsIndex(c *Context) {
	c.SendContentInLayout("stats_index.html", stats.Hits, 200)
}
