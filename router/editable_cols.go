package router

func GetEditableCols(c *Context, modelString string) ([]string, map[string]string) {
	model := c.FindModel(modelString)
	cols := []string{}
	editable := map[string]string{}
	for _, f := range model.Fields {
		if f.Flavor == "editable" {
			editable[f.Name] = "string"
		} else if f.Flavor == "select" {
			editable[f.Name] = "select"
		}
		cols = append(cols, f.Name)
	}
	return cols, editable
}
