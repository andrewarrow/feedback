package router

import (
	"fmt"
	"os"
)

func (c *Context) SelectAllFrom(name string) []map[string]any {
	prefix := os.Getenv("FEEDBACK_NAME")
	tableName := prefix + "_" + name
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
