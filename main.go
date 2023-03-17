package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/andrewarrow/feedback/gogen"
	"github.com/andrewarrow/feedback/htmlgen"
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

	if arg == "reset" {
		r := router.NewRouter()
		r.ResetDatabase()
	} else if arg == "gen" {
		name := util.GetArg(2)
		path := util.GetArg(3)
		gogen.MakeControllerAndView(name, path)
	} else if arg == "pp" {
		htmlgen.PrettyPrint()
	} else if arg == "run" {
		r := router.NewRouter()
		r.ListenAndServe(":3000")
	} else if arg == "help" {
		PrintHelp()
	}

}
