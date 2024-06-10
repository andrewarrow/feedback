package router

import "github.com/andrewarrow/feedback/sqlgen"

func (c *Context) Upsert(modelString, where string, lastParam any) string {
	model := c.FindModel(modelString)
	tableName := model.TableName()
	sql, params := sqlgen.InsertRowNoRandomDefaults(DB_FLAVOR, tableName, model.Fields, c.Params)
	_, err := c.Db.Exec(sql, params...)
	if err != nil {
		return c.Update(modelString, where, lastParam)
	}
	return ""
}
