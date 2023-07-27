package main

func showTemplate() string {

	t := `package app

{{$name := index . "name"}}
{{$lower := index . "lower"}}
{{$withS := index . "with_s"}}

import (
	"net/http"

	"github.com/andrewarrow/feedback/router"
	"github.com/andrewarrow/feedback/util"
)

func handle{{$name}}ShowPost(c *router.Context, guid string) {
	cols, editable := router.GetEditableCols(c, "{{$lower}}")
	list := []string{}
	for _, item := range cols {
		if router.IsEditable(item, editable) == false {
			continue
		}
		list = append(list, item)
	}
	list = append(list, "submit")
	c.ReadFormValuesIntoParams(list...)
	submit := c.Params["submit"].(string)
	if submit != "save" {
		//handleFooCreate(c, guid)
		return
	}

	c.ValidateUpdate("{{$lower}}")
	message := c.ValidateUpdate("{{$lower}}")
	returnPath := "/{{$withS}}"
	if message != "" {
		router.SetFlash(c, message)
		http.Redirect(c.Writer, c.Request, returnPath+"/"+guid, 302)
		return
	}
	message = c.Update("{{$lower}}", "where guid=", guid)
	if message != "" {
		router.SetFlash(c, message)
		http.Redirect(c.Writer, c.Request, returnPath+"/"+guid, 302)
		return
	}
	http.Redirect(c.Writer, c.Request, returnPath, 302)
}

func handle{{$name}}Show(c *router.Context, guid string) {
	item := c.One("{{$lower}}", "where guid=$1", guid)
	regexMap := map[string]string{}
	cols, editable := router.GetEditableCols(c, "{{$lower}}")
	//cols = append(cols, "save")
	//editable["save"] = "save"

	colAttributes := map[int]string{}
	colAttributes[1] = "w-3/4"

	m := map[string]any{}
	headers := []string{"field", "value"}

	params := map[string]any{}
	params["item"] = item
	params["editable"] = editable
	params["regex_map"] = regexMap
	m["headers"] = headers
	m["cells"] = c.MakeCells(util.ToAny(cols), headers, params, "_{{$lower}}_show")
	m["col_attributes"] = colAttributes
	m["save"] = true
	topVars := map[string]any{}
	topVars["name"] = item["name"]
	topVars["guid"] = guid
	send := map[string]any{}
	send["bottom"] = c.Template("table_show.html", m)
	send["top"] = c.Template("{{$withS}}_top.html", topVars)

	c.SendContentInLayout("generic_top_bottom.html", send, 200)
}`

	return t
}
