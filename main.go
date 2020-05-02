package main

import (
	"GolangNorthwindRestApi/database"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	databaaseConnection := database.InitDb()
	defer databaaseConnection.Close()

	fmt.Println(databaaseConnection)
}
