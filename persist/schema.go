package persist

import (
	"fmt"
	"os"

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
