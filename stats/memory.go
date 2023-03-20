package stats

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
)

type Hit struct {
	Remote    string
	Agent     string
	Path      string
	Timestamp time.Time
}

var hitMutex sync.Mutex
var hits = []*Hit{}

func AddHit(path string, request *http.Request) {
	h := Hit{}
	h.Agent = getHeader("User-Agent", request)
	h.Path = path
	h.Remote = request.RemoteAddr // TODO X-Forwarded-For, etc.
	h.Timestamp = time.Now()
	hitMutex.Lock()
	hits = append([]*Hit{&h}, hits...)
	if len(hits) > 100 {
		hits = hits[0:99]
	}
	hitMutex.Unlock()
}

func getHeader(field string, request *http.Request) string {
	val := request.Header[field]
	if len(val) == 0 {
		return ""
	}
	return strings.Join(val, ",")
}

func main() {
	fmt.Println("vim-go")
}
