package router

import "github.com/andrewarrow/feedback/models"

type ModelsVars struct {
	Models []models.Model
}

func (r *Router) ModelsResource(c *Context) {
	vars := ModelsVars{}
	vars.Models = r.Site.Models
	r.SendContentInLayout(c.writer, "models_index.html", vars, 200)
}
