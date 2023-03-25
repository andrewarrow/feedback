package router

func handleApi(c *Context, second, third string) {
	if second != "" && third == "" {
		handleApiCall(c)
		return
	}
	c.NotFound = true
}

func handleApiCall(c *Context) {
	m := map[string]any{}
	m["test"] = []string{"hi", "there"}
	c.SendContentAsJson(m, 200)
}
