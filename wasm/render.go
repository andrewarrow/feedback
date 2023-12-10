package wasm

import (
	"bytes"
	"embed"
	"strings"
	"text/template"

	"github.com/andrewarrow/feedback/common"
)

var EmbeddedTemplates embed.FS
var AllTemplates map[string]any
var UseLive = true

func (d *Document) Render(id, name string, vars map[string]any) {
	div := d.ById(id)
	templateText := ""
	if UseLive {
		templateText = AllTemplates[name].(string)
	} else {
		templateBytes, _ := EmbeddedTemplates.ReadFile("views/" + name)
		templateText = string(templateBytes)
	}
	t := template.New("")
	t = t.Funcs(common.TemplateFunctions())
	t, _ = t.Parse(string(templateText))
	content := new(bytes.Buffer)
	t.Execute(content, vars)
	//t.ExecuteTemplate(content, name, vars)
	cb := content.Bytes()
	div.Set("innerHTML", string(cb))
}

func LoadAllTemplates(list string, doGet func(string) string) {
	tokens := strings.Split(list, ",")

	AllTemplates = map[string]any{}
	for _, item := range tokens {
		AllTemplates[item] = doGet("/markup/" + item)
	}

}

/*
func runTemplate(name string, vars map[string]any) string {
	//fmt.Println(templateText)

}*/
