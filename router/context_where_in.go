package router

import (
	"fmt"
	"strings"
)

func (c *Context) WhereIn(modelString string, ids []any) MI64MSA {
	stringIds := []string{}
	for _, id := range ids {
		stringIds = append(stringIds, fmt.Sprintf("%d", id))
	}
	sql := fmt.Sprintf("where id in (%s)", strings.Join(stringIds, ","))
	items := c.All(modelString, sql, "")
	resultMap := MI64MSA{}
	for _, row := range items {
		id := row["id"].(int64)
		resultMap[id] = row
	}

	return resultMap
}
