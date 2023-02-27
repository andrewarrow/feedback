package controller

import (
	"html/template"
	"net/http"
)

type Render struct {
	Vars     *Vars
	Template *template.Template
	Site     *Site
}

func NewRender(t *template.Template, vars *Vars, site *Site) *Render {
	r := Render{}
	r.Template = t
	r.Vars = vars
	r.Site = site

	return &r
}

func (r *Render) Execute(writer http.ResponseWriter, file string, vars any) {
	r.Template.ExecuteTemplate(writer, file, vars)
}
