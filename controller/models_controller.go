package controller

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/andrewarrow/feedback/models"
)

type ModelsController struct {
	render *Render
	writer http.ResponseWriter
	models []models.Model
}

type ModelVars struct {
	Vars
	Models []models.Model
}

func NewModelsController(models []models.Model, render *Render) *ModelsController {
	m := ModelsController{}
	m.models = models
	for i, _ := range m.models {
		m.models[i].Index = i + 1
	}
	m.render = render
	return &m
}

func (m *ModelsController) Index() {
	vars := ModelVars{}
	vars.Models = m.models
	vars.Fill(m.render)
	m.render.Execute(m.writer, "models_index.html", vars)
}

func (m *ModelsController) Create() {
	h := m.writer.Header()
	h.Set("Location", "/models")
	m.writer.WriteHeader(301)
}

func (m *ModelsController) HandlePath(writer http.ResponseWriter,
	request *http.Request) {
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
