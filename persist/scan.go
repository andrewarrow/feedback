package persist

import (
	"fmt"

	"github.com/andrewarrow/feedback/files"
	"github.com/andrewarrow/feedback/models"
	"github.com/andrewarrow/feedback/util"
	"github.com/jmoiron/sqlx"
)

func ScanSchema() []*models.Model {
	db := PostgresConnection("DATABASE_URL")
	sql := "SELECT tablename FROM pg_catalog.pg_tables WHERE schemaname = 'public';"

	list := []*models.Model{}
	rows := SelectAll(db, sql)
	for _, row := range rows {
		table := fmt.Sprintf("%s", row["tablename"])
		single := util.Unplural(table)
		m := models.Model{}
		m.Name = single
		m.Fields = scanTable(db, table)
		list = append(list, &m)
	}

	return list
}

func scanTable(db *sqlx.DB, table string) []*models.Field {
	list := []*models.Field{}
	sql := fmt.Sprintf("SELECT column_name, data_type FROM information_schema.columns WHERE table_name = '%s'", table)
	rows := SelectAll(db, sql)
	for _, row := range rows {
		col := fmt.Sprintf("%s", row["column_name"])
		dt := fmt.Sprintf("%s", row["data_type"])
		field := models.Field{}
		field.Name = col
		field.Flavor = models.TypeToFlavor(dt)
		list = append(list, &field)
	}
	return list
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

func SaveSchema(asBytes []byte) {
	fmt.Println(string(asBytes))
	files.SaveFile("feedback.json", string(asBytes))
}
