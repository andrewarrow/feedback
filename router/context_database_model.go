package router

import (
	"fmt"
	"strings"

	"github.com/andrewarrow/feedback/util"
)

func (c *Context) SelectAll(modelName string, where string, params []any) []map[string]any {
	model := c.FindModel(modelName)
	sql := fmt.Sprintf("SELECT * FROM %s %s limit 30", model.TableName(), where)
	ms := []map[string]any{}
	rows, err := c.Db.Queryx(sql, params...)
	if err != nil {
		return ms
	}
	defer rows.Close()
	for rows.Next() {
		m := make(map[string]any)
		rows.MapScan(m)
		CastFields(model, m)
		ms = append(ms, m)
	}
	return ms
}

func (c *Context) SelectOne(modelName string, where string, params []any) map[string]any {
	return c.router.SelectOne(modelName, where, params)
}

func (r *Router) SelectOne(modelName string, where string, params []any) map[string]any {
	model := r.FindModel(modelName)
	sql := fmt.Sprintf("SELECT * FROM %s %s", model.TableName(), where)
	m := map[string]any{}
	rows, err := r.Db.Queryx(sql, params...)
	if err != nil {
		return m
	}
	defer rows.Close()
	rows.Next()
	rows.MapScan(m)
	CastFields(model, m)
	return m
}

func (c *Context) Insert(modelName string, params map[string]any) string {
	model := c.FindModel(modelName)

	fieldPositions := []string{}
	valueList := []any{}
	valuePositions := []string{}
	guid := util.PseudoUuid()
	count := 1
	for _, field := range model.Fields {
		if field.Name == "id" {
			continue
		} else if field.Name == "created_at" {
			continue
		}

		fieldPositions = append(fieldPositions, field.Name)
		valuePositions = append(valuePositions, fmt.Sprintf("$%d", count))
		count++

		if field.Name == "guid" {
			valueList = append(valueList, guid)
		} else {
			val := params[field.Name]
			if val != nil {
				valueList = append(valueList, val)
			} else {
				valueList = append(valueList, field.Default())
			}
		}
	}

	fields := strings.Join(fieldPositions, ",")
	values := strings.Join(valuePositions, ",")
	sql := fmt.Sprintf("insert into %s (%s) values (%s)", model.TableName(), fields, values)

	c.Db.Exec(sql, valueList...)
	return guid
}

func (c *Context) UpdateOne(modelName, setString, whereString string, params []any) {
	model := c.FindModel(modelName)
	sql := fmt.Sprintf("update %s set %s where %s", model.TableName(), setString, whereString)

	c.Db.Exec(sql, params...)
}
