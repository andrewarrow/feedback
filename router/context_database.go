package router

import (
	"fmt"
	"time"

	"github.com/andrewarrow/feedback/models"
	"github.com/andrewarrow/feedback/prefix"
	"github.com/xeonx/timeago"
)

func (c *Context) EmptyParams() []any {
	return []any{}
}
func (c *Context) FindModel(name string) *models.Model {
	return c.router.Site.FindModel(name)
}

func (c *Context) Count(name string, where string) int64 {
	tableName := prefix.Tablename(name)
	whereString := ""
	if where != "" {
		whereString = " where " + where
	}
	sql := fmt.Sprintf("SELECT count(1) as c FROM %s%s", tableName, whereString)
	m := map[string]any{}
	rows, err := c.Db.Queryx(sql)
	if err != nil {
		return 0
	}
	defer rows.Close()
	rows.Next()
	rows.MapScan(m)
	return m["c"].(int64)
}

func CastFields(model *models.Model, m map[string]any) {
	if len(m) == 0 {
		return
	}
	for _, field := range model.Fields {
		if field.Flavor == "timestamp" {
			tm := m[field.Name].(time.Time)
			ago := timeago.English.Format(tm)
			m[field.Name] = tm.Format(models.HUMAN)
			m[field.Name+"_ago"] = ago
		} else if field.Flavor == "int" {
			m[field.Name] = m[field.Name].(int64)
		} else {
			m[field.Name] = fmt.Sprintf("%s", m[field.Name])
		}
	}
}

func (r *Router) SelectOneFrom(model *models.Model, where string, params []any) map[string]any {
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

func (c *Context) SelectOneFrom(model *models.Model, where string, params []any) map[string]any {
	return c.router.SelectOneFrom(model, where, params)
}
