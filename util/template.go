package util

import (
	"fmt"
	"io/ioutil"
	"os"
)

func MakeTemplate(name string) {
	asBytes, _ := ioutil.ReadFile("t")
	top := `{{ define "_%s" }}`
	end := `{{ end }}`
	s := fmt.Sprintf(top, name) + "\n" + string(asBytes) + "\n" + end
	ioutil.WriteFile("views/_"+name+".html", []byte(s), 0644)
	result := fmt.Sprintf(`{{ template "_%s" . }}`, name)
	fmt.Printf("\n\n%s\n\n", result)
	os.Remove("t")
}
