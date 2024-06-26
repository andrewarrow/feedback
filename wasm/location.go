package wasm

import (
	"net/url"
	"strings"
	"syscall/js"
)

type Settable interface {
	Set(id, value string)
}

type Location struct {
	Value  js.Value
	Href   string
	Params url.Values
}

func NewLocation(g *Global) *Location {
	l := Location{}
	l.Value = g.Global.Get("location")
	l.Href = l.Value.Get("href").String()
	tokens := strings.Split(l.Href, "?")
	if len(tokens) == 2 {
		l.Params, _ = url.ParseQuery(tokens[1])
	}
	return &l
}

func (l *Location) Reload() {
	l.Value.Call("reload")
}
func (l *Location) GetParam(id string) string {
	val := l.Params[id]
	if len(val) == 0 {
		return ""
	}
	return val[0]
}

func (l *Location) Set(id, value string) {
	l.Value.Set(id, value)
}
