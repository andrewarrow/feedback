package wasm

import (
	"encoding/json"
	"net/url"
	"syscall/js"
)

type AutoForm struct {
	ReturnPath string
	Path       string
	Before     func() string
	After      func(string)
	Id         string
	Method     string
}

func NewAutoForm(id string) *AutoForm {
	a := AutoForm{}
	a.Id = id
	a.Method = "POST"
	return &a
}

func (g *Global) AddAutoForm(a *AutoForm) {
	form := g.Document.Id(a.Id)
	thefunc := func(this js.Value, p []js.Value) any {
		p[0].Call("preventDefault")
		q := ""
		if a.Before != nil {
			q = a.Before()
		}
		go a.Post(g, form, q)
		return nil
	}
	form.JValue.Set("onsubmit", js.FuncOf(thefunc))
}

func (a *AutoForm) Post(g *Global, w *Wrapper, q string) {
	var jsonString string
	var code int
	if a.Method == "POST" {
		jsonString, code = DoPost(a.Path, w.MapOfInputs())
	} else if a.Method == "PATCH" {
		jsonString, code = DoPatch(a.Path, w.MapOfInputs())
	} else if a.Method == "GET" {
		qValue := url.QueryEscape(q)
		jsonString = DoGet(a.Path + "?q=" + qValue)
		if jsonString == "" {
			code = 200
		}
	}
	var m map[string]any
	json.Unmarshal([]byte(jsonString), &m)
	if code == 200 {
		if a.After != nil {
			val, _ := m["val"].(string)
			a.After(val)
			return
		}
		g.Location.Set("href", a.ReturnPath)
		return
	}
	errorString, _ := m["error"].(string)
	g.flashThree("error: " + errorString)
}
