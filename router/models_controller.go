package router

import "github.com/andrewarrow/feedback/models"

type ModelsController struct {
}

func NewModelsController() Controller {
	mc := ModelsController{}
	return &mc
}

type ModelsVars struct {
	Models []models.Model
}

func (mc *ModelsController) Index(r *Router, context *Context) {
	vars := ModelsVars{}
	vars.Models = r.Site.Models
	r.SendContentInLayout(context.writer, "models_index.html", vars, 200)
}

func (mc *ModelsController) Create(r *Router, context *Context) {
}
