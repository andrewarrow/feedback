package wasm

func (w *Wrapper) MapOfInputs() map[string]any {
	m := map[string]any{}
	for _, input := range w.SelectAll("input") {
		m[input.Id] = input.Value
	}
	for _, input := range w.SelectAll("textarea") {
		m[input.Id] = input.Value
	}
	return m
}
