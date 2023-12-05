package wasm

import (
	"fmt"
	"syscall/js"
)

type Settable interface {
	Set(id, value string)
}

type Location struct {
	Value js.Value
}

func NewLocation(g *Global) *Location {
	l := Location{}
	l.Value = g.Global.Get("location")
	return &l
}

func (l *Location) Set(id, value string) {
	fmt.Println("!1")
	l.Value.Set(id, value)
}
