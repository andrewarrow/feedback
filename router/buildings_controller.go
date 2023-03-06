package router

func handleBuildings(c *Context, second, third string) {
	if second == "" {
		handleBuildingsIndex(c)
	} else if third != "" {
		c.notFound = true
	} else {
		c.notFound = true
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
