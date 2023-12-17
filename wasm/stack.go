package wasm

type StackItem struct {
	HTML     string
	Callback func()
}

func NewStackItem(s string) *StackItem {
	si := StackItem{}
	si.HTML = s
	return &si
}
