package wasm

import (
	"syscall/js"
)

type Global struct {
	Global   *js.Value
	Document *Document
	Window   *Wrapper
	Location *Location
	Start    string
	Ready    chan bool
	Space    map[string]string
	Stack    []*StackItem
}

func NewGlobal() (*Global, *Document) {
	g := Global{}
	g.Stack = []*StackItem{}
	g.Space = map[string]string{}
	g.Ready = make(chan bool, 1)
	temp := js.Global()
	temp.Set("WasmReady", js.FuncOf(g.WasmReady))
	g.Global = &temp
	g.Document = NewDocument(&g)
	return &g, g.Document
}

func (g *Global) WasmReady(this js.Value, p []js.Value) any {
	g.Location = NewLocation(g)
	g.Window = NewWrapper(g.Global.Get("window"))
	g.Start = p[0].String()
	g.Ready <- true
	return nil
}

func (g *Global) Click(id string, fn func(js.Value, []js.Value) any) {
	button := g.Document.ById(id)
	button.Set("onclick", js.FuncOf(fn))
}
func (g *Global) Submit(id string, fn func(js.Value, []js.Value) any) {
	form := g.Document.ById(id)
	form.Set("onsubmit", js.FuncOf(fn))
}
