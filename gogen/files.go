package gogen

import (
	"fmt"

	"github.com/andrewarrow/feedback/files"
	"github.com/andrewarrow/feedback/util"
)

func MakeControllerAndView(name, dir string) {
	c := `package app

import (
  "github.com/andrewarrow/feedback/router"
)

func Handle%s(c *router.Context, second, third string) {
  if c.User == nil {
    c.UserRequired = true
    return
  }
  c.SendContentInLayout("%s_index.html", nil, 200)
}`

	camel := util.ToCamelCase(name)
	controller := fmt.Sprintf(c, camel, name)
	files.SaveFile(dir+"/app/"+name+"_controller.go", controller)
	MakeView(name, dir+"/views/"+name+"_index.html")

	fmt.Printf("\nr.Paths[\"%s\"] = app.Handle%s\n", name, camel)
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
