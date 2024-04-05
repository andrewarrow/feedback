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

	fmt.Println("/Users/aa/cert.pem", "/Users/aa/key.pem")
	err := server.ListenAndServeTLS("/Users/aa/cert.pem", "/Users/aa/key.pem")
	fmt.Println(err)
	select {}

}
