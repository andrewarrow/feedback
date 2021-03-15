package command

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var paths []string = []string{}

func NewHelp() {
	fmt.Println("")
	fmt.Printf("%30s     %s\n", "new <name_of_app>", "create new app")
	fmt.Println("")
}

func NewMenu() {
	if len(os.Args) == 2 {
		NewHelp()
		return
	}
	if os.Args[2] != "" {
		NewProcess()
		return
	}
}

func NewProcess() {
	cliInstall(os.Args[2])
}

func cliInstall(name string) {
	dir, _ := os.Getwd()
	getDirsAndFiles(dir)
	for _, path := range paths {
		all, _ := ioutil.ReadFile(path)
		tokens := strings.Split(path, "/")
		index := 0
		for i, token := range tokens {
			if token == "feedback" {
				index = i
				break
			}
		}
		dest := name + "/" + strings.Join(tokens[index+1:len(tokens)-1], "/")
		fmt.Println(dest)
		os.MkdirAll(dest, 0755)
		if tokens[len(tokens)-1] == "conf.toml.dist" {
			tokens[len(tokens)-1] = "conf.toml"
		}
		ioutil.WriteFile(dest+"/"+tokens[len(tokens)-1],
			replacePackageNames(string(all), name, path), 0666)
	}
}

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
	if strings.HasSuffix(path, ".go") {
		fixed := strings.ReplaceAll(all, "github.com/andrewarrow/feedback/", name+"/")
		return []byte(fixed)
	}
	if strings.HasSuffix(path, ".mod") {
		fixed := strings.ReplaceAll(all, "github.com/andrewarrow/feedback", name)
		return []byte(fixed)
	}
	return []byte(all)

}
