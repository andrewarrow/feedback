package router

import (
	"encoding/json"
	"strings"

	"github.com/andrewarrow/feedback/models"
	"github.com/andrewarrow/feedback/sqlgen"
	"github.com/andrewarrow/feedback/util"
)

func (c *Context) SendRowAsJson(wrapper string, row map[string]any) {
	util.RemoveSensitiveKeys(row)
	m := map[string]any{wrapper: row}
	c.SendContentAsJson(m, 200)
}

func (c *Context) SendContentAsJson(thing any, status int) {
	list, ok := thing.([]map[string]any)
	if ok {
		for _, item := range list {
			util.RemoveSensitiveKeys(item)
		}
	}

	if c.Batch {
		c.BatchThing = thing
		return
	}

	asBytes, _ := json.Marshal(thing)
	//fmt.Println(string(asBytes))
	ae := c.Request.Header.Get("Accept-Encoding")
	doZip := false
	if strings.Contains(ae, "gzip") {
		doZip = true
	}

	doZippyJson(doZip, asBytes, status, c.Writer)
}

func (c *Context) SendContentAsJsonMessage(message string, status int) {
	m := map[string]any{"info": message}
	c.SendContentAsJson(m, status)
}

func (c *Context) Insert(modelString string) string {
	model := c.FindModel(modelString)
	tableName := model.TableName()
	funcToRun := c.Router.beforeFuncToRun(modelString)

	if funcToRun != nil {
		funcToRun(c)
	}
	sql, params := sqlgen.InsertRowNoRandomDefaults(tableName, model.Fields, c.Params)
	//fmt.Println(sql, params)
	_, err := c.Db.Exec(sql, params...)
	if err != nil {
		return err.Error()
	}
	return ""
}

func (c *Context) Update(modelString, where string, lastParam any) string {
	model := c.FindModel(modelString)
	tableName := model.TableName()

	already := map[string]bool{}
	list := []*models.Field{}
	for _, field := range model.Fields {
		if c.Params[field.Name] == nil {
			continue
		}
		if already[field.Name] {
			continue
		}
		already[field.Name] = true
		list = append(list, field)
	}
	sql, params := sqlgen.UpdateRowFromParams(tableName, list, c.Params, where)
	//fmt.Println(sql, params)
	params = append(params, lastParam)
	_, err := c.Db.Exec(sql, params...)
	if err != nil {
		return err.Error()
	}
	return ""
}
