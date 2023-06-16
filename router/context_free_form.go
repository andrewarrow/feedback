package router

import "fmt"

func (c *Context) FreeFormSelect(sql string, params ...any) []map[string]any {
	return c.Router.FreeFormSelect(sql, params...)
}

func (r *Router) FreeFormSelect(sql string, params ...any) []map[string]any {
	ms := []map[string]any{}
	rows, err := r.Db.Queryx(sql, params...)
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

func (c *Context) FreeFormUpdate(sql string, params ...any) error {
	_, err := c.Db.Exec(sql, params...)
	if err != nil {
		return err
	}
	return nil
}

func (r *Router) FreeFormUpdate(sql string, params ...any) error {
	_, err := r.Db.Exec(sql, params...)
	if err != nil {
		return err
	}
	return nil
}
