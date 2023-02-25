package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/andrewarrow/feedback/files"
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
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			welcome := files.ReadFile("views/welcome.html")
			fmt.Fprintf(w, welcome)
		})

		log.Fatal(http.ListenAndServe(":3000", nil))
	} else if arg == "help" {
		PrintHelp()
	}

}
