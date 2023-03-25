package router

import (
	"net/http"

	"github.com/andrewarrow/feedback/models"
)

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
	if second != "" && third != "" && c.Method == "GET" {
		handleFieldsShow(c, second, third)
		return
	}
	if second != "" && third != "" && c.Method == "PATCH" {
		handleFieldsPatch(c, second, third)
		return
	}
	c.NotFound = true
}

func handleFieldsShow(c *Context, modelName, fieldName string) {
	model := c.FindModel(modelName)
	field := models.FindField(model, fieldName)
	m := map[string]any{"model": model, "field": field}
	c.SendContentInLayout("fields_show.html", m, 200)
}

func handleFieldsPatch(c *Context, modelName, fieldName string) {
	http.Redirect(c.Writer, c.Request, "/models/"+modelName, 302)
}
