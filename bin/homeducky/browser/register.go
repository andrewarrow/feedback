package browser

import (
	"github.com/andrewarrow/feedback/wasm"
)

var Global *wasm.Global
var Document *wasm.Document

func RegisterEvents() {
	LogoutEvents()
	afterRegister := func(id int64) {
		Global.Location.Set("href", "/core/start")
	}
	afterLogin := func(id int64) {
		Global.Location.Set("href", "/core/start")
	}
	if Global.Start == "start.html" {
	} else if Global.Start == "login.html" {
		Global.AutoForm("login", "core", nil, afterLogin)
	} else if Global.Start == "register.html" {
		Global.AutoForm("register", "core", nil, afterRegister)
	}
}

func LogoutEvents() {
	if Document.Id("logout") == nil {
		return
	}
	Global.Event("logout", Global.Logout("/core"))
}
