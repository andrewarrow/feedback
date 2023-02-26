package controller

import (
	"html/template"
	"net/http"

	"github.com/andrewarrow/feedback/models"
)

type ModelsController struct {
	vars   Vars
	writer http.ResponseWriter
	models []models.Model
}

type ModelVars struct {
	Vars
	Models []models.Model
}

func NewModelsController(models []models.Model) *ModelsController {
	m := ModelsController{}
	m.models = models
	return &m
}

func (m *ModelsController) Index() {
	vars := ModelVars{}
	vars.Header = m.vars.Header
	vars.Footer = m.vars.Footer
	vars.Models = m.models
	t, _ := template.ParseFiles("views/models_index.html")
	t.Execute(m.writer, vars)
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
