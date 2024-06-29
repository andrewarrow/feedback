
import sys
import os
from placeit import placeit

path = sys.argv[1]
name = sys.argv[2]

def gomain():
    gomainA()
    gomainB()
    
def gomainB():
    template = """\
package app

import (
  "github.com/andrewarrow/feedback/router"
)

func Welcome(c *router.Context, second, third string) {
	if second == "" && third == "" && c.Method == "GET" {
		handleWelcomeIndex(c)
		return
	}
	c.NotFound = true
}

func handleWelcomeIndex(c *router.Context) {

	send := map[string]any{}
	if len(c.User) == 0 {
		c.SendContentInLayout("welcome.html", send, 200)
		return
	}
}
    """
    placeit("app/welcome.go", {}, template)

def gomainA():
    template = """\
package main

import (
  "{{name}}/app"
  "embed"
  "math/rand"
  "os"
  "time"

  "github.com/andrewarrow/feedback/router"
)

//go:embed app/feedback.json
var embeddedFile []byte

//go:embed views/*.html
var embeddedTemplates embed.FS

//go:embed assets/**/*.*
var embeddedAssets embed.FS

var buildTag string

func main() {
	rand.Seed(time.Now().UnixNano())
	if len(os.Args) == 1 {
		//PrintHelp()
		return
	}

	arg := os.Args[1]

	if arg == "import" {
	} else if arg == "render" {
		router.RenderMarkup()
	} else if arg == "run" {
		router.BuildTag = buildTag
		router.EmbeddedTemplates = embeddedTemplates
		//router.EmbeddedAssets = embeddedAssets
		r := router.NewRouter("DATABASE_URL", embeddedFile)
		r.Paths["/"] = app.Welcome
		//r.Paths["{{name}}"] = app.{{capName}}
		//r.Paths["api"] = app.HandleApi
		//r.Paths["login"] = app.Login
		//r.Paths["register"] = app.Register
		//r.Paths["admin"] = app.Admin
		r.Paths["markup"] = router.Markup
		r.BucketPath = "/Users/aa/bucket"
		r.ListenAndServe(":" + os.Args[2])
	} else if arg == "help" {
	}
}
    """

    placeit("main.go", {"name": name, "capName": name}, template)

