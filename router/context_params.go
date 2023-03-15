package router

import "encoding/json"

func (c *Context) ReadBodyIntoJson() map[string]any {
	body := c.BodyAsString()
	var params map[string]any
	json.Unmarshal([]byte(body), &params)
	return params
}
