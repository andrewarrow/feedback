package persist

import (
	"fmt"
	"os"
	"time"

	"github.com/andrewarrow/feedback/prefix"
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

func PostgresConnectionByUrl(url string) *sqlx.DB {

	db, err := sqlx.Connect("postgres", url)
	if err == nil {
		db.SetMaxOpenConns(9)
		db.SetMaxIdleConns(9)
		db.SetConnMaxLifetime(5 * time.Minute)
	} else {
		fmt.Println(err)
	}

	return db
}

func PostgresConnection(dbEnvVarName string) *sqlx.DB {

	url := os.Getenv(dbEnvVarName)
	if url == "" {
		return nil
	}
	return PostgresConnectionByUrl(url)
}

func DefaultJsonModels() string {
	jsonString := `{"footer": "github.com/andrewarrow/feedback",
"title": "%s",
  "routes": [{"root": "sessions", "paths": [
                     {"verb": "GET", "second": "", "third": ""},
                     {"verb": "GET", "second": "*", "third": ""},
                     {"verb": "POST", "second": "", "third": ""}
             ]},
             {"root": "users", "paths": [
                     {"verb": "GET", "second": "", "third": ""},
                     {"verb": "GET", "second": "*", "third": ""},
                     {"verb": "POST", "second": "", "third": ""}
             ]}
  ],
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
