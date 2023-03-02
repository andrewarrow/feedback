package router

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/andrewarrow/feedback/models"
)

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

func (mc *ModelsController) Create(context *Context, body string) {
}

func (mc *ModelsController) CreateWithJson(c *Context, body string) {
	var params map[string]any
	json.Unmarshal([]byte(body), &params)
	newModel := models.Model{}
	name := params["name"]
	if name != nil {
		newModel.Name = models.RemoveNonAlphanumeric(strings.ToLower(name.(string)))
	}

	if len(strings.TrimSpace(newModel.Name)) < 3 {
		c.writer.WriteHeader(422)
		fmt.Fprintf(c.writer, "length of name must be > 2")
	} else {
		c.router.Site.Models = append(c.router.Site.Models, newModel)
		vars := ModelsVars{}
		vars.Models = c.router.Site.Models
		c.router.Template.ExecuteTemplate(c.writer, "models_list.html", vars)
	}
}

type ModelVars struct {
	Model *models.Model
}

func (mc *ModelsController) Show(c *Context, id string) {
	model := c.router.Site.FindModel(id)
	vars := ModelVars{}
	vars.Model = model
	c.SendContentInLayout("models_show.html", vars, 200)
}

func (mc *ModelsController) New(c *Context) {
}
func (mc *ModelsController) Destroy(c *Context) {
}
