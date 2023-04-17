package router

import (
	"strings"
)

func (c *Context) Decorate(list []map[string]any, level int) {
	if level > 10 {
		return
	}
	ids := map[string]map[any]bool{}
	for _, item := range list {
		for k, v := range item {
			if strings.HasSuffix(k, "_id") {
				tokens := strings.Split(k, "_")
				modelString := tokens[0]
				if c.FindModel(modelString) == nil {
					continue
				}
				if ids[modelString] == nil {
					ids[modelString] = map[any]bool{}
				}
				ids[modelString][v] = true
			}
		}
	}
	itemMaps := map[string]any{}
	for k, v := range ids {
		whereInList := []any{}
		for kk, _ := range v {
			whereInList = append(whereInList, kk)
		}
		itemMaps[k] = c.WhereIn(k, whereInList)
	}
	for _, item := range list {
		for k, v := range item {
			if strings.HasSuffix(k, "_id") {
				tokens := strings.Split(k, "_")
				modelString := tokens[0]
				if itemMaps[modelString] == nil {
					continue
				}
				intId := v.(int64)
				if intId == 0 {
					continue
				}
				lookup := itemMaps[modelString].(map[int64]map[string]any)
				newList := []map[string]any{lookup[intId]}
				c.Decorate(newList, level+1)
				item[modelString] = newList[0]
			}
		}
	}
}
