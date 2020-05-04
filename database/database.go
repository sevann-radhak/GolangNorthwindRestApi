package database

import (
	"database/sql"

	"github.com/GolangNorthwindRestApi/helper"
)

func InitDB() *sql.DB {
	connectionString := "root:admin@tcp(localhost:3306)/northwind"
	databaseConnection, err := sql.Open("mysql", connectionString)

	helper.Catch(err)
	return databaseConnection
}
