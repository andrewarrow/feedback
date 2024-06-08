package wasm

import "time"

func (g *Global) Toast(s string) {
	flash := g.Document.Id("toast")
	flash.Show()
	//flash.Set("innerHTML", s)
	time.Sleep(time.Second * 3)
	flash.Hide()
}
