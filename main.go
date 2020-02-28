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

func main() {
	rand.Seed(time.Now().UnixNano())

	if len(os.Args) > 1 {
		if strings.HasPrefix(os.Args[1], "--install=") {
			tokens := strings.Split(os.Args[1], "=")
			install := tokens[1]
			getDirsAndFiles(".")
			for _, path := range paths {
				all, _ := ioutil.ReadFile(path)
				tokens = strings.Split(path, "/")
				os.MkdirAll(install+strings.Join(tokens[:len(tokens)-1], "/"), 0755)
				//outfile := tokens[len(tokens)-1]
				ioutil.WriteFile(install+path, all, 0755)

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
