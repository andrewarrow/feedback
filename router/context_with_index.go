package router

import "fmt"

func (c *Context) WithIndex(index int64, sql string, params ...any) []map[string]any {
	ms := []map[string]any{}
	rows, err := c.Dbs[index].Queryx(sql, params...)
	if err != nil {
		fmt.Println(sql, err)
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
