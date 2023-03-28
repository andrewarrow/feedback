package prefix

var FeedbackName string

func Tablename(table string) string {
	prefix := FeedbackName
	if prefix != "" {
		return prefix + "_" + table
	}
	return table
}
