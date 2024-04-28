package wasm

import "strings"

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
		m[input.Id] = strings.TrimSpace(input.Value)
	}
	for _, input := range w.SelectAll("textarea") {
		input.Id = input.Id[len(prefix):]
		m[input.Id] = strings.TrimSpace(input.Value)
	}
	for _, input := range w.SelectAll("select") {
		input.Id = input.Id[len(prefix):]
		m[input.Id] = input.Value
	}
	return m
}
