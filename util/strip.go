package util

import (
	"strings"
)

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

func RemoveSensitiveKeys(m map[string]any) {
	for k, v := range m {
		switch vv := v.(type) {
		case map[string]any:
			RemoveSensitiveKeys(vv)
		case []map[string]any:
			for _, item := range vv {
				RemoveSensitiveKeys(item)
			}
		default:
			if k == "password" || strings.HasSuffix(k, "_id") || k == "id" {
				delete(m, k)
			}
		}
	}
}
