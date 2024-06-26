package wasm

import "syscall/js"

func FuncOf(fn func(js.Value)) any {
	theFunc := func(this js.Value, p []js.Value) any {
		fn(p[0])
		return nil
	}
	return js.FuncOf(theFunc)
}
func SimpleFuncOf(fn func()) any {
	theFunc := func(this js.Value, p []js.Value) any {
		fn()
		return nil
	}
	return js.FuncOf(theFunc)
}
