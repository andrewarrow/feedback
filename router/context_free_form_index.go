package router

import (
	"fmt"
	"os"
)

func (c *Context) UpdateWithIndex(index int64, sql string, params ...any) error {
	if os.Getenv("DEBUG") == "1" {
		fmt.Println("UpdateWithIndex", sql, params)
	}
	_, err := c.Dbs[index].Exec(sql, params...)
	if err != nil {
		return err
	}
	return nil
}
