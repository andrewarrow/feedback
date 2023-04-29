package router

import (
	"net/http"
	"strings"
)

type Batch struct {
	TheFunc func(*Context, string, string)
	Context *Context
	Second  string
	Third   string
	Params  string
}

func (c *Context) FunctionToRun(route string, user map[string]any) *Batch {
	b := Batch{}
	b.Context = &Context{}
	b.Context.Db = c.Db
	b.Context.Writer = NewBatchWriter(route)

	tokens := strings.Split(route, "?")
	noParams := tokens[0]
	if len(tokens) == 2 {
		b.Params = "?" + tokens[1]
	}
	request, _ := http.NewRequest("GET", "/"+b.Params, nil)
	b.Context.Request = request
	b.Context.User = user
	b.Context.Method = "GET"
	b.Context.router = c.router

	b.Context.tokens = strings.Split(noParams+"/", "/")
	first := b.Context.tokens[1]

	second := ""
	third := ""
	if len(b.Context.tokens) == 4 {
		second = b.Context.tokens[2]
	} else if len(b.Context.tokens) == 5 {
		second = b.Context.tokens[2]
		third = b.Context.tokens[3]
	}
	b.Second = second
	b.Third = third
	b.TheFunc = c.router.pathFuncToRun(first)
	return &b
}

type BatchWriter struct {
	http.ResponseWriter
	TheHeader http.Header
	Route     string
	Results   map[string][]byte
	Code      int
}

func NewBatchWriter(route string) *BatchWriter {
	b := BatchWriter{}
	b.TheHeader = http.Header{}
	b.Route = route
	b.Results = map[string][]byte{}
	return &b
}

func (w *BatchWriter) WriteHeader(statusCode int) {
	w.Code = statusCode
}

func (w *BatchWriter) Header() http.Header {
	return w.TheHeader
}

func (w *BatchWriter) Write(data []byte) (int, error) {
	w.Results[w.Route] = data
	return 200, nil
}
