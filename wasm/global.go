package wasm

import (
	"strings"
	"syscall/js"
	"text/template"
)

var CustomFuncMap *template.FuncMap

type Global struct {
	Global       *js.Value
	Document     *Document
	Navigator    *Wrapper
	LocalStorage *Wrapper
	Window       *Wrapper
	Location     *Location
	Start        string
	Ready        chan bool
	Space        map[string]string
	Storage      map[string]any
	Stack        []*StackItem
}

func NewGlobal() (*Global, *Document) {
	g := Global{}
	g.Stack = []*StackItem{}
	g.Space = map[string]string{}
	g.Storage = map[string]any{}
	g.Ready = make(chan bool, 1)
	temp := js.Global()
	temp.Set("WasmReady", js.FuncOf(g.WasmReady))
	g.Global = &temp
	g.Document = NewDocument(&g)
	return &g, g.Document
}

func (g *Global) LastUrlToken() string {
	tokens := strings.Split(g.Location.Href, "/")
	return tokens[len(tokens)-1]
}

func (g *Global) WasmReady(this js.Value, p []js.Value) any {
	g.Location = NewLocation(g)
	g.LocalStorage = NewWrapper(g.Global.Get("localStorage"))
	g.Navigator = NewWrapper(g.Global.Get("navigator"))
	g.Window = NewWrapper(g.Global.Get("window"))
	g.Start = p[0].String()
	g.Ready <- true
	return nil
}

func (g *Global) Event(id string, fn func()) {
	button := g.Document.ById(id)
	if button.IsNull() {
		return
	}
	thefunc := func(this js.Value, p []js.Value) any {
		p[0].Call("preventDefault")
		fn()
		return nil
	}
	button.Set("onclick", js.FuncOf(thefunc))
}
func (g *Global) EventWithTarget(id string, fn func(target string)) {
	button := g.Document.ById(id)
	if button.IsNull() {
		return
	}
	thefunc := func(this js.Value, p []js.Value) any {
		p[0].Call("preventDefault")
		fn(p[0].Get("target").Get("id").String())
		return nil
	}
	button.Set("onclick", js.FuncOf(thefunc))
}

func (g *Global) SetClipboard(s string) {
	g.Navigator.JValue.Get("clipboard").Call("writeText", s)
}

func (g *Global) Click(id string, fn func(js.Value, []js.Value) any) {
	button := g.Document.ById(id)
	button.Set("onclick", js.FuncOf(fn))
}
func (g *Global) Submit(id string, fn func(js.Value, []js.Value) any) {
	form := g.Document.ById(id)
	form.Set("onsubmit", js.FuncOf(fn))
}
func (g *Global) SubmitEvent(id string, fn func()) {
	thefunc := func(this js.Value, p []js.Value) any {
		p[0].Call("preventDefault")
		fn()
		return nil
	}
	form := g.Document.ById(id)
	form.Set("onsubmit", js.FuncOf(thefunc))
}
func (g *Global) Focus(id string, fn func(js.Value, []js.Value) any) {
	form := g.Document.ById(id)
	form.Set("onfocus", js.FuncOf(fn))
}
func (g *Global) Change(id string, fn func(js.Value, []js.Value) any) {
	form := g.Document.ById(id)
	if !form.IsNull() {
		form.Set("onchange", js.FuncOf(fn))
	}
}

func (g *Global) Get(id string) string {
	w := g.Global.Get("window")
	return w.Get(id).String()
}
func (g *Global) SetWindowFunc(id string, fn func(id string)) {
	thefunc := func(this js.Value, p []js.Value) any {
		fn(p[0].String())
		return nil
	}
	g.Global.Set(id, js.FuncOf(thefunc))
}
