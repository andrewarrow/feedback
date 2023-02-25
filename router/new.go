package router

import (
	"encoding/json"

	"github.com/andrewarrow/feedback/files"
)

type Router struct {
	Paths map[string]string
	Site  Site
}

func NewRouter() *Router {
	r := Router{}
	r.Paths = map[string]string{}
	r.Paths["/admin/users"] = "GET"

	jsonString := files.ReadFile("data/site.json")
	json.Unmarshal([]byte(jsonString), &r.Site)
	return &r
}
