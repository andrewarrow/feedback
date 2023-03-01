package router

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/andrewarrow/feedback/controller"
	"github.com/andrewarrow/feedback/files"
)

type Router struct {
	Paths    map[string]func(c *Context)
	Template *template.Template
	Vars     *controller.Vars
	Site     *controller.Site
}

type Context struct {
	writer  http.ResponseWriter
	request *http.Request
	tokens  []string
}

func NewRouter(path string) *Router {
	r := Router{}
	r.Paths = map[string]func(c *Context){}

	var site controller.Site
	jsonString := files.ReadFile(path)
	json.Unmarshal([]byte(jsonString), &site)
	r.Site = &site

	r.Vars = controller.NewVars(&site)
	r.Template = LoadTemplates()
	r.Paths["models"] = r.ModelsResource
	r.Paths["sessions"] = r.SessionsResource

	return &r
}
