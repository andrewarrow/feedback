package router

import (
	"fmt"
	"net/http"

	"github.com/andrewarrow/feedback/files"
)

func (r *Router) RouteFromRequest(writer http.ResponseWriter, request *http.Request) {
	path := request.URL.Path
	if path == "/" {
		welcome := files.ReadFile("views/welcome.html")
		fmt.Fprintf(writer, welcome)
	} else if path == "/feedback/add" {
		fmt.Fprintf(writer, "ok")
	} else {
		writer.WriteHeader(404)
		notFound := files.ReadFile("views/404.html")
		fmt.Fprintf(writer, notFound)
	}
}
