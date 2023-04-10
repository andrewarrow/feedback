package router

import "fmt"

func (c *Context) Delete(modelName, fieldName string, id int64) {
	model := c.FindModel(modelName)
	sql := fmt.Sprintf("delete from %s where %s=$1", model.TableName(), fieldName)
	c.Db.Exec(sql, id)
}
