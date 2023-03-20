package stats

import (
	"net/http"
	"strings"
	"sync"
	"time"
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
	h.Agent = getHeader("User-Agent", request)
	h.Path = path
	h.Remote = getRealIp(request)
	h.Timestamp = time.Now()
	h.Referer = getHeader("Referer", request)
	if h.Referer == "" {
		h.Referer = getHeader("HTTP_REFERER", request)
	}
	hitMutex.Lock()
	Hits = append([]*Hit{&h}, Hits...)
	if len(Hits) > 100 {
		Hits = Hits[0:99]
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

func getRealIp(request *http.Request) string {
	ip := getHeader("X-Forwarded-For", request)
	if ip != "" {
		return ip
	}
	ip = getHeader("X-Real-IP", request)
	if ip != "" {
		return ip
	}
	ip = getHeader("Forwarded", request)
	if ip != "" {
		return ip
	}
	ip = getHeader("Via", request)
	if ip != "" {
		return ip
	}
	return request.RemoteAddr
}
