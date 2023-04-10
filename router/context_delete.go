package router

import "fmt"

func (c *Context) Delete(modelName string, id int64) {
	model := c.FindModel(modelName)
	sql := fmt.Sprintf("delete from %s where id=$1", model.TableName())
	c.Db.Exec(sql, id)
}
