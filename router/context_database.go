package router

import (
	"fmt"
	"os"
)

func TablenameFromName(name string) string {
	prefix := os.Getenv("FEEDBACK_NAME")
	return prefix + "_" + name
}

func (c *Context) SelectAllFrom(name string) []map[string]any {
	tableName := TablenameFromName(name)
	sql := fmt.Sprintf("SELECT * FROM %s ORDER BY created_at desc limit 30", tableName)
	ms := []map[string]any{}
	rows, err := c.Db.Queryx(sql)
	if err != nil {
		return ms
	}
	defer rows.Close()
	for rows.Next() {
		m := make(map[string]any)
		rows.MapScan(m)
		ms = append(ms, m)
	}
	return ms
}

func (c *Context) SelectOneFrom(id, name string) map[string]any {
	tableName := TablenameFromName(name)
	sql := fmt.Sprintf("SELECT * FROM %s where guid = $1", tableName)
	m := map[string]any{}
	rows, err := c.Db.Queryx(sql, id)
	if err != nil {
		return m
	}
	defer rows.Close()
	rows.Next()
	rows.MapScan(m)
	return m
}
