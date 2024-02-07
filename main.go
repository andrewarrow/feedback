package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/andrewarrow/feedback/aigen"
	"github.com/andrewarrow/feedback/files"
	"github.com/andrewarrow/feedback/gogen"
	"github.com/andrewarrow/feedback/persist"
	"github.com/andrewarrow/feedback/prefix"
	"github.com/andrewarrow/feedback/router"
	"github.com/andrewarrow/feedback/util"
)

//go:embed feedback.json
var embeddedFile []byte

//go:embed views/*.html
var embeddedTemplates embed.FS

//go:embed assets/**/*.*
var embeddedAssets embed.FS

func PrintHelp() {
	fmt.Println("")
	fmt.Println("feedback v1.0")
	fmt.Println("")
	fmt.Println("  help              # this menu")
	fmt.Println("  reset             # reset database")
	fmt.Println("  run               # ")
	fmt.Println("")
}

func main() {
	rand.Seed(time.Now().UnixNano())
	if len(os.Args) == 1 {
		PrintHelp()
		return
	}
	arg := os.Args[1]

	prefix.FeedbackName = os.Getenv("FEEDBACK_NAME")

	if arg == "reset" {
		//r := router.NewRouter("DATABASE_URL")
		//r.ResetDatabase()
	} else if arg == "genmd" {
		model := util.GetArg(2)
		path := util.GetArg(3)
		jsonBytes := getFeedbackJsonFile(path)
		r := router.NewRouter("NO_DB", jsonBytes)
		gogen.MakeMarkDown(r.Site, model)
	} else if arg == "gencode" {
		model := util.GetArg(2)
		path := util.GetArg(3)
		jsonBytes := getFeedbackJsonFile(path)
		r := router.NewRouter("NO_DB", jsonBytes)
		gogen.MakeRoutes(r.Site.Routes, model)
	} else if arg == "ai" {
		// davinci  2049 tokens
		// gpt-3.5-turbo 4096 tokens lowest cost
		// code-davinci-002 8001 tokens
		prompt := util.GetArg(2)
		aigen.RunImage(prompt)
	} else if arg == "hash" {
		text := util.PseudoUuid()
		text = "testing123"
		hash := router.HashPassword(text)
		fmt.Println(text)
		fmt.Printf("update lyfe_users set username='',password='%s' where firebase_uid='';\n\n", hash)
	} else if arg == "scan" {
		//r := router.NewRouter("DATABASE_URL", embeddedFile)
		db := os.Getenv("DATABASE_URL")
		var site router.FeedbackSite
		site.Models = persist.ScanSchema(db)
		asBytes, _ := json.Marshal(site)
		jqed := util.PipeToJq(string(asBytes))
		files.SaveFile("feedback.json", jqed)
		//tablesString := util.GetArg(2)
		//list := persist.ModelsForTables(r.Db, tablesString)
		//asBytes := router.ModelsToBytes(list)
		//persist.SaveSchema(asBytes)
	} else if arg == "guids" {
		path := util.GetArg(2)
		jsonBytes := getFeedbackJsonFile(path)
		r := router.NewRouter("DATABASE_URL", jsonBytes)
		router.MakeGuidsInTables(r.Db, r.Site.Models)
	} else if arg == "init" {
		path := util.GetArg(2)
		router.InitNewApp(path)
	} else if arg == "run" {
		path := util.GetArg(2)
		jsonBytes := getFeedbackJsonFile(path)
		router.EmbeddedTemplates = embeddedTemplates
		router.EmbeddedAssets = embeddedAssets
		r := router.NewRouter("DATABASE_URL", jsonBytes)
		r.BucketPath = "/Users/aa/bucket"
		r.ListenAndServe(":3000")
	} else if arg == "help" {
		PrintHelp()
	}

}

func getFeedbackJsonFile(path string) []byte {
	var jsonBytes []byte
	if path != "" {
		asString := files.ReadFile(path)
		jsonBytes = []byte(asString)
	} else {
		jsonBytes = embeddedFile
	}
	return jsonBytes
}
