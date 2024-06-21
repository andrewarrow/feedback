package router

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/andrewarrow/feedback/models"
	"github.com/andrewarrow/feedback/sqlgen"
	"github.com/andrewarrow/feedback/util"
	"github.com/lib/pq"
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

func (c *Context) ValidateAndInsert(modelString string) string {
	msg := c.ValidateCreate(modelString)
	if msg != "" {
		return msg
	}
	return c.Insert(modelString)
}

func (c *Context) Insert(modelString string) string {
	model := c.FindModel(modelString)
	tableName := model.TableName()
	funcToRun := c.Router.beforeFuncToRun(modelString)

	if funcToRun != nil {
		funcToRun(c)
	}
	sql, params := sqlgen.InsertRowNoRandomDefaults(DB_FLAVOR, tableName, model.Fields, c.Params)
	//fmt.Println(sql, params)
	//customDB := &persist.CustomDB{DB: c.Db}
	//_, err := customDB.ExecWithLogging(sql, params...)
	_, err := c.Db.Exec(sql, params...)
	if err != nil {
		if sqlErr, ok := err.(*pq.Error); ok {
			return sqlErr.Error() + " " + sqlErr.Detail + " " + sqlErr.Message + " " +
				sqlErr.Hint + " " +
				sqlErr.Position + " " + sqlErr.Column + " " + sqlErr.DataTypeName
		}
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
	sql, params := sqlgen.UpdateRowFromParams(DB_FLAVOR, tableName, list, c.Params, where)
	if os.Getenv("DEBUG") == "1" {
		fmt.Println("sqlgen.UpdateRowFromParams", sql, params)
	}
	params = append(params, lastParam)
	_, err := c.Db.Exec(sql, params...)
	if err != nil {
		if sqlErr, ok := err.(*pq.Error); ok {
			return sqlErr.Error() + " " + sqlErr.Detail + " " + sqlErr.Message + " " +
				sqlErr.Hint + " " +
				sqlErr.Position + " " + sqlErr.Column + " " + sqlErr.DataTypeName
		}
		return err.Error()
	}
	return ""
}
