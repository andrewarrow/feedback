package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/andrewarrow/feedback/models"
	"github.com/andrewarrow/feedback/sqlgen"
	"github.com/andrewarrow/feedback/util"
	"github.com/jmoiron/sqlx"
)

type ModelsVars struct {
	Models []*models.Model
}
type ModelVars struct {
	Model *models.Model
	Rows  []map[string]any
}
type FieldVars struct {
	Model *models.Model
	Rows  []*models.Field
	Row   map[string]any
}

func handleModels(c *Context, second, third string) {
	if c.user.IsAdmin() == false {
		c.notFound = true
		return
	}
	if second == "" {
		handleModelsIndex(c)
	} else if third != "" {
		handleThird(c, second, third)
	} else {
		if c.method == "GET" {
			ModelsShow(c, second)
		} else {
			ModelsCreateWithId(c, second)
		}
	}
}

func handleThird(c *Context, second, third string) {
	model := c.router.Site.FindModel(second)
	if model == nil {
		c.notFound = true
		return
	}

	if c.method == "DELETE" {
		safeName := models.RemoveNonAlphanumeric(second)
		c.db.Exec(fmt.Sprintf("delete from %s where guid=$1", util.Plural(safeName)), third)
		http.Redirect(c.writer, c.request, "/models/"+second, 302)
		return
	} else if c.method == "GET" {
		vars := FieldVars{}
		vars.Rows = model.Fields
		vars.Model = model
		vars.Row = FetchOneRow(c.db, model, third)
		c.SendContentInLayout("models_edit.html", &vars, 200)
		return
	} else if c.method == "POST" {

		params := []any{}
		for _, field := range model.Fields {
			value := c.request.FormValue(field.Name)
			params = append(params, value)
		}
		params = append(params, third)
		sql := sqlgen.UpdateRow(model)
		c.db.Exec(sql, params...)
		http.Redirect(c.writer, c.request, "/models/"+second, 302)
		return
	}
	c.notFound = true
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
		c.saveSchema()
		MakeTable(c.db, &newModel)
		vars := ModelsVars{}
		vars.Models = c.router.Site.Models
		c.router.Template.ExecuteTemplate(c.writer, "models_list.html", vars)
	}
}

func ModelsShow(c *Context, rawId string) {
	id := models.RemoveNonAlphanumeric(rawId)
	model := c.router.Site.FindModel(id)
	if model == nil {
		c.notFound = true
		return
	}

	tableName := util.Plural(model.Name)
	vars := ModelVars{}
	vars.Rows = []map[string]any{}
	rows, err := c.db.Queryx(fmt.Sprintf("SELECT * FROM %s ORDER BY id limit 30", tableName))
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		m := make(map[string]any)
		rows.MapScan(m)
		for _, f := range model.Fields {
			if f.Flavor == "int" {
				m[f.Name] = m[f.Name].(int64)
			} else {
				m[f.Name] = fmt.Sprintf("%s", m[f.Name])
			}
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
	index := c.request.FormValue("index")
	flavor := c.request.FormValue("flavor")
	if fieldName == "" {
		sql, params := sqlgen.InsertRow(tableName, model.Fields)
		c.db.Exec(sql, params...)
	} else {
		f := models.Field{}
		f.Name = fieldName
		f.Flavor = flavor
		f.Index = index
		model.Fields = append(model.Fields, &f)
		c.saveSchema()
		MakeTable(c.db, model)
	}
	http.Redirect(c.writer, c.request, c.path, 302)
}

func FetchOneRow(db *sqlx.DB, model *models.Model, guid string) map[string]any {
	tableName := util.Plural(model.Name)
	rows, err := db.Queryx(fmt.Sprintf("SELECT * FROM %s where guid=$1", tableName), guid)
	if err != nil {
		return map[string]any{}
	}
	defer rows.Close()

	rows.Next()
	m := make(map[string]any)
	rows.MapScan(m)
	for _, f := range model.Fields {
		if f.Flavor == "int" {
			m[f.Name] = m[f.Name].(int64)
		} else {
			m[f.Name] = fmt.Sprintf("%s", m[f.Name])
		}
	}
	return m
}
