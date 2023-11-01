package router

import (
	"bytes"
	"html/template"
)

func (c *Context) Template(name string, vars any) template.HTML {
	t := c.Router.GetLiveOrCachedTemplate(name)
	if t == nil {
		return template.HTML("")
	}
	content := new(bytes.Buffer)
	t.Execute(content, vars)
	cs := content.String()
	return template.HTML(cs)
}
