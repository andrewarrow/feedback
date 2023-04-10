package router

import (
	"encoding/json"
)

func (c *Context) ReadJsonBodyIntoParams() {
	c.Params = map[string]any{}
	body := c.BodyAsString()
	json.Unmarshal([]byte(body), &c.Params)
}

func (c *Context) ReadJsonBodyAsArray() []any {
	var list []any
	body := c.BodyAsString()
	json.Unmarshal([]byte(body), &list)
	return list
}

func (c *Context) ExecuteTemplate(filename string, vars any) {
	c.router.Template.ExecuteTemplate(c.Writer, filename, vars)
}
