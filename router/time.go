package router

import (
	"fmt"
	"time"

	"github.com/xeonx/timeago"
)

func FixTime(m map[string]any) (string, string) {
	tm := m["created_at"].(time.Time)
	//tm = tm.Add(time.Hour * 8)
	timestamp := fmt.Sprintf("%s", tm)
	ago := timeago.English.Format(tm)
	return timestamp, ago
}
