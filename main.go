package main

import "github.com/andrewarrow/feedback/server"
import "time"
import "math/rand"
import "github.com/andrewarrow/feedback/util"
import "fmt"
import "os"

func main() {
	rand.Seed(time.Now().UnixNano())
	if util.InitConfig() == false {
		print("no config")
		return
	}
	fmt.Println(util.AllConfig)
	if len(os.Args) > 1 {
		server.Serve(os.Args[1])
	} else {
		server.Serve(util.AllConfig.Http.Port)
	}
}
