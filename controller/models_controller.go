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
	h := m.writer.Header()
	h.Set("Location", "/models")
	m.writer.WriteHeader(301)
}

func (m *ModelsController) HandlePath(writer http.ResponseWriter,
	path, method string, vars Vars) {
	m.vars = vars
	m.writer = writer
	if method == "GET" {
		m.Index()
	} else if method == "POST" {
		m.Create()
	}
}
