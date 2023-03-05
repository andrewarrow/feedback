package persist

import (
	"fmt"
	"os"

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

	return db
}

func SchemaJson(db *sqlx.DB) string {
	db.Exec(sqlgen.PgCreateSchemaTable())
	m := make(map[string]any)
	rows, _ := db.Queryx("select json_string from feedback_schema limit 1")
	rows.Next()
	rows.MapScan(m)
	if len(m) == 0 {
		jsonString := `{"phone": "github.com/andrewarrow/feedback",
"models": [
  {"name": "user", "fields": [
    {"name": "username", "flavor": "username", "index": "unique"},
    {"name": "password", "flavor": "fewWords"},
    {"name": "guid", "flavor": "uuid", "index": "yes"}
  ]},
  {"name": "story", "fields": [
    {"name": "title", "flavor": "fewWords"},
    {"name": "url", "flavor": "fewWords"},
		{"name": "username", "flavor": "fewWords", "index": "yes"},
    {"name": "body", "flavor": "text"},
    {"name": "comments", "flavor": "int"},
    {"name": "guid", "flavor": "uuid", "index": "yes"}
  ]},
  {"name": "comment", "fields": [
		{"name": "username", "flavor": "fewWords", "index": "yes"},
    {"name": "body", "flavor": "text"},
		{"name": "story_id", "flavor": "int", "index": "yes"},
		{"name": "story_guid", "flavor": "uuid"},
    {"name": "guid", "flavor": "uuid", "index": "yes"}
  ]}
]}`
		db.Exec(fmt.Sprintf("insert into feedback_schema (json_string) values ('%s')", jsonString))
		return jsonString
	}
	return fmt.Sprintf("%s", m["json_string"])
}
