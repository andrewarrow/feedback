package wasm

import (
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

func (w *Wrapper) Set(s string, thing any) {
	thingS, ok := thing.(string)
	if ok {
		w.JValue.Set(s, thingS)
		return
	}
	w.JValue.Set(s, js.FuncOf(thing.(func(this js.Value, args []js.Value) any)))
}

func (w *Wrapper) Get(s string) string {
	return w.JValue.Get(s).String()
}
func (w *Wrapper) GetInt(s string) int {
	return w.JValue.Get(s).Int()
}
func (w *Wrapper) Call(s string, thing any) any {
	return w.JValue.Call(s, thing)
}
func (w *Wrapper) Focus() {
	w.JValue.Call("focus")
}
func (w *Wrapper) AppendChild(c any) {
	w.JValue.Call("appendChild", c)
}
func (w *Wrapper) FireClick() {
	w.JValue.Call("click")
}
func (w *Wrapper) SetItem(key, value any) {
	w.JValue.Call("setItem", key, value)
}
func (w *Wrapper) GetItem(key any) string {
	return w.JValue.Call("getItem", key).String()
}

func (w *Wrapper) Click(fn func(js.Value, []js.Value) any) {
	w.JValue.Set("onclick", js.FuncOf(fn))
}

func (w *Wrapper) Show() {
	RemoveClass(w.JValue, "hidden")
}
func (w *Wrapper) Hide() {
	AddClass(w.JValue, "hidden")
}
func (w *Wrapper) AddClass(c string) {
	AddClass(w.JValue, c)
}
func (w *Wrapper) RemoveClass(c string) {
	RemoveClass(w.JValue, c)
}

func (w *Wrapper) SelectAllByClass(s string) []*Wrapper {
	return w.SelectAllByQuery("getElementsByClassName", s)
}

func (w *Wrapper) SelectAllByQuery(call, s string) []*Wrapper {
	list := w.JValue.Call(call, s)
	items := []*Wrapper{}
	for i := 0; i < list.Length(); i++ {
		item := list.Index(i)
		w := NewWrapper(item)
		w.Name = item.Get("name").String()
		w.Id = item.Get("id").String()
		w.Value = item.Get("value").String()
		items = append(items, w)
	}
	return items
}

func (w *Wrapper) SelectAll(s string) []*Wrapper {
	return w.SelectAllByQuery("querySelectorAll", s)
}
