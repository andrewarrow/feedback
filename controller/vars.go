package controller

import (
	"html/template"
	"net/http"
)

type Vars struct {
	Title string
	Phone string
}

func NewVars() *Vars {
	v := Vars{}
	v.Title = "Feedback"
	return &v
}

func (v *Vars) Fill(r *Render) {
	v.Title = r.Vars.Title
	v.Phone = r.Vars.Phone
}

type Render struct {
	Vars     *Vars
	Template *template.Template
}

func NewRender(v *Vars, t *template.Template) *Render {
	r := Render{}
	r.Vars = v
	r.Template = t
	return &r
}

func (r *Render) Execute(writer http.ResponseWriter, file string, vars any) {
	r.Template.ExecuteTemplate(writer, file, vars)
}
