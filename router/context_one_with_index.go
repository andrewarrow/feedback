package router

import "fmt"

func (c *Context) OneWithIndex(index int64, modelName string, where string, params ...any) map[string]any {
	model := c.FindModel(modelName)
	sql := fmt.Sprintf("SELECT * FROM %s %s", model.TableName(), where)
	m := map[string]any{}
	rows, err := c.Dbs[index].Queryx(sql, params...)
	if err != nil {
		return m
	}
	defer rows.Close()
	rows.Next()
	rows.MapScan(m)
	CastFields(model, m)
	return m
}
