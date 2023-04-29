package router

import (
	"strings"
)

type MSA map[string]any
type MSAS map[string][]string
type MSMAB map[string]map[any]bool
type MI64MSA map[int64]map[string]any

func (c *Context) DecorateListWithFields(list []map[string]any,
	fields map[string]bool) {
	c.DecorateList(list)
	for _, item := range list {
		for k, v := range item {
			m, isMap := v.(map[string]any)
			if isMap == false {
				continue
			}
			handleMap(k, m, fields, 0)
		}
	}
}

func handleMap(k string, m map[string]any, fields map[string]bool, level int) {
	if level > 10 {
		return
	}
	for k, v := range m {
		if fields[k] == false {
			delete(m, k)
		}
		otherMap, isMap := v.(map[string]any)
		if isMap {
			handleMap(k, otherMap, fields, level+1)
		}
	}

}

func (c *Context) DecorateSingle(item map[string]any) {
	list := []map[string]any{item}
	c.DecorateList(list)
}

func (c *Context) DecorateList(list []map[string]any) {
	topLevel := c.Decorate(list)
	for _, modelString := range topLevel {
		thingList := []map[string]any{}
		for _, item := range list {
			thing := item[modelString]
			if thing != nil {
				thingList = append(thingList, thing.(map[string]any))
			}
		}
		c.Decorate(thingList)
	}
}

func (c *Context) Decorate(list []map[string]any) []string {
	topLevel := []string{}
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
					topLevel = append(topLevel, modelString)
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
				lookup := itemMaps[modelString].(MI64MSA)
				item[modelString] = lookup[intId]
			}
		}
	}
	return topLevel
}
