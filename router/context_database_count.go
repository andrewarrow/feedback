package router

import (
	"fmt"
)

func (c *Context) Count(modelName string, whereString string, params []any) int64 {
	model := c.FindModel(modelName)
	sql := fmt.Sprintf("SELECT count(1) as c FROM %s %s", model.TableName(), whereString)
	m := map[string]any{}
	rows, err := c.Db.Queryx(sql, params...)
	if err != nil {
		return 0
	}
	defer rows.Close()
	rows.Next()
	rows.MapScan(m)
	return m["c"].(int64)
}

func (c *Context) SendIntAsJson(wrapper string, val int64) {
	m := map[string]any{wrapper: val}
	c.SendContentAsJson(m, 200)
}
