package controller

import "net/http"

type ModelsController struct {
}

func NewModelsController() *ModelsController {
	m := ModelsController{}
	return &m
}

func (m *ModelsController) Index() {
}
func (m *ModelsController) Create() {
}
func (r *ModelsController) HandlePath(writer http.ResponseWriter, path string, vars Vars) {
}
