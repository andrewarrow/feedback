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

func (c *Context) FunctionToRun(route string) *Batch {
	b := Batch{}
	b.Context = &Context{}
	b.Context.Db = c.Db
	b.Context.Writer = &BatchWriter{}

	request, _ := http.NewRequest("GET", "/", nil)
	b.Context.Request = request

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
}

func (w *BatchWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *BatchWriter) Write(data []byte) (int, error) {
	return w.ResponseWriter.Write(data)
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	BatchWriter := &BatchWriter{ResponseWriter: w}

	BatchWriter.WriteHeader(http.StatusOK)
	BatchWriter.Write([]byte("Hello, world!"))
}
