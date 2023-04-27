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
	list := []any{""}
	params := map[string]any{}
	params["thing"] = "hi"

	colAttributes := map[int]string{}
	colAttributes[1] = "w-1/2"

	m := map[string]any{}
	headers := []string{"title", "description", "price", "something", "else", ""}
	m["headers"] = headers
	m["cells"] = c.MakeCells(list, headers, params, "_welcome_show")
	m["col_attributes"] = colAttributes

	c.SendContentInLayout("table_show.html", m, 200)
}
