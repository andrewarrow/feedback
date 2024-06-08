package wasm

import (
	"encoding/json"
	"syscall/js"
)

type AutoForm struct {
	ReturnPath string
	Path       string
	After      func(int64)
	Id         string
}

func NewAutoForm(id string) *AutoForm {
	a := AutoForm{}
	a.Id = id
	return &a
}

func (g *Global) AddAutoForm(a *AutoForm) {
	form := g.Document.Id(a.Id)
	thefunc := func(this js.Value, p []js.Value) any {
		p[0].Call("preventDefault")
		go a.Post(g, form)
		return nil
	}
	form.JValue.Set("onsubmit", js.FuncOf(thefunc))
}

func (a *AutoForm) Post(g *Global, w *Wrapper) {
	jsonString, code := DoPost(a.Path, w.MapOfInputs())
	var m map[string]any
	json.Unmarshal([]byte(jsonString), &m)
	if code == 200 {
		if a.After != nil {
			id, _ := m["id"].(float64)
			a.After(int64(id))
			return
		}
		g.Location.Set("href", a.ReturnPath)
		return
	}
	errorString, _ := m["error"].(string)
	g.flashThree("error: " + errorString)
}
