package router

import (
	"fmt"
	"strconv"
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

func (c *Context) WhereInWithId(modelString, id string, ids []any) map[int64]map[string]any {
	stringIds := []string{}
	for _, id := range ids {
		stringIds = append(stringIds, fmt.Sprintf("%d", id))
	}
	sql := fmt.Sprintf("where %s in (%s)", id, strings.Join(stringIds, ","))
	items := c.All(modelString, sql, "")
	resultMap := map[int64]map[string]any{}
	for _, row := range items {
		id := row[id].(int64)
		resultMap[id] = row
	}

	return resultMap
}

/*

select coach_id, tag from coach_tags where tag in ('test2','test4') order by created_at desc limit 30;
*/

func listSizeOf(size int) string {
	buffer := []string{}
	for i := 1; i < size+1; i++ {
		buffer = append(buffer, fmt.Sprintf("$%d", i))
	}

	return strings.Join(buffer, ",")
}

func (c *Context) AllIn(field, modelName string, offset, other string, tokens []any) []map[string]any {
	model := c.FindModel(modelName)
	offsetString := ""
	if offset != "" {
		offsetInt, _ := strconv.Atoi(offset)
		offsetString = fmt.Sprintf("OFFSET %d", offsetInt)
	}
	sql := fmt.Sprintf("select distinct(%s) from %s where %s in (%s) order by %s desc limit 30 %s", field, model.TableName(), other, listSizeOf(len(tokens)), field, offsetString)
	ms := []map[string]any{}
	rows, err := c.Db.Queryx(sql, tokens...)
	if err != nil {
		return ms
	}
	defer rows.Close()
	for rows.Next() {
		m := make(map[string]any)
		rows.MapScan(m)
		CastFields(model, m)
		ms = append(ms, m)
	}
	return ms
}

func (c *Context) WhereInList(modelString string, ids []any) []map[string]any {
	stringIds := []string{}
	for _, id := range ids {
		stringIds = append(stringIds, fmt.Sprintf("%d", id))
	}
	sql := fmt.Sprintf("where id in (%s)", strings.Join(stringIds, ","))
	return c.All(modelString, sql, "")
}
