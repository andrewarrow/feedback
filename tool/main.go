package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/andrewarrow/feedback/util"
)

func main() {
	path := os.Args[1]
	name := os.Args[2]
	fmt.Println(path)
	os.Mkdir(path+"/app", 0775)
	os.Mkdir(path+"/views", 0775)
	os.Mkdir(path+"/assets", 0775)

	lower := strings.ToLower(name)
	withS := util.Plural(lower)

	m := map[string]string{"name": name,
		"lower":  lower,
		"with_s": withS,
	}
	tmpl, _ := template.New("").Parse(controllerTemplate())
	result := bytes.NewBuffer([]byte{})
	tmpl.Execute(result, m)

	filename := lower + "_controller.go"
	ioutil.WriteFile(path+"/app/"+filename, result.Bytes(), 0644)

	tmpl, _ = template.New("").Parse(mainTemplate())
	m["package"] = lower
	result = bytes.NewBuffer([]byte{})
	tmpl.Execute(result, m)
	filename = "main.go"
	ioutil.WriteFile(path+"/"+filename, result.Bytes(), 0644)

	tmpl, _ = template.New("").Parse(modTemplate())
	result = bytes.NewBuffer([]byte{})
	tmpl.Execute(result, m)
	filename = "go.mod"
	ioutil.WriteFile(path+"/"+filename, result.Bytes(), 0644)
}
