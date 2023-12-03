package wasm

import "syscall/js"

type Document struct {
	Document js.Value
}

func NewDocument() *Document {
	d := Document{}
	d.Document = js.Global().Get("document")
	return &d
}

func (d *Document) ById(id string) js.Value {
	return d.Document.Call("getElementById", id)
}
