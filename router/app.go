package router

import (
	"fmt"
	"go/build"
	"io/ioutil"
	"os"

	"github.com/andrewarrow/feedback/files"
)

func InitNewApp() {
	importPath := "github.com/andrewarrow/feedback"
	pkg, _ := build.Import(importPath, "", build.FindOnly)
	fmt.Println(pkg.Dir)

	dirs := []string{"views", "assets/css", "assets/images", "assets/javascript"}
	for _, dir := range dirs {
		os.MkdirAll(dir, 0755)
		list, _ := ioutil.ReadDir(pkg.Dir + "/" + dir)
		for _, file := range list {
			path := pkg.Dir + "/" + dir + "/" + file.Name()
			asString := files.ReadFile(path)
			files.SaveFile(dir+"/"+file.Name(), asString)
		}
	}

}
