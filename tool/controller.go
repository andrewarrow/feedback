package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
	"text/template"

	"github.com/andrewarrow/feedback/util"
)

func controller(path, name string) {
	fmt.Println(path)

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

	tmpl, _ = template.New("").Parse(createTemplate())
	result = bytes.NewBuffer([]byte{})
	tmpl.Execute(result, m)

	filename = lower + "_create.go"
	ioutil.WriteFile(path+"/app/"+filename, result.Bytes(), 0644)

	tmpl, _ = template.New("").Parse(showTemplate())
	result = bytes.NewBuffer([]byte{})
	tmpl.Execute(result, m)

	filename = lower + "_show.go"
	ioutil.WriteFile(path+"/app/"+filename, result.Bytes(), 0644)
}
