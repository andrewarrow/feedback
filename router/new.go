package router

import (
	"encoding/json"
	"html/template"

	"github.com/andrewarrow/feedback/controller"
	"github.com/andrewarrow/feedback/files"
	"github.com/andrewarrow/feedback/persist"
)

type Router struct {
	Paths    map[string]controller.InterfaceController
	Site     Site
	Database persist.Database
	Template *template.Template
	Vars     controller.Vars
}

func NewRouter() *Router {
	r := Router{}
	r.Paths = map[string]controller.InterfaceController{}
	r.Database = persist.NewInMemory()

	jsonString := files.ReadFile("data/site.json")
	json.Unmarshal([]byte(jsonString), &r.Site)

	r.Paths["/models"] = controller.NewModelsController(r.Site.Models)
	//for _, model := range r.Site.Models {
	//	r.Paths[fmt.Sprintf("/admin/%s", util.Plural(model.Name))] = "GET"
	//}
	r.Template = LoadTemplates()
	r.Vars = r.NewVars()

	return &r
}
