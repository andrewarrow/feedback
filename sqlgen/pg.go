package sqlgen

import "fmt"

func PgCreateTable(tableName string) string {
	sql := `CREATE TABLE %s (
  id SERIAL PRIMARY KEY,
	created_at TIMESTAMP NOT NULL DEFAULT NOW()
);`
	return fmt.Sprintf(sql, tableName)
}

func PgCreateSchemaTable() string {
	sql := `CREATE TABLE feedback_schema (
  id SERIAL PRIMARY KEY,
	json_string text
);`
	return sql
}
