package router

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
)

func (r *Router) ListenAndServe(port string) {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		r.RouteFromRequest(writer, request)
	})

	log.Fatal(http.ListenAndServe(port, nil))
}

func (r *Router) ListenAndServeTLS() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		r.RouteFromRequest(writer, request)
	})

	server := &http.Server{
		Addr: ":443",
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	err := server.ListenAndServeTLS("cert.pem", "key.pem")
	fmt.Println(err)
	select {}

}
