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
	m["next"] = c.LayoutMap["ajax_next"]
	m["table_large"] = c.LayoutMap["table_large"]
	m["table_small"] = c.LayoutMap["table_small"]
	m["table_large_div"] = c.LayoutMap["table_large_div"]
	m["table_small_div"] = c.LayoutMap["table_small_div"]
	m["div"] = c.LayoutMap["ajax_div"]
	if m["div"] == nil {
		m["div"] = "feedback-ajax"
	}
	//fmt.Println(m["div"], m["next"])
	asBytes, _ := json.Marshal(m)
	doZippyJson(doZip, asBytes, status, writer)
}
