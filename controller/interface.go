package controller

import "net/http"

type InterfaceController interface {
	Index()
	Create()
	HandlePath(http.ResponseWriter, *http.Request, Vars)
}
