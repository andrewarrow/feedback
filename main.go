package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/andrewarrow/feedback/server"
	"github.com/andrewarrow/feedback/util"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	gpPlusFeedback := os.Getenv("GOPATH") + "/src/github.com/andrewarrow/feedback/"
	fi, err := os.Stat(gpPlusFeedback + "main.go")
	if err != nil {
		fmt.Println("go get github.com/andrewarrow/feedback")
		return
	}
	if fi.Size() == 0 {
		fmt.Println("go get github.com/andrewarrow/feedback")
		return
	}

	if len(os.Args) > 1 && handledByCli(gpPlusFeedback) {
		return
	}

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
	}
}
