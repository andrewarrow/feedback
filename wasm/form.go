package wasm

import (
	"encoding/json"
	"strings"
	"syscall/js"
	"time"
)

func (w *Wrapper) MapOfInputs() map[string]any {
	m := map[string]any{}
	for _, input := range w.SelectAll("input") {
		if input.Get("type") == "submit" {
			continue
		}
		m[input.Id] = input.Value
		input.Set("value", "")
	}
	for _, input := range w.SelectAll("textarea") {
		m[input.Id] = input.Value
		input.Set("value", "")
	}
	for _, input := range w.SelectAll("select") {
		m[input.Id] = input.Value
		input.Set("value", "")
	}
	for _, input := range w.SelectAll("hidden") {
		m[input.Id] = input.Value
	}
	return m
}

func (w *Wrapper) NoClearInputs(prefix string) map[string]any {
	m := map[string]any{}
	for _, input := range w.SelectAll("input") {
		if input.Get("type") == "submit" {
			continue
		}
		input.Id = input.Id[len(prefix):]
		if input.Get("type") == "checkbox" {
			m[input.Id] = input.Checked
		} else {
			m[input.Id] = strings.TrimSpace(input.Value)
		}
	}
	for _, input := range w.SelectAll("textarea") {
		input.Id = input.Id[len(prefix):]
		m[input.Id] = strings.TrimSpace(input.Value)
	}
	for _, input := range w.SelectAll("select") {
		input.Id = input.Id[len(prefix):]
		m[input.Id] = input.Value
	}
	for _, input := range w.SelectAll("hidden") {
		input.Id = input.Id[len(prefix):]
		m[input.Id] = input.Value
	}
	return m
}

func (g *Global) AutoForm(id, after string, cb func()) {
	form := g.Document.Id(id)
	thefunc := func(this js.Value, p []js.Value) any {
		p[0].Call("preventDefault")
		go form.AutoFormPost(g, id, after, cb)
		return nil
	}
	form.JValue.Set("onsubmit", js.FuncOf(thefunc))
}

func (w *Wrapper) AutoFormPost(g *Global, id, after string, cb func()) {
	jsonString, code := DoPost("/"+after+"/"+id, w.MapOfInputs())
	var m map[string]any
	json.Unmarshal([]byte(jsonString), &m)
	if code == 200 {
		returnPath, _ := m["return"].(string)
		if returnPath == "" {
			returnPath = after
		}
		if cb != nil {
			cb()
			return
		}
		g.Location.Set("href", returnPath)
		return
	}
	errorString, _ := m["error"].(string)
	g.flashThree("error: " + errorString)
}

func (g *Global) AutoDel(route string, w *Wrapper, name string, cb func()) {
	for _, item := range w.SelectAllByClass(name) {
		lid := item.Id[2:]
		click := func() {
			go func() {
				DoDelete(route + lid)
				cb()
			}()
		}
		item.EventWithId(click)
	}
}

func (g *Global) flashThree(s string) {
	flash := g.Document.ById("flash")
	flash.Set("innerHTML", s)
	time.Sleep(time.Second * 3)
	flash.Set("innerHTML", "")
}
