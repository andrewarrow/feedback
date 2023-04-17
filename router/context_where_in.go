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

func (c *Context) WhereInMap(modelString string, rows any, field, otherField string) (map[int64]map[string]any, []string) {
	resultMap := map[int64]map[string]any{}
	list := []string{}
	ids := []string{}

	if field != "" {
		for _, row := range rows.([]map[string]any) {
			ids = append(ids, fmt.Sprintf("%d", row[field]))
		}
	} else {
		for _, s := range rows.([]string) {
			ids = append(ids, s)
		}
	}

	sql := fmt.Sprintf("where id in (%s)", strings.Join(ids, ","))
	items := c.SelectAll(modelString, sql, []any{}, "")
	for _, row := range items {
		id := row["id"].(int64)
		resultMap[id] = row
		if otherField != "" {
			list = append(list, fmt.Sprintf("%d", row[otherField]))
		}
	}

	return resultMap, list
}

func (c *Context) WhereInWithId(modelString string, rows any, field, id string) map[int64]map[string]any {
	resultMap := map[int64]map[string]any{}
	ids := []string{}

	for _, row := range rows.([]map[string]any) {
		ids = append(ids, fmt.Sprintf("%d", row[field]))
	}

	sql := fmt.Sprintf("where %s in (%s)", id, strings.Join(ids, ","))
	items := c.SelectAll(modelString, sql, []any{}, "")
	for _, row := range items {
		id := row["id"].(int64)
		resultMap[id] = row
	}

	return resultMap
}
