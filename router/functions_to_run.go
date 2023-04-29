package router

import (
	"fmt"
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
	//auth := util.GetHeader("Authorization", c.Request)

	b := Batch{}
	b.Context = &Context{}
	b.Context.Db = c.Db
	b.Context.Writer = NewBatchWriter()

	request, _ := http.NewRequest("GET", "/", nil)
	//request.Header.Set("Authorization", auth)
	b.Context.Request = request
	b.Context.User = user
	b.Context.Method = "GET"
	b.Context.router = c.router

	tokens := strings.Split(route, "?")
	noParams := tokens[0]
	if len(tokens) == 2 {
		b.Params = tokens[1]
	}
	fmt.Println("noParams", noParams)
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
}

func NewBatchWriter() *BatchWriter {
	b := BatchWriter{}
	b.TheHeader = http.Header{}
	return &b
}

func (w *BatchWriter) WriteHeader(statusCode int) {
	fmt.Println("statusCode", statusCode)
	//w.ResponseWriter.WriteHeader(statusCode)
}

func (w *BatchWriter) Header() http.Header {
	return w.TheHeader
}

func (w *BatchWriter) Write(data []byte) (int, error) {
	//return w.ResponseWriter.Write(data)
	fmt.Println("Write", len(data))
	return 200, nil
}
