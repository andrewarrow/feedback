package router

func (c *Context) TableJson(tableName string) {
	send := map[string]any{}
	items := c.All(tableName, "order by created_at desc", "")
	send["items"] = items
	c.SendContentAsJson(send, 200)
}
func (c *Context) TableJsonParams(tableName, where string, params ...any) {
	send := map[string]any{}
	items := c.All(tableName, where+" order by created_at desc", "", params...)
	send["items"] = items
	c.SendContentAsJson(send, 200)
}
