package router

import (
	"fmt"
	"strings"
)

type MSA map[string]any
type MSAS map[string][]string
type MSMAB map[string]map[any]bool

func (c *Context) DecorateSingle(item map[string]any) {
	list := []map[string]any{item}
	c.Decorate(list, 0)
}

func (c *Context) DecorateList(list []map[string]any) {
	c.Decorate(list, 0)
}

func gatherDecorateIds(list []MSA, fill MSAS, level int) {
	if level > 10 {
		return
	}
	ids := MSMAB{}
	for _, item := range list {
		for k, v := range item {
			if strings.HasSuffix(k, "_id") == false {
				continue
			}
			tokens := strings.Split(k, "_")
			modelString := tokens[0]
			if ids[modelString] == nil {
				ids[modelString] = map[any]bool{}
			}
			ids[modelString][v] = true
		}
	}
	itemMaps := MSAS{}
	for k, v := range ids {
		whereInList := []string{}
		for kk, _ := range v {
			intId := kk.(int64)
			whereInList = append(whereInList, fmt.Sprintf("%d", intId))
		}
		itemMaps[k] = whereInList
	}
	fmt.Println(itemMaps)
	return
}

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
