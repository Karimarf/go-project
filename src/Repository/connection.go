package Repository

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"src/Types"
	"strings"
)

func getMysqlVariables() Types.MysqlVariables {
	pathEnFile := "../.env"
	file, err := os.ReadFile(pathEnFile)
	if err != nil {
		panic(err.Error())
	}
	strFile := string(file)
	lines := strings.Split(strFile, "\n")
	var mysqlVariables Types.MysqlVariables
	for _, line := range lines {
		if strings.Contains(line, "MYSQL_HOST") {
			mysqlVariables.Host = strings.Split(line, "=")[1]
		}
		if strings.Contains(line, "MYSQL_PORT") {
			mysqlVariables.Port = strings.Split(line, "=")[1]
		}
		if strings.Contains(line, "MYSQL_USER") {
			mysqlVariables.User = strings.Split(line, "=")[1]
		}
		if strings.Contains(line, "MYSQL_PASSWORD") {
			mysqlVariables.Password = strings.Split(line, "=")[1]
		}
		if strings.Contains(line, "MYSQL_DATABASE") {
			mysqlVariables.DBName = strings.Split(line, "=")[1]
		}
	}
	return mysqlVariables
}

func createConnection() *sql.DB {
	mysqlVariables := getMysqlVariables()
	db, err := sql.Open("mysql", mysqlVariables.User+":"+mysqlVariables.Password+"@tcp(localhost:"+mysqlVariables.Port+")/"+mysqlVariables.DBName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func closeConnection(db *sql.DB) {
	err := db.Close()
	if err != nil {
		return
	}
}
