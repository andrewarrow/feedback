package command

import (
	"fmt"
	"os"
)

func NewHelp() {
	fmt.Println("")
	fmt.Printf("%30s     %s\n", "new <name_of_app>", "create new app")
	fmt.Println("")
}

func NewMenu() {
	if len(os.Args) == 2 {
		NewHelp()
		return
	}
	if os.Args[2] != "" {
		return
	}
}
