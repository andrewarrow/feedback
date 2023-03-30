package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/andrewarrow/feedback/aigen"
	"github.com/andrewarrow/feedback/gogen"
	"github.com/andrewarrow/feedback/htmlgen"
	"github.com/andrewarrow/feedback/persist"
	"github.com/andrewarrow/feedback/prefix"
	"github.com/andrewarrow/feedback/router"
	"github.com/andrewarrow/feedback/util"
)

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
	} else if arg == "gen" {
		name := util.GetArg(2)
		path := util.GetArg(3)
		gogen.MakeControllerAndView(name, path)
	} else if arg == "ai" {
		// davinci  2049 tokens
		// gpt-3.5-turbo 4096 tokens lowest cost
		// code-davinci-002 8001 tokens
		prompt := util.GetArg(2)
		aigen.RunImage(prompt)
	} else if arg == "pp" {
		path := util.GetArg(2)
		htmlgen.PrettyPrint(path)
	} else if arg == "export" {
		db := persist.PostgresConnection("DATABASE_URL")
		jsonString := persist.SchemaJson(db)
		fmt.Println(jsonString)
	} else if arg == "scan" {
		list := persist.ScanSchema()
		asBytes := router.ModelsToBytes(list)
		persist.SaveSchema(asBytes)
	} else if arg == "run" {
		r := router.NewRouter("DATABASE_URL", nil)
		r.ListenAndServe(":3000")
	} else if arg == "help" {
		PrintHelp()
	}

}
