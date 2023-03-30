package router

import (
	"fmt"
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

func (c *Context) UpdateOne(modelName, setString, whereString string, params []any) error {
	model := c.FindModel(modelName)
	sql := fmt.Sprintf("update %s set %s where %s", model.TableName(), setString, whereString)

	_, err := c.Db.Exec(sql, params...)
	return err
}
