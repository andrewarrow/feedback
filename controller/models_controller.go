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
	site   *Site
}

type ModelVars struct {
	Vars
	Models []models.Model
}

func NewModelsController(render *Render) *ModelsController {
	m := ModelsController{}
	m.site = &render.Site
	m.render = render
	return &m
}

func (m *ModelsController) Index() {
	vars := ModelVars{}
	vars.Models = m.site.Models
	vars.Fill(m.render)
	m.render.Execute(m.writer, "models_index.html", vars)
}

func (m *ModelsController) Create() {
	h := m.writer.Header()
	h.Set("Location", "/models")
	m.writer.WriteHeader(301)
}

func (m *ModelsController) CreateWithJson(jsonString string) {
	vars := ModelVars{}
	newModel := models.Model{}
	newModel.Name = "foo"
	m.site.Models = append(m.site.Models, newModel)
	vars.Models = m.site.Models
	m.render.Execute(m.writer, "models_list.html", vars)
}

func (m *ModelsController) HandlePath(writer http.ResponseWriter,
	request *http.Request) {
	m.writer = writer
	method := request.Method
	// path := request.URL.Path
	if method == "GET" {
		m.Index()
	} else if method == "POST" {
		//fmt.Printf("%+v\n", request.Header)
		buffer := new(bytes.Buffer)
		buffer.ReadFrom(request.Body)
		fmt.Println("POST", buffer.String())
		contentType := request.Header["Content-Type"]
		if contentType[0] == "application/x-www-form-urlencoded" {
			m.Create()
		} else {
			m.CreateWithJson(buffer.String())
		}
	}
}
