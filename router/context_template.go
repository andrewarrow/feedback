package router

import (
	"bytes"
	"fmt"
)

func (c *Context) Template(name string, vars any) {
	t := c.router.Template.Lookup(name)
	if t == nil {
		return
	}
	content := new(bytes.Buffer)
	t.Execute(content, vars)
	cs := content.String()
	fmt.Println(cs)
	//template.HTML(cs)
}
