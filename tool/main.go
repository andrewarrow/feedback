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

	"github.com/andrewarrow/feedback/util"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	if len(os.Args) == 1 {
		return
	}
	arg := os.Args[1]

	if arg == "init" {
		initApp(os.Args[2], os.Args[3])
	} else if arg == "controller" {
		controller(os.Args[2], os.Args[3])
	} else if arg == "table" {
		table(os.Args[2], os.Args[3])
	} else if arg == "editable" {
		editable(os.Args[2], os.Args[3])
	} else if arg == "list" {
		list(os.Args[2], os.Args[3])
	} else if arg == "" {
	}
}

func initApp(path, name string) {
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

	tmpl, _ := template.New("").Parse(mainTemplate())
	m["package"] = lower
	result := bytes.NewBuffer([]byte{})
	tmpl.Execute(result, m)
	filename := "main.go"
	ioutil.WriteFile(path+"/"+filename, result.Bytes(), 0644)

	tmpl, _ = template.New("").Parse(modTemplate())
	result = bytes.NewBuffer([]byte{})
	tmpl.Execute(result, m)
	filename = "go.mod"
	ioutil.WriteFile(path+"/"+filename, result.Bytes(), 0644)

	tmpl, _ = template.New("").Parse(runTemplate())
	result = bytes.NewBuffer([]byte{})
	tmpl.Execute(result, m)
	filename = "run"
	ioutil.WriteFile(path+"/"+filename, result.Bytes(), 0644)

	tmpl, _ = template.New("").Parse(ignoreTemplate())
	result = bytes.NewBuffer([]byte{})
	tmpl.Execute(result, m)
	filename = ".gitignore"
	ioutil.WriteFile(path+"/"+filename, result.Bytes(), 0644)
}
