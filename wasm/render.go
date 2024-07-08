package wasm

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"strings"
	"syscall/js"
	"text/template"

	"github.com/andrewarrow/feedback/common"
)

var EmbeddedTemplates embed.FS
var AllTemplates map[string]any
var UseLive = true
var NamedTemplates *template.Template

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

func LoadTemplates(tf template.FuncMap) *template.Template {
	t := template.New("")
	t = t.Funcs(tf)

	templateFiles, _ := EmbeddedTemplates.ReadDir("views")
	for _, file := range templateFiles {
		name := file.Name()
		tokens := strings.Split(name, ".")
		name = tokens[0]
		fileContents, _ := EmbeddedTemplates.ReadFile("views/" + file.Name())

		_, err := t.New(name).Parse(string(fileContents))
		if err != nil {
			fmt.Println(file.Name(), err)
		}
	}
	return t
}
func LoadLiveTemplates(tf template.FuncMap) *template.Template {
	t := template.New("")
	t = t.Funcs(tf)

	for name, v := range AllTemplates {
		_, err := t.New(name).Parse(v.(string))
		if err != nil {
			fmt.Println(name, err)
		}
	}

	return t
}

func Render(name string, vars any) string {
	if NamedTemplates == nil {
		if UseLive {
			NamedTemplates = LoadLiveTemplates(common.TemplateFunctions())
		} else {
			NamedTemplates = LoadTemplates(common.TemplateFunctions())
		}
	}
	fmt.Println(name, NamedTemplates)
	t := NamedTemplates.Lookup(name)
	content := new(bytes.Buffer)
	t.Execute(content, vars)
	//t.ExecuteTemplate(content, name, vars)
	cb := content.Bytes()
	return string(cb)
}

func LoadAllTemplates(files []fs.DirEntry) {

	AllTemplates = map[string]any{}
	for _, item := range files {
		name := item.Name()
		tokens := strings.Split(name, ".")
		AllTemplates[tokens[0]], _ = DoGet("/markup/" + name)
	}

}
