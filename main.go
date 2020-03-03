package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/andrewarrow/feedback/server"
	"github.com/andrewarrow/feedback/util"
)

var paths []string = []string{}

func getDirsAndFiles(dir string) {
	files, _ := ioutil.ReadDir(dir)
	for _, f := range files {
		if f.Name() == ".git" || f.Name() == ".gitattributes" {
			continue
		}
		fi, _ := os.Lstat(dir + "/" + f.Name())
		if fi.IsDir() {
			getDirsAndFiles(dir + "/" + f.Name())
		} else {
			path := dir + "/" + f.Name()
			tokens := strings.Split(path, "/")
			if len(tokens) > 2 {
				fmt.Println(path)
				paths = append(paths, path)
			}
		}
	}
}

func replacePackageNames(all, name, path string) []byte {
	if !strings.HasSuffix(path, ".go") {
		return []byte(all)
	}

	fixed := strings.ReplaceAll(all, "github.com/andrewarrow/feedback/", name+"/")
	return []byte(fixed)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	if len(os.Args) > 1 {
		if strings.HasPrefix(os.Args[1], "--install=") {
			getDirsAndFiles(".")
			tokens := strings.Split(os.Args[1], "=")
			install := tokens[1]
			tokens = strings.Split(install, "/")
			name := tokens[len(tokens)-2]
			for _, path := range paths {
				all, _ := ioutil.ReadFile(path)
				tokens = strings.Split(path, "/")
				os.MkdirAll(install+strings.Join(tokens[:len(tokens)-1], "/"), 0755)
				fmt.Println(name)
				ioutil.WriteFile(install+path, replacePackageNames(string(all), name, path), 0666)
			}
		} else if strings.HasPrefix(os.Args[1], "--form=") {
			tokens := strings.Split(os.Args[1], "=")
			thing := tokens[1]

			// - tmpl
			// - controller
			// - routes

			routes := `{{define "T"}}
{{.thing}} := router.Group("/{{.thing}}")
sessions.GET("/new", controllers.{{.upperThing}}New)
sessions.POST("/", controllers.{{.upperThing}}Create)
sessions.POST("/destroy", controllers.{{.upperThing}}Destroy)
{{end}}
`
			var buf bytes.Buffer
			t, err := template.New("").Parse(routes)
			capital := strings.ToUpper(thing[0:1])
			m := map[string]interface{}{"thing": thing, "upperThing": capital + thing[1:]}
			err = t.ExecuteTemplate(&buf, "T", m)
			fmt.Println(err, buf.String())
		}
		return
	}

	if util.InitConfig() == false {
		print("no config")
		return
	}
	fmt.Println(util.AllConfig)
	if len(os.Args) > 2 {
		server.Serve(os.Args[1])
		util.AllConfig.Http.Host = os.Args[2]
	} else {
		server.Serve(util.AllConfig.Http.Port)
	}
}
