package controller

import (
	"net/http"
)

type SessionsController struct {
	FeedbackController
}

func NewSessionsController(render *Render) *SessionsController {
	s := SessionsController{}
	s.render = render
	return &s
}

func (m *SessionsController) Create() {
	h := m.writer.Header()
	h.Set("Location", "/")
	m.writer.WriteHeader(301)
}

func (m *SessionsController) Index() {
}

func (m *SessionsController) HandlePath(writer http.ResponseWriter,
	request *http.Request, tokens []string) {
}
