package router

import (
	"encoding/json"
	"fmt"
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
	Rows  []map[string]any
}

func ModelsShow(c *Context, id string) {
	model := c.router.Site.FindModel(id)
	if model == nil {
		c.notFound = true
		return
	}

	tableName := util.Plural(model.Name)
	//c.db.Exec(sqlgen.MysqlCreateTable(tableName))
	c.db.Exec(sqlgen.PgCreateTable(tableName))
	sql := `ALTER TABLE %s ADD COLUMN %s varchar(255) default '';`
	for _, field := range model.Fields {
		c.db.Exec(fmt.Sprintf(sql, tableName, field.Name))
		if field.Index == "yes" {
			sql := `create index %s_index on %s(%s);`
			c.db.Exec(fmt.Sprintf(sql, field.Name, tableName, field.Name))
		}
	}

	vars := ModelVars{}
	vars.Rows = []map[string]any{}
	rows, _ := c.db.Queryx(fmt.Sprintf("SELECT * FROM %s ORDER BY id limit 30", tableName))
	for rows.Next() {
		m := make(map[string]any)
		rows.MapScan(m)
		for k, v := range m {
			m[k] = fmt.Sprintf("%s", v)
		}
		vars.Rows = append(vars.Rows, m)
	}

	vars.Model = model
	c.SendContentInLayout("models_show.html", vars, 200)
}

func ModelsCreateWithId(c *Context, id string) {
	model := c.router.Site.FindModel(id)
	if model == nil {
		c.notFound = true
		return
	}
	tableName := util.Plural(model.Name)
	fieldName := c.request.FormValue("name")
	if fieldName == "" {
		c.db.Exec(sqlgen.InsertRow(tableName, model.Fields))
	} else {
		f := models.Field{}
		f.Name = fieldName
		f.Flavor = "bar"
		//c.router.Site.AddField(id, f)
		model.Fields = append(model.Fields, f)
	}
	http.Redirect(c.writer, c.request, c.path, 302)
}
