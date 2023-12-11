package wasm

import (
	"fmt"
	"syscall/js"
)

type Settable interface {
	Set(id, value string)
}

type Location struct {
	Value  js.Value
	href   string
	Params js.Value
}

func NewLocation(g *Global) *Location {
	l := Location{}
	l.Value = g.Global.Get("location")
	l.href = l.Value.Get("href").String()
	l.Params = g.Global.Get("URLSearchParams").New(l.href)
	return &l
}

func (l *Location) GetParam(id string) string {
	return l.Params.Call("get", id).String()
}

func (l *Location) Set(id, value string) {
	fmt.Println("!1")
	l.Value.Set(id, value)
}
