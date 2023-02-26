package controller

import (
	"bytes"
	"fmt"
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
	t := template.New("models_index.html")
	t = t.Funcs(TemplateFunctions())
	t, _ = t.ParseFiles("views/models_index.html")
	t.Execute(m.writer, vars)
}

func (m *ModelsController) Create() {
	h := m.writer.Header()
	h.Set("Location", "/models")
	m.writer.WriteHeader(301)
}

func (m *ModelsController) HandlePath(writer http.ResponseWriter,
	request *http.Request, vars Vars) {
	m.vars = vars
	m.writer = writer
	method := request.Method
	// path := request.URL.Path
	if method == "GET" {
		m.Index()
	} else if method == "POST" {
		buffer := new(bytes.Buffer)
		buffer.ReadFrom(request.Body)
		fmt.Println("POST", buffer.String())
		m.Create()
	}
}
