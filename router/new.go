package router

import (
	"html/template"

	"github.com/andrewarrow/feedback/controller"
	"github.com/andrewarrow/feedback/persist"
)

type Router struct {
	Paths    map[string]controller.InterfaceController
	Database persist.Database
	Template *template.Template
	Vars     *controller.Vars
}

func NewRouter() *Router {
	r := Router{}
	r.Paths = map[string]controller.InterfaceController{}
	r.Database = persist.NewInMemory()

	r.Template = LoadTemplates()
	render := controller.NewRender(r.Template)
	r.Vars = render.Vars
	r.Paths["/models"] = controller.NewModelsController(render)
	//for _, model := range r.Site.Models {
	//	r.Paths[fmt.Sprintf("/admin/%s", util.Plural(model.Name))] = "GET"
	//}

	return &r
}
