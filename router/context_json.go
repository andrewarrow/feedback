package router

import (
	"encoding/json"
	"fmt"
)

func (c *Context) SendContentAsJson(thing any, status int) {
	asBytes, _ := json.Marshal(thing)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Header().Set("Cache-Control", "none")
	c.Writer.Header().Set("Content-Length", fmt.Sprintf("%d", len(asBytes)))
	c.Writer.WriteHeader(status)
	c.Writer.Write(asBytes)
}
