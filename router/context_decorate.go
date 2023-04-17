package router

import (
	"strings"
)

type MSA map[string]any
type MSAS map[string][]string
type MSMAB map[string]map[any]bool
type MI64MSA map[int64]map[string]any

func (c *Context) DecorateSingle(item map[string]any) {
	list := []map[string]any{item}
	c.Decorate(list)
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
