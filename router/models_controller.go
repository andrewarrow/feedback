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

func (mc *ModelsController) Index(context *Context) {
	vars := ModelsVars{}
	vars.Models = context.router.Site.Models
	context.SendContentInLayout("models_index.html", vars, 200)
}

func (mc *ModelsController) Create(context *Context) {
}
