package wasm

import (
	"syscall/js"
)

type Wrapper struct {
	JValue  js.Value
	Name    string
	Id      string
	Value   string
	Checked bool
}

func NewWrapper(v js.Value) *Wrapper {
	if v.IsNull() {
		return nil
	}
	w := Wrapper{}
	w.JValue = v
	return &w
}

func (w *Wrapper) Set(s string, thing any) {
	if w == nil {
		return
	}
	if w.JValue.IsNull() {
		return
	}
	if thing == nil {
		return
	}
	thingS, ok := thing.(string)
	if ok {
		w.JValue.Set(s, thingS)
		return
	}
	thingB, ok := thing.(bool)
	if ok {
		w.JValue.Set(s, thingB)
		return
	}
	w.JValue.Set(s, js.FuncOf(thing.(func(this js.Value, args []js.Value) any)))
}

func (w *Wrapper) GetAttribute(s string) string {
	return w.JValue.Call("getAttribute", s).String()
}
func (w *Wrapper) Get(s string) string {
	return w.JValue.Get(s).String()
}
func (w *Wrapper) GetInt(s string) int {
	return w.JValue.Get(s).Int()
}
func (w *Wrapper) GetBool(s string) bool {
	return w.JValue.Get(s).Bool()
}
func (w *Wrapper) Call(s string, thing any) any {
	return w.JValue.Call(s, thing)
}
func (w *Wrapper) Focus() {
	w.JValue.Call("focus")
}
func (w *Wrapper) Blur() {
	w.JValue.Call("blur")
}
func (w *Wrapper) AppendChild(c any) {
	w.JValue.Call("appendChild", c)
}
func (w *Wrapper) FireClick() {
	w.JValue.Call("click")
}
func (w *Wrapper) FireSubmit() {
	w.JValue.Call("submit")
}
func (w *Wrapper) IsNull() bool {
	return w.JValue.IsNull()
}
func (w *Wrapper) SetItem(key, value any) {
	w.JValue.Call("setItem", key, value)
}
func (w *Wrapper) GetItem(key any) string {
	item := w.JValue.Call("getItem", key)
	if item.IsNull() {
		return ""
	}
	return w.JValue.Call("getItem", key).String()
}

func (w *Wrapper) Click(fn func(js.Value, []js.Value) any) {
	w.JValue.Set("onclick", js.FuncOf(fn))
}
func (w *Wrapper) Event(fn func(id string)) {
	thefunc := func(this js.Value, p []js.Value) any {
		p[0].Call("preventDefault")
		fn(p[0].Get("target").Get("id").String())
		return nil
	}
	w.JValue.Set("onclick", js.FuncOf(thefunc))
}

func (w *Wrapper) EventWithId(fn func()) {
	thefunc := func(this js.Value, p []js.Value) any {
		fn()
		return nil
	}
	w.JValue.Set("onclick", js.FuncOf(thefunc))
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
func (w *Wrapper) HasClass(c string) bool {
	return HasClass(w.JValue, c)
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
		if item.Get("checked").IsUndefined() == false {
			w.Checked = item.Get("checked").Bool()
		}
		items = append(items, w)
	}
	return items
}

func (w *Wrapper) SelectAll(s string) []*Wrapper {
	return w.SelectAllByQuery("querySelectorAll", s)
}
