package router

func handleBuildings(c *Context, second, third string) {
	if second == "" {
		handleBuildingsIndex(c)
	} else if third != "" {
		c.notFound = true
	} else {
		handleBuildingsShow(c)
	}
}

type BuildingVars struct {
	Rows []*Building
}

func handleBuildingsIndex(c *Context) {
	vars := BuildingVars{}
	vars.Rows = FetchBuildings(c.db)
	c.SendContentInLayout("buildings_index.html", &vars, 200)
}

func handleBuildingsShow(c *Context) {
	c.SendContentInLayout("buildings_show.html", nil, 200)
}
