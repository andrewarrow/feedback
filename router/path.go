package router

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/andrewarrow/feedback/controller"
	"github.com/andrewarrow/feedback/files"
)

func (r *Router) NewVars() controller.Vars {
	vars := controller.NewVars()
	vars.Phone = r.Site.Phone
	return vars
}

func (r *Router) RouteFromRequest(writer http.ResponseWriter, request *http.Request) {
	path := request.URL.Path
	fmt.Println(path)
	if path == "/" {
		r.Template.ExecuteTemplate(writer, "welcome.html", r.Vars)
	} else if strings.HasPrefix(path, "/assets") {
		r.HandleAsset(path, writer)
	} else if path == "/feedback/add" {
		fmt.Fprintf(writer, "ok")
	} else {
		match := r.Paths[path]
		if match == nil {
			writer.WriteHeader(404)
			r.Template.ExecuteTemplate(writer, "404.html", r.Vars)
		} else {
			match.HandlePath(writer, request, r.Vars)
		}
	}
}

func (r *Router) HandleAsset(path string, writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "text/css")
	matchFile := files.ReadFile(fmt.Sprintf("%s", path[1:]))
	fmt.Fprintf(writer, matchFile)
}
