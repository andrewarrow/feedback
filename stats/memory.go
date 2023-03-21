package stats

import (
	"net/http"
	"sync"
	"time"

	"github.com/andrewarrow/feedback/util"
)

type Hit struct {
	Remote    string
	Agent     string
	Path      string
	Referer   string // keep the mis-spelling going Referrer
	Timestamp time.Time
}

var hitMutex sync.Mutex
var Hits = []*Hit{}

func AddHit(path string, request *http.Request) {
	h := Hit{}
	h.Agent = util.GetHeader("User-Agent", request)
	h.Path = path
	h.Remote = getRealIp(request)
	h.Timestamp = time.Now()
	h.Referer = util.GetHeader("Referer", request)
	if h.Referer == "" {
		h.Referer = util.GetHeader("HTTP_REFERER", request)
	}
	hitMutex.Lock()
	Hits = append([]*Hit{&h}, Hits...)
	if len(Hits) > 100 {
		Hits = Hits[0:99]
	}
	hitMutex.Unlock()
}

func getRealIp(request *http.Request) string {
	ip := util.GetHeader("X-Forwarded-For", request)
	if ip != "" {
		return ip
	}
	ip = util.GetHeader("X-Real-IP", request)
	if ip != "" {
		return ip
	}
	ip = util.GetHeader("Forwarded", request)
	if ip != "" {
		return ip
	}
	ip = util.GetHeader("Via", request)
	if ip != "" {
		return ip
	}
	return request.RemoteAddr
}
