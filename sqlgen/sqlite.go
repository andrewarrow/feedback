package sqlgen

import "fmt"

func SqliteCreateTable(tableName string) string {
	sql := `CREATE TABLE %s (
  id INTEGER PRIMARY KEY,
	guid NOT NULL,
	created_at datetime CURRENT_TIMESTAMP,
	updated_at datetime CURRENT_TIMESTAMP
);`
	return fmt.Sprintf(sql, tableName)
}
