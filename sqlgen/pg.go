package sqlgen

import (
	"fmt"
	"os"
)

func PgCreateTable(tableName string) string {
	sql := `CREATE TABLE %s (
  id SERIAL PRIMARY KEY,
	created_at TIMESTAMP NOT NULL DEFAULT NOW()
);`
	return fmt.Sprintf(sql, tableName)
}

func PgCreateSchemaTable() string {
	prefix := os.Getenv("FEEDBACK_NAME")
	sql := `CREATE TABLE %s_feedback_schema (
  id SERIAL PRIMARY KEY,
	json_string text
);`
	return fmt.Sprintf(sql, prefix)
}
