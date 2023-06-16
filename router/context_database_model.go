package router

import (
	"fmt"
	"strconv"
)

func (r *Router) All(modelName string, where, offset string, params ...any) []map[string]any {
	return r.SelectAll(modelName, where, params, offset)
}

func (c *Context) All(modelName string, where, offset string, params ...any) []map[string]any {
	return c.SelectAll(modelName, where, params, offset)
}

func (c *Context) SelectAll(modelName string, where string, params []any, offset string) []map[string]any {
	return c.Router.SelectAll(modelName, where, params, offset)
}

func (r *Router) SelectAll(modelName string, where string, params []any, offset string) []map[string]any {
	model := r.FindModel(modelName)
	offsetString := ""
	if offset != "" {
		offsetInt, _ := strconv.Atoi(offset)
		offsetString = fmt.Sprintf("OFFSET %d", offsetInt)
	}
	sql := fmt.Sprintf("SELECT * FROM %s %s limit 30 %s", model.TableName(), where, offsetString)
	ms := []map[string]any{}
	rows, err := r.Db.Queryx(sql, params...)
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

func (r *Router) One(modelName string, where string, params ...any) map[string]any {
	return r.SelectOne(modelName, where, params)
}

func (c *Context) One(modelName string, where string, params ...any) map[string]any {
	return c.Router.SelectOne(modelName, where, params)
}

func (c *Context) SelectOne(modelName string, where string, params []any) map[string]any {
	return c.Router.SelectOne(modelName, where, params)
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
