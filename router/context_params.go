package router

import "encoding/json"

func (c *Context) ReadBodyIntoJson() map[string]any {
	body := c.BodyAsString()
	var params map[string]any
	json.Unmarshal([]byte(body), &params)
	return params
}

func (c *Context) ExecuteTemplate(filename string, vars any) {
	c.router.Template.ExecuteTemplate(c.Writer, filename, vars)
}
