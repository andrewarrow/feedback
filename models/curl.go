package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/andrewarrow/feedback/util"
)

func (m *Model) CurlCreatePayload() string {
	buffer := []string{}
	q := `"`
	for _, field := range m.Fields {
		if field.Required == "" {
			continue
		}
		name := field.Name
		if strings.HasSuffix(name, "_id") {
			tokens := strings.Split(name, "_")
			name = tokens[0] + "_guid"
		}
		val := q + "some_string" + q
		if strings.HasSuffix(name, "_guid") {
			val = q + util.PseudoUuid() + q
		} else if field.Flavor == "timestamp" {
			val = fmt.Sprintf("%d", time.Now().Unix())
		}

		buffer = append(buffer, fmt.Sprintf("%s%s%s: %s", q, name, q, val))
	}
	list := strings.Join(buffer, ",")
	return "'{" + list + "}'"
}
