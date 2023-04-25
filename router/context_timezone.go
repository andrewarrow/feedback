package router

import (
	"time"

	"github.com/xeonx/timeago"
)

func (c *Context) TimezoneList(list []map[string]any,
	field1, field2 string, tz *time.Location) {
	cfg := timeago.English
	cfg.Max = 9223372036854775807

	for _, thing := range list {
		t1 := thing[field1].(int64)
		t2 := thing[field2].(int64)

		newT1 := time.Unix(t1, 0).In(tz)
		newT2 := time.Unix(t2, 0).In(tz)

		thing[field1] = newT1.Unix()
		thing[field2] = newT2.Unix()

		thing[field1+"_human"] = newT1.Format(HUMAN)
		thing[field2+"_human"] = newT2.Format(HUMAN)
	}
}
