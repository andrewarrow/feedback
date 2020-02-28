package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/andrewarrow/feedback/server"
	"github.com/andrewarrow/feedback/util"
)

var paths []string = []string{}

func getDirsAndFiles(dir string) {
	files, _ := ioutil.ReadDir(dir)
	for _, f := range files {
		if f.Name() == ".git" || f.Name() == ".gitattributes" {
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
		if strings.HasPrefix(os.Args[1], "--install=") {
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
		}
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
