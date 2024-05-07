package wasm

import (
	"bytes"
	"embed"
	"encoding/json"
	"io/fs"
	"strings"
	"syscall/js"
	"text/template"

	"github.com/andrewarrow/feedback/common"
)

var EmbeddedTemplates embed.FS
var AllTemplates map[string]any
var UseLive = true

func (d *Document) RenderToId(id, name string, vars any) *Wrapper {
	div := d.ById(id)
	div.Set("innerHTML", d.Render(name, vars))
	return NewWrapper(div)
}

func (d *Document) RenderAndAppend(location, template, key, jsonString string) *Wrapper {
	var vars map[string]any
	json.Unmarshal([]byte(jsonString), &vars)
	div := d.RenderToNewDiv(template, vars[key])
	d.Id(location).Call("appendChild", div)
	return NewWrapper(div)
}

func (d *Document) RenderToNewDiv(name string, vars any) js.Value {
	newHTML := d.Render(name, vars)
	newDiv := d.Document.Call("createElement", "div")
	newDiv.Set("innerHTML", newHTML)
	return newDiv.Get("firstElementChild")
}

func (d *Document) NewTag(t, s string) *Wrapper {
	newTag := d.Document.Call("createElement", t)
	newTag.Set("innerHTML", s)
	return NewWrapper(newTag)
}

func (d *Document) Render(name string, vars any) string {
	return Render(name, vars)
}

func Render(name string, vars any) string {
	templateText := ""
	if UseLive {
		templateText = AllTemplates[name].(string)
	} else {
		templateBytes, _ := EmbeddedTemplates.ReadFile("views/" + name + ".html")
		templateText = string(templateBytes)
	}
	//fmt.Println(templateText)
	t := template.New("")
	t = t.Funcs(common.TemplateFunctions())
	if CustomFuncMap != nil {
		t = t.Funcs(*CustomFuncMap)
	}
	t, _ = t.Parse(string(templateText))
	content := new(bytes.Buffer)
	t.Execute(content, vars)
	t.ExecuteTemplate(content, name, vars)
	cb := content.Bytes()
	return string(cb)
}

func LoadAllTemplates(files []fs.DirEntry) {

	AllTemplates = map[string]any{}
	for _, item := range files {
		name := item.Name()
		tokens := strings.Split(name, ".")
		AllTemplates[tokens[0]] = DoGet("/markup/" + name)
	}

}
