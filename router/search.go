package router

import (
	"fmt"
	"strings"
)

func AddSearchResults(c *Context, field, token string, allRows map[int64]any) {
	rows := c.SelectAll("user", fmt.Sprintf("where LOWER(%s) like $1", field),
		[]any{"%" + strings.ToLower(token) + "%"}, "")
	for _, row := range rows {
		id := row["id"].(int64)
		allRows[id] = row
	}
}
