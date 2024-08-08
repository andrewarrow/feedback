package sqlgen

import (
	"fmt"

	"github.com/andrewarrow/feedback/prefix"
)

func PgCreateTable(tableName string, small bool) string {
	sql := `CREATE TABLE %s (
  id SERIAL PRIMARY KEY,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);`
	if small {
		sql = `CREATE TABLE %s ();`
	}
	return fmt.Sprintf(sql, tableName)
}

func FeedbackSchemaTable() string {
	return prefix.Tablename("feedback_schema")
}

func PgCreateSchemaTable() string {
	sql := `CREATE TABLE %s (
  id SERIAL PRIMARY KEY,
	json_string text
);`
	return fmt.Sprintf(sql, FeedbackSchemaTable())
}
