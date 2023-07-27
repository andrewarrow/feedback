package router

import (
	"io/ioutil"
	"os"

	"github.com/andrewarrow/feedback/files"
)

func InitNewApp(path string) {

	place := path + "/tailwind"
	os.MkdirAll(place, 0755)
	place = path + "/app"
	os.MkdirAll(place, 0755)
	name := "feedback.json"
	asString := files.ReadFile(name)
	files.SaveFile(place+"/"+name, asString)
	place = path
	name = "run"
	asString = files.ReadFile(name)
	files.SaveFile(place+"/"+name, asString)
	name = "tailwind.config.js"
	asString = files.ReadFile(name)
	files.SaveFile(place+"/"+name, asString)
	name = "extra.html"
	place = path + "/tailwind"
	asString = files.ReadFile("tailwind/" + name)
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
