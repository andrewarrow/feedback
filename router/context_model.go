package router

import "github.com/andrewarrow/feedback/models"

func (c *Context) Model(modelName string) *models.Model {
	return c.Router.Site.FindModel(modelName)
}

func (c *Context) Models() []*models.Model {
	return c.Router.Site.Models
}
