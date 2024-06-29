package wasm

import (
	"syscall/js"
)

func GetItemMap(item js.Value, count int) map[string]any {
	if count > 10 {
		return nil
	}
	m := map[string]any{}
	o := js.Global().Get("Object")
	keys := o.Call("keys", item)

	for i := 0; i < keys.Length(); i++ {
		key := keys.Index(i).String()
		value := item.Get(key)
		if value.Type() == js.TypeNumber {
			m[key] = value.Float()
		} else if value.Type() == js.TypeString {
			m[key] = value.String()
		} else if value.Type() == js.TypeBoolean {
			m[key] = value.Bool()
		} else if value.Type() == js.TypeObject {
			//m[key] = GetItemMap(value, count+1)
		} else if value.IsNull() || value.IsUndefined() {
			m[key] = nil
		} else {
			m[key] = value.String()
		}
	}
	return m
}
