package wasm

import (
	"encoding/json"
	"syscall/js"
)

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

func (d *Document) ByIdWrap(id string) *Wrapper {
	return NewWrapper(d.ById(id))
}
func (d *Document) ByIdWrapped(id string) *Wrapper {
	return d.ByIdWrap(id)
}

func (d *Document) SelectAllFrom(id, s string) []*Wrapper {
	return d.ByIdWrap(id).SelectAll(s)
}
func (d *Document) AppendTo(id, jsonString, template string) {
	w := d.ByIdWrap(id)
	var m map[string]any
	json.Unmarshal([]byte(jsonString), &m)
	newHTML := Render(template, m)
	ih := w.Get("innerHTML")
	w.Set("innerHTML", ih+newHTML)
}

func (d *Document) ByIdString(id string) string {
	input := d.Document.Call("getElementById", id)
	return input.Get("value").String()
}
