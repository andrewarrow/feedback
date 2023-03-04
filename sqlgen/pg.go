package sqlgen

import "fmt"

func PgCreateTable(tableName string) string {
	sql := `CREATE TABLE %s (
  id SERIAL PRIMARY KEY
);`
	return fmt.Sprintf(sql, tableName)
}
