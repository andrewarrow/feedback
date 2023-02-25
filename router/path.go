package router

import (
	"fmt"
	"net/http"

	"github.com/andrewarrow/feedback/files"
)

func (r *Router) RouteFromRequest(writer http.ResponseWriter, request *http.Request) {
	path := request.URL.Path
	fmt.Println(path)
	if path == "/" {
		welcome := files.ReadFile("views/welcome.html")
		fmt.Fprintf(writer, welcome)
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
