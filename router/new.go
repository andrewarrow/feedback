package router

import (
	"encoding/json"

	"github.com/andrewarrow/feedback/controller"
	"github.com/andrewarrow/feedback/files"
	"github.com/andrewarrow/feedback/persist"
)

type Router struct {
	Paths    map[string]controller.InterfaceController
	Site     Site
	Database persist.Database
}

func NewRouter() *Router {
	r := Router{}
	r.Paths = map[string]controller.InterfaceController{}
	r.Database = persist.NewInMemory()

	jsonString := files.ReadFile("data/site.json")
	json.Unmarshal([]byte(jsonString), &r.Site)

	r.Paths["models"] = controller.NewModelsController()
	//for _, model := range r.Site.Models {
	//	r.Paths[fmt.Sprintf("/admin/%s", util.Plural(model.Name))] = "GET"
	//}

	return &r
}
