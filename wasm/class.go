package wasm

import (
	"strings"
	"syscall/js"
)

func AddClass(w js.Value, className string) {
	currentClass := w.Get("className").String()
	newClass := currentClass + " " + className
	w.Set("className", newClass)
}

func RemoveClass(w js.Value, className string) {
	currentClass := w.Get("className").String()
	tokens := strings.Split(currentClass, " ")
	buffer := []string{}
	for _, item := range tokens {
		if item == className {
			continue
		}
		buffer = append(buffer, item)
	}
	w.Set("className", strings.Join(buffer, " "))
}

func HasClass(w js.Value, className string) bool {
	currentClass := w.Get("className").String()
	tokens := strings.Split(currentClass, " ")
	for _, item := range tokens {
		if item == className {
			return true
		}
	}
	return false
}
