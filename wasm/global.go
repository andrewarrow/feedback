package wasm

import (
	"fmt"
	"syscall/js"
)

type Global struct {
	Global   *js.Value
	Document *Document
}

func NewGlobal() *Global {
	g := Global{}
	temp := js.Global()
	temp.Set("WasmReady", js.FuncOf(g.WasmReady))
	g.Global = &temp
	g.Document = NewDocument(&g)
	return &g
}

func (g *Global) WasmReady(this js.Value, p []js.Value) any {
	fmt.Println("here")
	return nil
}

func (g *Global) Click(id string, fn func(js.Value, []js.Value)) {
	button := g.ById(id)
	button.Set("onclick", js.FuncOf(fn))
}
