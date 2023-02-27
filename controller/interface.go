package controller

import "net/http"

type InterfaceController interface {
	Index()
	Create()
	HandlePath(http.ResponseWriter, *http.Request, []string)
}

type FeedbackController struct {
	render *Render
	writer http.ResponseWriter
}
