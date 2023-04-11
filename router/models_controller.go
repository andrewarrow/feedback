package router

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/andrewarrow/feedback/models"
	"github.com/andrewarrow/feedback/sqlgen"
	"github.com/jmoiron/sqlx"
)

type ModelsVars struct {
	Models     []*models.Model
	SchemaJson string
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
	c.Layout = "models_layout.html"
	if len(c.User) == 0 {
		c.UserRequired = true
		return
	}
	if IsAdmin(c.User) == false {
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
		} else if c.Method == "PATCH" {
			handleModelPatch(c, second)
		} else if c.Method == "POST" && second == "json" {
			http.Redirect(c.Writer, c.Request, "/models", 302)
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
		tableName := model.TableName()
		c.Db.Exec(fmt.Sprintf("delete from %s where guid=$1", tableName), third)
		http.Redirect(c.Writer, c.Request, "/models/"+second, 302)
		return
	} else if c.Method == "GET" {
		vars := FieldVars{}
		vars.Rows = model.Fields
		vars.Model = model
		vars.Row = c.SelectOne(second, "where guid=$1", []any{third})
		c.SendContentInLayout("models_edit.html", &vars, 200)
		return
	} else if c.Method == "POST" {

		list := []string{}
		for _, field := range model.Fields {
			list = append(list, field.Name)
		}
		c.ReadFormValuesIntoParams(list...)
		c.ValidateUpdate(second)
		c.Update(second, "where guid=", third)
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
	c.ReadJsonBodyIntoParams()
	newModel := models.Model{}
	name := c.Params["name"]
	if name != nil {
		newModel.Name = models.RemoveNonAlphanumeric(strings.ToLower(name.(string)))
	}

	if len(strings.TrimSpace(newModel.Name)) < 3 {
		c.Writer.WriteHeader(422)
		fmt.Fprintf(c.Writer, "length of name must be > 2")
	} else {
		f := models.Field{}
		f.Name = "guid"
		f.Flavor = "uuid"
		f.Index = "yes"
		newModel.Fields = append(newModel.Fields, &f)
		c.router.Site.Models = append(c.router.Site.Models, &newModel)
		c.saveSchema()
		MakeTable(c.Db, &newModel)
		vars := ModelsVars{}
		vars.Models = c.router.Site.Models
		c.ExecuteTemplate("models_list.html", vars)
	}
}

func ModelsShow(c *Context, rawId string) {
	id := models.RemoveNonAlphanumeric(rawId)
	model := c.router.Site.FindModel(id)
	if model == nil {
		c.NotFound = true
		return
	}

	rows := c.SelectAll(id, "order by id", []any{}, "")
	vars := ModelVars{}
	vars.Rows = rows
	vars.Model = model
	c.SendContentInLayout("models_show.html", vars, 200)
}

func ModelsCreateWithId(c *Context, id string) {
	model := c.router.Site.FindModel(id)
	if model == nil {
		c.NotFound = true
		return
	}
	tableName := model.TableName()
	fieldName := c.Request.FormValue("name")
	index := c.Request.FormValue("index")
	flavor := c.Request.FormValue("flavor")
	if fieldName == "" {
		sql, params := sqlgen.InsertRowWithRandomDefaults(tableName, model.Fields, map[string]any{})
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
	tableName := model.TableName()
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
		} else if f.Flavor == "timestamp" && m[f.Name] != nil {
			t := m[f.Name].(time.Time)
			m[f.Name] = t.Format(models.ISO8601)
		} else {
			m[f.Name] = fmt.Sprintf("%s", m[f.Name])
		}
	}
	return m
}

func handleModelPatch(c *Context, modelName string) {
	model := c.FindModel(modelName)
	model.Name = c.Request.FormValue("name")
	c.saveSchema()
	http.Redirect(c.Writer, c.Request, "/models", 302)
}
