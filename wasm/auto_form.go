package wasm

type AutoForm struct {
	ReturnPath string
	Path       string
	After      func(int64)
}

func NewAutoForm(id string) *AutoForm {
	a := AutoForm{}
	return &a
}
