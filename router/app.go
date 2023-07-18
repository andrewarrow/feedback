package router

import (
	"io/ioutil"
	"os"

	"github.com/andrewarrow/feedback/files"
)

func InitNewApp(path string) {

	dirs := []string{"views", "assets/css", "assets/images", "assets/javascript"}
	for _, dir := range dirs {
		place := path + "/" + dir
		os.MkdirAll(place, 0755)
		list, _ := ioutil.ReadDir(dir)
		for _, file := range list {
			path := dir + "/" + file.Name()
			asString := files.ReadFile(path)
			files.SaveFile(place+"/"+file.Name(), asString)
		}
	}

}
