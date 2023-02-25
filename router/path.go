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
		welcome := files.ReadFile("views/welcome.html")
		fmt.Fprintf(writer, welcome)
	} else if strings.HasPrefix(path, "/assets") {
		r.HandleAsset(path, writer)
	} else if path == "/feedback/add" {
		fmt.Fprintf(writer, "ok")
	} else {
		match := r.Paths[path]
		if match == "" {
			writer.WriteHeader(404)
			notFound := files.ReadFile("views/404.html")
			fmt.Fprintf(writer, notFound)
		} else {
			matchFile := files.ReadFile(fmt.Sprintf("views%s.html", path))
			fmt.Fprintf(writer, matchFile)
		}
	}
}

func (r *Router) HandleAsset(path string, writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "text/css")
	matchFile := files.ReadFile(fmt.Sprintf("%s", path[1:]))
	fmt.Fprintf(writer, matchFile)
}
