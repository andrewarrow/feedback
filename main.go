package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
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
	command := os.Args[1]

	if command == "new" {
	} else if command == "login" {
	} else if command == "help" {
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
