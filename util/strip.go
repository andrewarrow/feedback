package util

import "strings"

func StripFields(row map[string]any) map[string]any {

	for k, _ := range row {
		if k == "password" || k == "id" {
			delete(row, k)
		} else if strings.HasSuffix(k, "_id") {
			delete(row, k)
		}
	}

	return row

}
