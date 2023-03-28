package router

import "strings"

func (c *Context) ReadFormValuesIntoParams(list ...string) {
	for _, name := range list {
		val := strings.TrimSpace(c.Request.FormValue(name))
		c.Params[name] = val
	}
}
