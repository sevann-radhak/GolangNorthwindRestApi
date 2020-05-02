package database

import "database/sql"

func InitDb() *sql.DB {
	connectionString := "foot:admin@tcp(localhost:3306)/nothwind"

	databaaseConnection, err := sql.Open("mysql", connectionString)

	if err != nil {
		panic(err.Error())
	}

	return databaaseConnection
}
