package router

import (
	"log"
	"net/http"
)

func (r *Router) ListenAndServe(port string) {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		r.RouteFromRequest(writer, request)
	})

	log.Fatal(http.ListenAndServe(port, nil))
}
