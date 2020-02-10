package controllers

import "github.com/gin-gonic/gin"
import "github.com/andrewarrow/feedback/util"
import "net/http"
import "io/ioutil"
import "os"

func DirectoriesIndex(c *gin.Context) {
	files, dirs := getDirsAndFiles("")
	c.HTML(http.StatusOK, "list.tmpl", gin.H{
		"files": files,
		"dirs":  dirs,
	})

}
func DirectoriesDownload(c *gin.Context) {
	active := util.AllConfig.Directories.Active
	fileName := c.Param("name")
	targetPath := active + "/" + fileName
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Type", "application/octet-stream")
	c.File(targetPath)
}
func DirectoriesNameIndex(c *gin.Context) {
	fileName := c.Param("name")
	files, dirs := getDirsAndFiles("/" + fileName)
	c.HTML(http.StatusOK, "list.tmpl", gin.H{
		"files": files,
		"dirs":  dirs,
	})
}
func DirectoriesDownloadExtra(c *gin.Context) {
	name := c.Param("name")
	extra := c.Param("extra")
	active := util.AllConfig.Directories.Active
	targetPath := active + "/" + name + "/" + extra
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename="+extra)
	c.Header("Content-Type", "application/octet-stream")
	c.File(targetPath)
}

func getDirsAndFiles(extra string) ([]string, []string) {
	active := util.AllConfig.Directories.Active
	files, _ := ioutil.ReadDir(active + extra)
	dirs := []string{}
	buff := []string{}
	for _, f := range files {
		fi, _ := os.Stat(active + extra + "/" + f.Name())
		if fi.IsDir() {
			dirs = append(dirs, f.Name())
		} else {
			buff = append(buff, f.Name())
		}
	}
	return buff, dirs
}
