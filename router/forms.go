package router

import (
	"io"
	"io/ioutil"
	"strings"

	"github.com/andrewarrow/feedback/buckets"
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

type UploadedFile struct {
	OrigName     string
	GuidFilename string
	Size         int64
}

func SaveMultiFiles(c *Context, guid string) []UploadedFile {
	list := []UploadedFile{}
	files := c.Request.MultipartForm.File["file"]

	for _, fileHeader := range files {
		name := fileHeader.Filename
		file, _ := fileHeader.Open()
		asBytes, _ := io.ReadAll(file)
		file.Close()
		filename := util.GuidFilename(name, guid)
		ioutil.WriteFile(c.Router.BucketPath+"/"+filename, asBytes, 0644)
		c.Params["photo"] = filename
		up := UploadedFile{}
		up.OrigName = name
		up.GuidFilename = filename
		up.Size = int64(len(asBytes))
		list = append(list, up)
	}
	return list
}

func SaveMultiFilesAws(c *Context, guid string) {
	list := []string{"", "_2", "_3", "_4", "_5"}
	for _, item := range list {
		files := c.Request.MultipartForm.File["file"+item]

		for _, fileHeader := range files {
			name := fileHeader.Filename
			file, _ := fileHeader.Open()
			asBytes, _ := io.ReadAll(file)
			file.Close()
			filename := util.GuidFilename(name, guid)
			buckets.StoreInAws(asBytes, filename)
			c.Params["photo"+item] = filename
		}
	}
}
