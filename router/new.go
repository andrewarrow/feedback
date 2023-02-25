package router

import (
	"encoding/json"

	"github.com/andrewarrow/feedback/files"
	"github.com/andrewarrow/feedback/persist"
)

type Router struct {
	Paths    map[string]string
	Site     Site
	Database persist.Database
}

func NewRouter() *Router {
	r := Router{}
	r.Paths = map[string]string{}
	r.Paths["/admin/users"] = "GET"

	r.Database = persist.NewInMemory()

	jsonString := files.ReadFile("data/site.json")
	json.Unmarshal([]byte(jsonString), &r.Site)
	return &r
}
