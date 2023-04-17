package router

import (
	"fmt"
	"strings"
)

func (c *Context) Decorate(list []map[string]any) {
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
	fmt.Println(ids)
}
