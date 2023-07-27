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
	tmpl, _ := template.New("").Parse(controllerTemplate())
	result := bytes.NewBuffer([]byte{})
	tmpl.Execute(result, m)

	filename := "foo_controller.go"
	ioutil.WriteFile(path+"/app/"+filename, result.Bytes(), 0644)

	tmpl, _ = template.New("").Parse(mainTemplate())
	m["package"] = "foo"
	result = bytes.NewBuffer([]byte{})
	tmpl.Execute(result, m)
	filename = "main.go"
	ioutil.WriteFile(path+"/"+filename, result.Bytes(), 0644)
}
