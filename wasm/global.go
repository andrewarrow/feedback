package wasm

import (
	"fmt"
	"syscall/js"
)

type Global struct {
	Global *js.Value
}

func NewGlobal() *Global {
	g := Global{}
	temp := js.Global()
	temp.Set("WasmReady", js.FuncOf(g.WasmReady))
	g.Global = &temp
	return &g
}

func (g *Global) WasmReady(this js.Value, p []js.Value) any {
	fmt.Println("here")
	return nil
}
