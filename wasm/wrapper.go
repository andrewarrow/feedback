package wasm

import (
	"fmt"
	"syscall/js"
)

type Wrapper struct {
	JValue js.Value
	Name   string
	Id     string
	Value  string
}

func NewWrapper(v js.Value) *Wrapper {
	w := Wrapper{}
	w.JValue = v
	return &w
}

func (w *Wrapper) SelectAll(s string) []*Wrapper {
	list := w.JValue.Call("querySelectorAll", s)
	items := []*Wrapper{}
	for i := 0; i < list.Length(); i++ {
		item := list.Index(i)
		w := NewWrapper(item)
		w.Name = item.Get("name").String()
		w.Id = item.Get("id").String()
		w.Value = item.Get("value").String()
		fmt.Println(w)
		items = append(items, w)
	}
	return items
}
