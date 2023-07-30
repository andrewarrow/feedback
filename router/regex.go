package router

func (c *Context) RegexMap(s string) map[string]string {
	model := c.FindModel(s)
	regexMap := map[string]string{}
	for _, f := range model.Fields {
		regexMap[f.Name] = f.Regex
	}
	return regexMap
}
