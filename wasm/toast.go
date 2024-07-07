package wasm

import "time"

func (g *Global) Toast(s string) {
	flash := g.Document.Id("toast")
	tn := g.Document.Id("toast-name")
	tn.Set("innerHTML", s)
	flash.Show()
	//flash.Set("innerHTML", s)
	time.Sleep(time.Second * 3)
	flash.Hide()
}
func (g *Global) ToastFlash(s string) {
	flash := g.Document.ById("flash")
	flash.Set("innerHTML", s)
	time.Sleep(time.Second * 3)
	flash.Set("innerHTML", "")
}
