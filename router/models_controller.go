package router

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/andrewarrow/feedback/models"
	"github.com/andrewarrow/feedback/sqlgen"
	"github.com/andrewarrow/feedback/util"
)

type ModelsVars struct {
	Models []*models.Model
}

func handleModelsIndex(c *Context) {
	if c.request.Method == "GET" {
		vars := ModelsVars{}
		vars.Models = c.router.Site.Models
		c.SendContentInLayout("models_index.html", vars, 200)
		return
	}
	handleModelsCreateWithJson(c)
}

func handleModels(c *Context, second, third string) {
	if second == "" {
		handleModelsIndex(c)
	} else if third != "" {
		c.notFound = true
	} else {
		if c.method == "GET" {
			ModelsShow(c, second)
		} else {
			ModelsCreateWithId(c, second)
		}
	}
}

func handleModelsCreateWithJson(c *Context) {
	body := c.BodyAsString()
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
		c.router.Site.Models = append(c.router.Site.Models, &newModel)
		vars := ModelsVars{}
		vars.Models = c.router.Site.Models
		c.router.Template.ExecuteTemplate(c.writer, "models_list.html", vars)
	}
}

type ModelVars struct {
	Model *models.Model
	Rows  []template.HTML
}

func ModelsShow(c *Context, id string) {
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
	rows, _ := c.db.Queryx(fmt.Sprintf("SELECT * FROM %s ORDER BY id limit 30", tableName))
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

func ModelsCreateWithId(c *Context, id string) {
	model := c.router.Site.FindModel(id)
	tableName := util.Plural(model.Name)
	fieldName := c.request.FormValue("name")
	if fieldName == "" {
		c.db.Exec(sqlgen.InsertRow(tableName, model.Fields))
	} else {
		f := models.Field{}
		f.Name = fieldName
		f.Flavor = "bar"
		c.router.Site.AddField(id, f)
	}
	http.Redirect(c.writer, c.request, c.path, 302)
}
