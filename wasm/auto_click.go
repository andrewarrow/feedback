package wasm

import "fmt"

func (g *Global) AutoClick(prefix, suffix string, w *Wrapper, name string, cb func(string)) {
	for _, item := range w.SelectAllByClass(name) {
		id := item.Id[2:]
		click := func() {
			go func() {
				DoPost(fmt.Sprintf("/%s/%s/%s", prefix, id, suffix), nil)
				cb(id)
			}()
		}
		item.EventWithId(click)
	}
}
