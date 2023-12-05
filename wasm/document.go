package wasm

import "syscall/js"

type Document struct {
	Document js.Value
}

func NewDocument(g *Global) *Document {
	d := Document{}
	d.Document = g.Global.Get("document")
	return &d
}

func (d *Document) ById(id string) js.Value {
	return d.Document.Call("getElementById", id)
}

func (d *Document) ByIdWrapped(id string) *Wrapper {
	return NewWrapper(d.ById(id))
}
