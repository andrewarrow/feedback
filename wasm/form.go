package wasm

import (
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

func (g *Global) AutoForm(id, after string) {
	form := g.Document.Id(id)
	thefunc := func(this js.Value, p []js.Value) any {
		p[0].Call("preventDefault")
		go form.AutoFormPost(g, id, after)
		return nil
	}
	form.JValue.Set("onsubmit", js.FuncOf(thefunc))
}

func (w *Wrapper) AutoFormPost(g *Global, id, after string) {
	_, code := DoPost(after+"/"+id, w.MapOfInputs())
	if code == 200 {
		g.Location.Set("href", after)
		return
	}
	g.flashThree("error")
}

func (g *Global) flashThree(s string) {
	flash := g.Document.ById("flash")
	flash.Set("innerHTML", s)
	time.Sleep(time.Second * 3)
	flash.Set("innerHTML", "")
}
