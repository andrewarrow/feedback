package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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
	if c.User.IsAdmin() == false {
		c.NotFound = true
		return
	}
	if second == "" {
		handleModelsIndex(c)
	} else if third != "" {
		handleThird(c, second, third)
	} else {
		if c.Method == "GET" {
			ModelsShow(c, second)
		} else {
			ModelsCreateWithId(c, second)
		}
	}
}

func handleThird(c *Context, second, third string) {
	model := c.router.Site.FindModel(second)
	if model == nil {
		c.NotFound = true
		return
	}

	if c.Method == "DELETE" {
		safeName := models.RemoveNonAlphanumeric(second)
		c.Db.Exec(fmt.Sprintf("delete from %s where guid=$1", util.Plural(safeName)), third)
		http.Redirect(c.Writer, c.Request, "/models/"+second, 302)
		return
	} else if c.Method == "GET" {
		vars := FieldVars{}
		vars.Rows = model.Fields
		vars.Model = model
		vars.Row = FetchOneRow(c.Db, model, third)
		c.SendContentInLayout("models_edit.html", &vars, 200)
		return
	} else if c.Method == "POST" {

		params := []any{}
		for _, field := range model.Fields {
			var value any
			if field.Flavor == "int" {
				stringValue := c.Request.FormValue(field.Name)
				value, _ = strconv.Atoi(stringValue)
			} else {
				value = c.Request.FormValue(field.Name)
			}
			params = append(params, value)
		}
		params = append(params, third)
		sql := sqlgen.UpdateRow(model)
		c.Db.Exec(sql, params...)
		http.Redirect(c.Writer, c.Request, "/models/"+second, 302)
		return
	}
	c.NotFound = true
}

func handleModelsIndex(c *Context) {
	if c.Request.Method == "GET" {
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
		c.Writer.WriteHeader(422)
		fmt.Fprintf(c.Writer, "length of name must be > 2")
	} else {
		c.router.Site.Models = append(c.router.Site.Models, &newModel)
		c.saveSchema()
		MakeTable(c.Db, &newModel)
		vars := ModelsVars{}
		vars.Models = c.router.Site.Models
		c.router.Template.ExecuteTemplate(c.Writer, "models_list.html", vars)
	}
}

func ModelsShow(c *Context, rawId string) {
	id := models.RemoveNonAlphanumeric(rawId)
	model := c.router.Site.FindModel(id)
	if model == nil {
		c.NotFound = true
		return
	}

	tableName := util.Plural(model.Name)
	vars := ModelVars{}
	vars.Rows = []map[string]any{}
	rows, err := c.Db.Queryx(fmt.Sprintf("SELECT * FROM %s ORDER BY id limit 30", tableName))
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
		c.NotFound = true
		return
	}
	tableName := util.Plural(model.Name)
	fieldName := c.Request.FormValue("name")
	index := c.Request.FormValue("index")
	flavor := c.Request.FormValue("flavor")
	if fieldName == "" {
		sql, params := sqlgen.InsertRow(tableName, model.Fields)
		c.Db.Exec(sql, params...)
	} else {
		f := models.Field{}
		f.Name = fieldName
		f.Flavor = flavor
		f.Index = index
		model.Fields = append(model.Fields, &f)
		c.saveSchema()
		MakeTable(c.Db, model)
	}
	http.Redirect(c.Writer, c.Request, c.path, 302)
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
