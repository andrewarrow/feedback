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

func (mc *ModelsController) Index(c *Context) {
	vars := ModelsVars{}
	vars.Models = c.router.Site.Models
	c.SendContentInLayout("models_index.html", vars, 200)
}

func (mc *ModelsController) Create(context *Context) {
}

func (mc *ModelsController) Show(c *Context, id string) {
	c.SendContentInLayout("models_show.html", nil, 200)
}
