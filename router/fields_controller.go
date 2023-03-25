package router

import "github.com/andrewarrow/feedback/models"

func handleFields(c *Context, second, third string) {
	c.Layout = "models_layout.html"
	if c.User == nil {
		c.UserRequired = true
		return
	}
	if IsAdmin(c.User) == false {
		c.NotFound = true
		return
	}
	if second != "" && third != "" {
		handleFieldsShow(c, second, third)
		return
	}
	c.NotFound = true
}

func handleFieldsShow(c *Context, modelName, fieldName string) {
	model := c.FindModel(modelName)
	field := models.FindField(model, fieldName)
	c.SendContentInLayout("fields_show.html", field, 200)
}
