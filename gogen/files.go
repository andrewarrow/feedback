package gogen

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/andrewarrow/feedback/files"
	"github.com/andrewarrow/feedback/util"
)

func MakeControllerAndView(name, dir string) {
	c := `package app

import (
	"github.com/andrewarrow/feedback/router"
)

func Handle{{index . "camel" }}(c *router.Context, second, third string) {
	if c.User == nil {
		c.UserRequired = true
		return
	}
	if second == "" {
		handle{{index . "camel" }}Index(c)
	} else if second != "" && third == "" {
		handle{{index . "camel" }}Show(c, second)
	} else {
		c.NotFound = true
	}
}

func handle{{index . "camel" }}Index(c *router.Context) {
	if c.Method == "GET" {
		c.SendContentInLayout("{{index . "name" }}_index.html", nil, 200)
		return
	}
	handle{{index . "camel" }}Create(c)
}

func handle{{index . "camel" }}Create(c *router.Context) {
	c.NotFound = true
}

func handle{{index . "camel" }}Show(c *router.Context, id string) {
	if c.Method == "GET" {
		c.SendContentInLayout("{{index . "name" }}_show.html", nil, 200)
		return
	}
	handle{{index . "camel" }}Updates(c, id)
}

func handle{{index . "camel" }}Updates(c *router.Context, id string) {
	if c.Method == "POST" {
		c.NotFound = true
	} else if c.Method == "DELETE" {
		c.NotFound = true
	}
}`

	m := map[string]string{"name": name, "camel": util.ToCamelCase(name)}
	t, _ := template.New("c").Parse(c)
	content := new(bytes.Buffer)
	t.Execute(content, m)
	controller := content.String()
	files.SaveFile(dir+"/app/"+name+"_controller.go", controller)
	MakeView(name, dir+"/views/"+name+"_index.html")
	MakeView(name, dir+"/views/"+name+"_show.html")

	fmt.Printf("\nr.Paths[\"%s\"] = app.Handle%s\n", name, m["camel"])
}

func MakeView(name, path string) {
	v := `<article class="grid">
        <div>
          <hgroup>
            <h1>%s</h1>
            <h2>hi</h2>
          </hgroup>
        </div>
        <div></div>
</article>`
	view := fmt.Sprintf(v, name)
	files.SaveFile(path, view)
}
