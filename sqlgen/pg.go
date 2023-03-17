package sqlgen

import (
	"fmt"

	"github.com/andrewarrow/feedback/prefix"
)

func PgCreateTable(tableName string) string {
	sql := `CREATE TABLE %s (
  id SERIAL PRIMARY KEY,
	created_at TIMESTAMP NOT NULL DEFAULT NOW()
);`
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
