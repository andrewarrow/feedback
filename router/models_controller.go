package router

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/andrewarrow/feedback/models"
	"github.com/andrewarrow/feedback/util"
	"github.com/brianvoe/gofakeit/v6"
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

func (mc *ModelsController) CreateWithId(c *Context, id, body string) {
	model := c.router.Site.FindModel(id)

	tableName := util.Plural(model.Name)
	sql := `INSERT INTO %s (username, guid) values ('%s', '%s');`
	username := gofakeit.Username()
	guid := util.PseudoUuid()
	c.db.Exec(fmt.Sprintf(sql, tableName, username, guid))
	http.Redirect(c.writer, c.request, c.path, 302)
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
	Rows  []template.HTML
}

func (mc *ModelsController) Show(c *Context, id string) {
	model := c.router.Site.FindModel(id)

	tableName := util.Plural(model.Name)
	sql := `CREATE TABLE %s (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    username varchar(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE KEY unique_username (username)
) ENGINE InnoDB;`
	c.db.Exec(fmt.Sprintf(sql, tableName))
	sql = `ALTER TABLE %s ADD COLUMN %s varchar(255) default '';`
	for _, field := range model.Fields {
		c.db.Exec(fmt.Sprintf(sql, tableName, field.Name))
		if field.Index == "yes" {
			sql := `create index %s_index on %s(%s);`
			c.db.Exec(fmt.Sprintf(sql, field.Name, tableName, field.Name))
		}
	}

	vars := ModelVars{}
	tds := []string{}
	for _, field := range model.Fields {
		tds = append(tds, fmt.Sprintf("<th>%s</th>", field.Name))
	}
	vars.Rows = append(vars.Rows, template.HTML(strings.Join(tds, "")))
	rows, _ := c.db.Queryx("SELECT * FROM users ORDER BY id limit 30")
	for rows.Next() {
		m := make(map[string]any)
		rows.MapScan(m)
		tds := []string{}
		for _, field := range model.Fields {
			data := m[field.Name]
			tds = append(tds, fmt.Sprintf("<td>%s</td>", data))
		}
		vars.Rows = append(vars.Rows, template.HTML(strings.Join(tds, "")))
	}

	vars.Model = model
	c.SendContentInLayout("models_show.html", vars, 200)
}

func (mc *ModelsController) New(c *Context) {
}
func (mc *ModelsController) Destroy(c *Context) {
}
