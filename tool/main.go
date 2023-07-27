package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
)

func main() {
	path := os.Args[1]
	fmt.Println(path)
	os.Mkdir(path+"/app", 0775)
	os.Mkdir(path+"/views", 0775)
	os.Mkdir(path+"/assets", 0775)

	m := map[string]string{"name": "Foo",
		"lower":  "foo",
		"with_s": "foos",
	}
	tmpl, _ := template.New("t1").Parse(controllerTemplate())
	var result bytes.Buffer
	tmpl.Execute(&result, m)

	filename := "foo_controller.go"
	ioutil.WriteFile(path+"/app/"+filename, result.Bytes(), 0644)
}
