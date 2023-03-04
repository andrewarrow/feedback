package sqlgen

import "fmt"

func MysqlCreateTable(tableName string) string {
	sql := `CREATE TABLE %s (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    username varchar(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE KEY unique_username (username)
) ENGINE InnoDB;`
	return fmt.Sprintf(sql, tableName)
}
