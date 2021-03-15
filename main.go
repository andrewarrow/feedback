package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/andrewarrow/feedback/command"
)

func PrintHelp() {
	fmt.Println("")
	fmt.Println("feedback v1.0")
	fmt.Println("")
	fmt.Println("  help              # this menu")
	fmt.Println("  new               # just like rails new")
	fmt.Println("")
}

func main() {
	rand.Seed(time.Now().UnixNano())
	if len(os.Args) == 1 {
		PrintHelp()
		return
	}
	arg := os.Args[1]

	if arg == "new" {
		command.NewMenu()
	} else if arg == "login" {
	} else if arg == "help" {
		PrintHelp()
	}

	/*
		if util.InitConfig() == false {
			print("no config")
			return
		}
		fmt.Println(util.AllConfig)
		if len(os.Args) > 2 {
			server.Serve(os.Args[1])
			util.AllConfig.Http.Host = os.Args[2]
		} else {
			server.Serve(util.AllConfig.Http.Port)
		}*/
}
