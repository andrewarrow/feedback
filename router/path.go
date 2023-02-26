package router

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/andrewarrow/feedback/files"
)

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
			match.HandlePath(writer, request)
		}
	}
}

func (r *Router) HandleAsset(path string, writer http.ResponseWriter) {
	contentType := "text/css"
	if strings.HasSuffix(path, ".js") {
		contentType = "application/javascript"
	}
	writer.Header().Set("Content-Type", contentType)
	matchFile := files.ReadFile(fmt.Sprintf("%s", path[1:]))
	fmt.Fprintf(writer, matchFile)
}
