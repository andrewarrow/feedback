package router

import (
	"encoding/json"
	"fmt"

	"github.com/andrewarrow/feedback/files"
	"github.com/andrewarrow/feedback/persist"
	"github.com/andrewarrow/feedback/util"
)

type Router struct {
	Paths    map[string]string
	Site     Site
	Database persist.Database
}

func NewRouter() *Router {
	r := Router{}
	r.Paths = map[string]string{}
	r.Database = persist.NewInMemory()

	jsonString := files.ReadFile("data/site.json")
	json.Unmarshal([]byte(jsonString), &r.Site)

	for _, model := range r.Site.Models {
		r.Paths[fmt.Sprintf("/admin/%s", util.Plural(model.Name))] = "GET"
	}

	return &r
}
