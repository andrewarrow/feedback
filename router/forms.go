package router

import "strings"

func (c *Context) ReadFormValuesIntoParams(list ...string) {
	c.Params = map[string]any{}
	for _, name := range list {
		val := strings.TrimSpace(c.Request.FormValue(name))
		c.Params[name] = val
	}
}
