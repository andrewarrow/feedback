package controller

import (
	"html/template"
	"net/http"
)

type ModelsController struct {
	vars   Vars
	writer http.ResponseWriter
}

func NewModelsController() *ModelsController {
	m := ModelsController{}
	return &m
}

func (m *ModelsController) Index() {
	t, _ := template.ParseFiles("views/models_index.html")
	t.Execute(m.writer, m.vars)
}

func (m *ModelsController) Create() {
}

func (m *ModelsController) HandlePath(writer http.ResponseWriter, path string, vars Vars) {
	m.vars = vars
	m.writer = writer
	m.Index()
}
