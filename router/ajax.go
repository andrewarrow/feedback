package router

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func (c *Context) SendContentForAjax(doZip bool, user map[string]any, writer http.ResponseWriter,
	filename string, contentVars any, status int) {

	t := c.router.Template.Lookup(filename)
	content := new(bytes.Buffer)
	t.Execute(content, contentVars)
	cb := content.Bytes()
	m := map[string]any{}
	m["html"] = string(cb)
	m["next"] = ""
	m["div"] = "feedback-ajax"
	asBytes, _ := json.Marshal(m)
	doZippyJson(doZip, asBytes, status, writer)
}
