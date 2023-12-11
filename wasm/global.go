package wasm

import (
	"fmt"
	"syscall/js"
)

type Global struct {
	Global   *js.Value
	Document *Document
	Window   *Wrapper
	Location *Location
	Start    string
	Ready    chan bool
	Space    map[string]any
}

func NewGlobal() (*Global, *Document) {
	g := Global{}
	g.Space = map[string]any{}
	g.Ready = make(chan bool, 1)
	temp := js.Global()
	temp.Set("WasmReady", js.FuncOf(g.WasmReady))
	g.Global = &temp
	g.Document = NewDocument(&g)
	return &g, g.Document
}

func (g *Global) WasmReady(this js.Value, p []js.Value) any {
	fmt.Println("here")
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
