package wasm

import (
	"fmt"
	"syscall/js"
)

type Global struct {
	Global   *js.Value
	Document *Document
	Location *Settable
}

func NewGlobal() *Global {
	g := Global{}
	temp := js.Global()
	temp.Set("WasmReady", js.FuncOf(g.WasmReady))
	g.Global = &temp
	g.Document = NewDocument(&g)
	g.Location = NewLocation(&g)
	return &g
}

func (g *Global) WasmReady(this js.Value, p []js.Value) any {
	fmt.Println("here")
	return nil
}

func (g *Global) Click(id string, fn func(js.Value, []js.Value) any) {
	button := g.Document.ById(id)
	button.Set("onclick", js.FuncOf(fn))
}
