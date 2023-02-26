package controller

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/andrewarrow/feedback/files"
)

type Render struct {
	Vars     *Vars
	Template *template.Template
	Site     Site
}

func NewRender(t *template.Template) *Render {
	r := Render{}
	r.Template = t
	jsonString := files.ReadFile("data/site.json")
	json.Unmarshal([]byte(jsonString), &r.Site)
	r.Vars = NewVars(&r.Site)

	return &r
}

func (r *Render) Execute(writer http.ResponseWriter, file string, vars any) {
	r.Template.ExecuteTemplate(writer, file, vars)
}
