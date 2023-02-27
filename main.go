package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/andrewarrow/feedback/router"
)

func PrintHelp() {
	fmt.Println("")
	fmt.Println("feedback v1.0")
	fmt.Println("")
	fmt.Println("  help              # this menu")
	fmt.Println("  new               # just like rails new")
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

	if arg == "new" {
	} else if arg == "run" {
		r := router.NewRouter("data/site.json")
		http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
			r.RouteFromRequest(writer, request)
		})

		log.Fatal(http.ListenAndServe(":3000", nil))
	} else if arg == "help" {
		PrintHelp()
	}

}
