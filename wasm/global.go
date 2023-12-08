package wasm

import (
	"fmt"
	"syscall/js"
)

type Global struct {
	Global   *js.Value
	Document *Document
	Location *Location
	Start    string
}

func NewGlobal() (*Global, *Document) {
	g := Global{}
	temp := js.Global()
	temp.Set("WasmReady", js.FuncOf(g.WasmReady))
	g.Global = &temp
	g.Document = NewDocument(&g)
	return &g, g.Document
}

func (g *Global) WasmReady(this js.Value, p []js.Value) any {
	fmt.Println("here")
	g.Location = NewLocation(g)
	g.Start = p[0].String()
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
