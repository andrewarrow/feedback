package persist

import (
	"fmt"
	"os"
	"time"

	"github.com/andrewarrow/feedback/prefix"
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

func PostgresConnection(dbEnvVarName string) *sqlx.DB {

	url := os.Getenv(dbEnvVarName)

	db, err := sqlx.Connect("postgres", url)
	if err == nil {
		db.SetMaxOpenConns(30)
		db.SetMaxIdleConns(30)
		db.SetConnMaxLifetime(5 * time.Minute)
	}

	return db
}

func SchemaJson(db *sqlx.DB) string {
	if db == nil {
		return ""
	}
	db.Exec(sqlgen.PgCreateSchemaTable())
	m := make(map[string]any)
	sql := fmt.Sprintf("select json_string from %s limit 1", sqlgen.FeedbackSchemaTable())
	rows, _ := db.Queryx(sql)
	defer rows.Close()
	rows.Next()
	rows.MapScan(m)
	if len(m) == 0 {
		jsonStringWithTitle := DefaultJsonModels()
		db.Exec(fmt.Sprintf("insert into %s (json_string) values ('%s')", sqlgen.FeedbackSchemaTable(), jsonStringWithTitle))
		return jsonStringWithTitle
	}
	return fmt.Sprintf("%s", m["json_string"])
}

func DefaultJsonModels() string {
	jsonString := `{"footer": "github.com/andrewarrow/feedback",
"title": "%s",
"models": [
  {"name": "user", "fields": [
		{"name": "username", "flavor": "username", "index": "unique", "regex": "^[a-zA-Z0-9_]{2,20}$"},
		{"name": "password", "flavor": "fewWords", "regex": "^.{8,100}$"},
    {"name": "id", "flavor": "int"},
		{"name": "created_at", "flavor": "timestamp", "index": "yes"},
    {"name": "guid", "flavor": "uuid", "index": "yes"}
  ]}
]}`
	prefix := prefix.FeedbackName
	return fmt.Sprintf(jsonString, prefix)
}
