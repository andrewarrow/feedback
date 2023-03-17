package prefix

import "os"

func Tablename(table string) string {
	prefix := os.Getenv("FEEDBACK_NAME")
	if prefix != "" {
		return prefix + "_" + table
	}
	return table
}
