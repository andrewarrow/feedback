package router

import (
	"encoding/json"
	"fmt"

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
	list, ok := thing.([]any)
	if ok {
		for _, item := range list {
			util.RemoveSensitiveKeys(item.(map[string]any))
		}
	}

	asBytes, _ := json.Marshal(thing)

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Header().Set("Cache-Control", "none")
	c.Writer.Header().Set("Content-Length", fmt.Sprintf("%d", len(asBytes)))
	c.Writer.WriteHeader(status)
	c.Writer.Write(asBytes)
}

func (c *Context) SendContentAsJsonMessage(message string, status int) {
	m := map[string]any{"info": message}
	c.SendContentAsJson(m, status)
}

func (c *Context) Insert(modelString string) string {
	model := c.FindModel(modelString)
	tableName := model.TableName()
	funcToRun := c.router.beforeFuncToRun(modelString)

	if funcToRun != nil {
		funcToRun(c)
	}
	sql, params := sqlgen.InsertRowNoRandomDefaults(tableName, model.Fields, c.Params)
	_, err := c.Db.Exec(sql, params...)
	if err != nil {
		return err.Error()
	}
	return ""
}

func (c *Context) Update(modelString, where string, lastParam any) string {
	model := c.FindModel(modelString)
	tableName := model.TableName()

	list := []*models.Field{}
	for _, field := range model.Fields {
		if c.Params[field.Name] == nil {
			continue
		}
		list = append(list, field)
	}
	sql, params := sqlgen.UpdateRowFromParams(tableName, list, c.Params, where)
	params = append(params, lastParam)
	_, err := c.Db.Exec(sql, params...)
	if err != nil {
		return err.Error()
	}
	return ""
}
