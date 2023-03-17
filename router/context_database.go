package router

import (
	"fmt"

	"github.com/andrewarrow/feedback/models"
	"github.com/andrewarrow/feedback/prefix"
)

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

func (c *Context) SelectAllFrom(name, order, where string) []*map[string]any {
	tableName := prefix.Tablename(name)
	whereString := ""
	if where != "" {
		whereString = "where " + where
	}
	orderString := "created_at desc"
	if order != "" {
		orderString = order
	}
	sql := fmt.Sprintf("SELECT * FROM %s %s ORDER BY %s limit 30", tableName, whereString, orderString)
	ms := []*map[string]any{}
	rows, err := c.Db.Queryx(sql)
	if err != nil {
		return ms
	}
	defer rows.Close()
	for rows.Next() {
		m := make(map[string]any)
		rows.MapScan(m)
		ms = append(ms, &m)
	}
	return ms
}

func (c *Context) SelectOneFrom(id, name string) *map[string]any {
	tableName := prefix.Tablename(name)
	sql := fmt.Sprintf("SELECT * FROM %s where guid = $1", tableName)
	m := map[string]any{}
	rows, err := c.Db.Queryx(sql, id)
	if err != nil {
		return &m
	}
	defer rows.Close()
	rows.Next()
	rows.MapScan(m)
	return &m
}
