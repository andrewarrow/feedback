package wasm

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
	return m
}
