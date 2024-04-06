package persist

import (
	"database/sql"
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

type CustomDB struct {
	*sqlx.DB
}

func (db *CustomDB) ExecWithLogging(query string, args ...interface{}) (sql.Result, error) {
	fmt.Printf("Executing SQL: %s\n", db.Rebind(query))

	sqlValues := make([]string, len(args))
	for i, arg := range args {
		switch v := arg.(type) {
		case int, int64, float64:
			sqlValues[i] = fmt.Sprintf("%v", v)
		case string:
			sqlValues[i] = fmt.Sprintf("'%v'", v) // Surround strings with single quotes
		case []byte:
			sqlValues[i] = fmt.Sprintf("E'\\x%X'", v) // Convert byte slices to SQL hex format
		case time.Time:
			sqlValues[i] = fmt.Sprintf("'%v'", v.Format("2006-01-02 15:04:05"))
		default:
			sqlValues[i] = "NULL"
		}
	}
	fmt.Println("SQL argument values:", sqlValues)

	return db.DB.Exec(query, args...)
}

func PostgresConnectionByUrl(url string) *sqlx.DB {
	var db *sqlx.DB
	var err error
	if os.Getenv("DEBUG") == "1" {
		fmt.Println("PostgresConnectionByUrl", url)
	}

	for {
		db, err = sqlx.Connect("postgres", url)
		if err == nil {
			if os.Getenv("DEBUG") == "1" {
				fmt.Println("db", db)
			}
			db.SetMaxOpenConns(9)
			db.SetMaxIdleConns(9)
			db.SetConnMaxLifetime(5 * time.Minute)
			break
		} else {
			fmt.Println(err.Error())
			time.Sleep(time.Second * 3)
		}
	}

	return db
}

func PostgresConnection(dbEnvVarName string) *sqlx.DB {

	if os.Getenv("DEBUG") == "1" {
		fmt.Println("PostgresConnection", dbEnvVarName)
	}
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
