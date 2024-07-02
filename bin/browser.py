
import sys
import os
from placeit import placeit

path = sys.argv[1]
name = sys.argv[2]

def browser():
    template = """\
package browser

import (
	"github.com/andrewarrow/feedback/wasm"
)

var Global *wasm.Global
var Document *wasm.Document

func RegisterEvents() {
	afterRegister := func(id int64) {
		Global.Location.Set("href", "/{{name}}/start")
	}
	afterLogin := func(id int64) {
		Global.Location.Set("href", "/{{name}}/start")
	}
	if Global.Start == "start.html" {
	} else if Global.Start == "login.html" {
		Global.AutoForm("login", "{{name}}", nil, afterLogin)
	} else if Global.Start == "register.html" {
		Global.AutoForm("register", "{{name}}", nil, afterRegister)
	}
}
    """

    placeit("browser/register.go", {"name": name}, template)

