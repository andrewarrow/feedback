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
	for _, input := range w.SelectAll("select") {
		m[input.Id] = input.Value
		input.Set("value", "")
	}
	for _, input := range w.SelectAll("hidden") {
		m[input.Id] = input.Value
	}
	return m
}

func (w *Wrapper) NoClearInputs() map[string]any {
	m := map[string]any{}
	for _, input := range w.SelectAll("input") {
		if input.Get("type") == "submit" {
			continue
		}
		m[input.Id] = input.Value
	}
	for _, input := range w.SelectAll("textarea") {
		m[input.Id] = input.Value
	}
	for _, input := range w.SelectAll("select") {
		m[input.Id] = input.Value
	}
	for _, input := range w.SelectAll("hidden") {
		m[input.Id] = input.Value
	}
	return m
}
