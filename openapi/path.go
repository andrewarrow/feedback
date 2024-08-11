package openapi

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"

	"github.com/andrewarrow/feedback/router"
)

func (oa *OpenAPI) AddPath(path string, fn func(*router.Context, string, string)) {
	v := reflect.ValueOf(fn)
	name := runtime.FuncForPC(v.Pointer()).Name()
	tokens := strings.Split(name, ".")
	name = tokens[len(tokens)-1]

	oa.FuncToPath[name] = path
	fmt.Println(name, path)
}
