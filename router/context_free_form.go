package router

func (c *Context) FreeFormSelect(sql string, params ...any) []map[string]any {
	ms := []map[string]any{}
	rows, err := c.Db.Queryx(sql, params...)
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

func (c *Context) FreeFormUpdate(sql string, params ...any) error {
	_, err := c.Db.Exec(sql, params...)
	if err != nil {
		return err
	}
	return nil
}
