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
		if f.Name() == ".git" || f.Name() == ".gitattributes" || f.Name() == ".gitignore" {
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
		if strings.HasPrefix(os.Args[1], "--help") {
			fmt.Println("--install=dir")
			fmt.Println("--form=thing")
			fmt.Println("--model=thing")
		} else if strings.HasPrefix(os.Args[1], "--install=") {
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
		} else if strings.HasPrefix(os.Args[1], "--model=") {
			tokens := strings.Split(os.Args[1], "=")
			thing := tokens[1]
			models := `
{{define "T"}}
package models

import "github.com/jmoiron/sqlx"
import "fmt"

type {{.upperThing}} struct {
	Id        int   {{.jsonId}} 
	CreatedAt int64 {{.jsonCreatedAt}}
}

const {{.allUpper}}_SELECT = "SELECT id, UNIX_TIMESTAMP(created_at) as createdat from {{.thing}}s"

func Select{{.upperThing}}s(db *sqlx.DB) ([]{{.upperThing}}, string) {
	{{.thing}}s := []{{.upperThing}}{}
	sql := fmt.Sprintf("%s order by created_at desc", {{.allUpper}}_SELECT)
	err := db.Select(&{{.thing}}s, sql)
	s := ""
	if err != nil {
		s = err.Error()
	}

	return {{.thing}}s, s
}
func Insert{{.upperThing}}(db *sqlx.DB) string {
	_, err := db.NamedExec("INSERT INTO {{.thing}}s (col) values (:col)",
		map[string]interface{}{"": ""})
	if err != nil {
		return err.Error()
	}
	return ""
}
{{end}}
`
			var buf bytes.Buffer
			path, _ := os.Getwd()

			capital := strings.ToUpper(thing[0:1])
			m := map[string]interface{}{"thing": thing,
				"jsonId":        "`json:\"id\"`",
				"jsonCreatedAt": "`json:\"created_at\"`",
				"allUpper":      strings.ToUpper(thing),
				"upperThing":    capital + thing[1:]}

			t, _ := template.New("").Parse(models)
			t.ExecuteTemplate(&buf, "T", m)
			ioutil.WriteFile(fmt.Sprintf("%s/models/%s.go", path, thing),
				buf.Bytes(), 0666)
		} else if strings.HasPrefix(os.Args[1], "--form=") {
			tokens := strings.Split(os.Args[1], "=")
			thing := tokens[1]

			// - tmpl
			// - controller
			// - routes

			templates := `{{define "T"}}
{{printf "%s" "{{template \"_header\" .}}"}}
<form method="post" action="/{{.thing}}">
<input type="text" name="thing"/>
<br/>
<input type="submit"/>
</form>
{{printf "%s" "{{template \"_footer\" .}}"}}
{{end}}`

			routes := `{{define "T"}}
{{.thing}} := router.Group("/{{.thing}}")
{{.thing}}.GET("/new", controllers.{{.upperThing}}New)
{{.thing}}.POST("/", controllers.{{.upperThing}}Create)
{{.thing}}.POST("/destroy", controllers.{{.upperThing}}Destroy)
{{end}}
`

			controllers := `
{{define "T"}}
package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)
func {{.upperThing}}New(c *gin.Context) {
	BeforeAll("", c)
	c.HTML(http.StatusOK, "{{.thing}}__new.tmpl", gin.H{
		"flash": "",
	})

}
func {{.upperThing}}Create(c *gin.Context) {
	BeforeAll("", c)
	c.Redirect(http.StatusFound, "/")
	c.Abort()
}
func {{.upperThing}}Destroy(c *gin.Context) {

	c.Redirect(http.StatusFound, "/")
	c.Abort()
}
{{end}}
`
			var buf bytes.Buffer
			path, _ := os.Getwd()

			t, _ := template.New("").Parse(routes)
			capital := strings.ToUpper(thing[0:1])
			m := map[string]interface{}{"thing": thing, "upperThing": capital + thing[1:]}
			t.ExecuteTemplate(&buf, "T", m)
			fmt.Println(buf.String())

			buf.Reset()

			t, _ = template.New("").Parse(controllers)
			t.ExecuteTemplate(&buf, "T", m)
			ioutil.WriteFile(fmt.Sprintf("%s/controllers/%s.go", path, thing),
				buf.Bytes(), 0666)

			buf.Reset()

			t, _ = template.New("").Parse(templates)
			t.ExecuteTemplate(&buf, "T", m)
			ioutil.WriteFile(fmt.Sprintf("%s/templates/%s__new.tmpl", path, thing),
				buf.Bytes(), 0666)
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
