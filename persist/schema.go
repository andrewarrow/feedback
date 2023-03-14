package persist

import (
	"fmt"
	"os"
	"time"

	"github.com/andrewarrow/feedback/sqlgen"
	"github.com/jmoiron/sqlx"

	//_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func MysqlConnection() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	url := fmt.Sprintf("%s:%s@(%s:%s)/%s", dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName)
	db := sqlx.MustConnect("mysql", url)

	return db
}

func PostgresConnection() *sqlx.DB {

	url := os.Getenv("DATABASE_URL")

	db := sqlx.MustConnect("postgres", url)
	db.SetMaxOpenConns(30)
	db.SetMaxIdleConns(30)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db
}

func SchemaJson(db *sqlx.DB) string {
	db.Exec(sqlgen.PgCreateSchemaTable())
	prefix := os.Getenv("FEEDBACK_NAME")
	m := make(map[string]any)
	sql := fmt.Sprintf("select json_string from %s_feedback_schema limit 1", prefix)
	rows, _ := db.Queryx(sql)
	defer rows.Close()
	rows.Next()
	rows.MapScan(m)
	if len(m) == 0 {
		jsonString := `{"footer": "github.com/andrewarrow/feedback",
"title": "%s",
"models": [
  {"name": "user", "fields": [
    {"name": "username", "flavor": "username", "index": "unique"},
    {"name": "password", "flavor": "fewWords"},
    {"name": "guid", "flavor": "uuid", "index": "yes"}
  ]}
]}`
		jsonStringWithTitle := fmt.Sprintf(jsonString, prefix)
		db.Exec(fmt.Sprintf("insert into %s_feedback_schema (json_string) values ('%s')", prefix, jsonStringWithTitle))
		return jsonStringWithTitle
	}
	return fmt.Sprintf("%s", m["json_string"])
}
