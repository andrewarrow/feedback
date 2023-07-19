package router

import (
	"io/ioutil"
	"os"

	"github.com/andrewarrow/feedback/files"
)

func InitNewApp(path string) {

	place := path + "/" + "app"
	os.MkdirAll(place, 0755)
	name := "welcome_controller.go"
	asString := files.ReadFile("router/" + name)
	files.SaveFile(place+"/"+name, asString)
	name = "feedback.json"
	asString = files.ReadFile(name)
	files.SaveFile(place+"/"+name, asString)
	place = path
	name = "main.go"
	asString = files.ReadFile(name)
	files.SaveFile(place+"/"+name, asString)

	dirs := []string{"views", "assets/css", "assets/images", "assets/javascript"}
	for _, dir := range dirs {
		place = path + "/" + dir
		os.MkdirAll(place, 0755)
		list, _ := ioutil.ReadDir(dir)
		for _, file := range list {
			name := file.Name()
			if dir == "views" {
				if name != "_table_large.html" && name != "_table_small.html" &&
					name != "_welcome_show_cols.html" &&
					name != "_nav_user.html" &&
					name != "application_layout.html" &&
					name != "sessions_new.html" &&
					name != "table_show.html" &&
					name != "404.html" &&
					name != "generic_top_bottom.html" {
					continue
				}
			}
			path := dir + "/" + name
			asString := files.ReadFile(path)
			files.SaveFile(place+"/"+file.Name(), asString)
		}
	}

}
