package router

import "strings"

func (c *Context) ReadFormValuesIntoParams(list ...string) {
	c.Params = map[string]any{}
	for _, name := range list {
		val := strings.TrimSpace(c.Request.FormValue(name))
		c.Params[name] = val
	}
}

func (c *Context) ReadMultipleFormValues(list ...string) {
	c.Params = map[string]any{}
	for _, name := range list {
		selectedValues := c.Request.PostForm[name]
		buffer := []string{}
		for _, item := range selectedValues {
			buffer = append(buffer, strings.TrimSpace(item))
		}

		val := strings.Join(buffer, ",")
		c.Params[name] = val
	}
}
