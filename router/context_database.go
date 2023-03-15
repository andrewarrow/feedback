package router

import (
	"fmt"
	"os"

	"github.com/andrewarrow/feedback/models"
)

func (c *Context) FindModel(name string) *models.Model {
	return c.router.Site.FindModel(name)
}

func (c *Context) Count(name string) int64 {
	tableName := TablenameFromName(name)
	sql := fmt.Sprintf("SELECT count(1) as c FROM %s", tableName)
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

func TablenameFromName(name string) string {
	prefix := os.Getenv("FEEDBACK_NAME")
	return prefix + "_" + name
}

func (c *Context) SelectAllFrom(name string) []*map[string]any {
	tableName := TablenameFromName(name)
	sql := fmt.Sprintf("SELECT * FROM %s ORDER BY created_at desc limit 30", tableName)
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
	tableName := TablenameFromName(name)
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
