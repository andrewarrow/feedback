package persist

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func ScanSchema() {
	db := PostgresConnection()
	sql := "SELECT tablename FROM pg_catalog.pg_tables WHERE schemaname = 'public';"

	rows := SelectAll(db, sql)
	for _, row := range rows {
		table := fmt.Sprintf("%s", row["tablename"])
		//single := util.Unplural(table)
		scanTable(db, table)
	}

}

func scanTable(db *sqlx.DB, table string) {
	fmt.Println(table)
	sql := fmt.Sprintf("SELECT column_name, data_type FROM information_schema.columns WHERE table_name = '%s'", table)
	rows := SelectAll(db, sql)
	for _, row := range rows {
		col := fmt.Sprintf("%s", row["column_name"])
		flavor := fmt.Sprintf("%s", row["data_type"])
		fmt.Printf("     %s %s\n", col, flavor)
	}
}

func SelectAll(db *sqlx.DB, sql string) []map[string]any {
	ms := []map[string]any{}
	rows, err := db.Queryx(sql)
	if err != nil {
		return ms
	}
	defer rows.Close()
	for rows.Next() {
		m := make(map[string]any)
		rows.MapScan(m)
		ms = append(ms, m)
	}
	return ms
}
