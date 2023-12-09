package wasm

func (d *Document) Render(id, template string, payload map[string]any) {
	div := d.ById(id)
	div.Set("innerHTML", "hi")
}
