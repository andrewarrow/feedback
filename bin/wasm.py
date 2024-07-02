
import sys
import os
from placeit import placeit

path = sys.argv[1]
name = sys.argv[2]

def wasm():
    template = """\
package main

import (
	"embed"
	"fmt"
	"{{name}}/browser"
	"math/rand"
	"time"

	"github.com/andrewarrow/feedback/wasm"
)

//go:embed views/*.html
var embeddedTemplates embed.FS

var useLive string
var viewList string

func main() {
	fmt.Println(viewList)
	wasm.UseLive = useLive == "true"
	wasm.EmbeddedTemplates = embeddedTemplates
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Go Web Assembly")
	browser.Global, browser.Document = wasm.NewGlobal()

	<-browser.Global.Ready
	if wasm.UseLive {
		files, _ := embeddedTemplates.ReadDir("views")
		go func() {
			wasm.LoadAllTemplates(files)
			browser.RegisterEvents()
		}()
	} else {
		browser.RegisterEvents()
	}

	select {}
}
    """

    placeit("wasm/main.go", {"name": name}, template)

