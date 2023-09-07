package router

import (
	"io"
	"io/ioutil"
	"strings"

	"github.com/andrewarrow/feedback/util"
)

func (c *Context) ReadFormValuesIntoParams(list ...string) {
	c.Params = map[string]any{}
	for _, name := range list {
		val := strings.TrimSpace(c.Request.FormValue(name))
		c.Params[name] = val
	}
}

func (c *Context) ReadMultipleFormValues(list ...string) {
	c.Params = map[string]any{}
	for _, name := range list {
		selectedValues := c.Request.PostForm[name]
		buffer := []string{}
		for _, item := range selectedValues {
			buffer = append(buffer, strings.TrimSpace(item))
		}

		val := strings.Join(buffer, ",")
		c.Params[name] = val
	}
}

func SaveMultiFiles(c *Context) {
	files := c.Request.MultipartForm.File["file"]

	for _, fileHeader := range files {
		name := fileHeader.Filename
		file, _ := fileHeader.Open()
		asBytes, _ := io.ReadAll(file)
		file.Close()
		filename := util.GuidFilename(name)
		ioutil.WriteFile(c.Router.BucketPath+"/"+filename, asBytes, 0644)
		c.Params["photo"] = filename
	}
}
