package wasm

import "syscall/js"

func FuncOf(fn func()) {
	thefunc := func(this js.Value, p []js.Value) any {
		fn()
		return nil
	}
}
