package router

import (
	"fmt"
	"net/http"

	"github.com/andrewarrow/feedback/files"
)

func RouteFromRequest(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		welcome := files.ReadFile("views/welcome.html")
		fmt.Fprintf(w, welcome)
	} else {
		w.WriteHeader(404)
		notFound := files.ReadFile("views/404.html")
		fmt.Fprintf(w, notFound)
	}
}
