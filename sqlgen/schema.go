package sqlgen

import "fmt"

func UpdateSchema(asBytes []byte) string {
	return fmt.Sprintf("update %s set json_string = '%s'", FeedbackSchemaTable(), string(asBytes))
}
