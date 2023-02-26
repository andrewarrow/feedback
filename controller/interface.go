package controller

import "net/http"

type InterfaceController interface {
	Index()
	Create()
	HandlePath(http.ResponseWriter, string, string, Vars)
}
