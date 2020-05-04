package main

import (
	"database/sql"
	"net/http"

	"github.com/GolangNorthwindRestApi/database"
	"github.com/GolangNorthwindRestApi/employee"
	"github.com/GolangNorthwindRestApi/product"

	"github.com/go-chi/chi"

	_ "github.com/go-sql-driver/mysql"
)

var databaaseConnection *sql.DB

func main() {
	databaaseConnection = database.InitDB()
	defer databaaseConnection.Close()

	var (
		employeeRepository = employee.NewRepository(databaaseConnection)
		productRepository  = product.NewRepository(databaaseConnection)
	)

	var (
		employeeService employee.Service
		productService  product.Service
	)

	employeeService = employee.NewService(employeeRepository)
	productService = product.NewService(productRepository)

	r := chi.NewRouter()

	r.Mount("/employees", employee.MakeHttpHandler(employeeService))
	r.Mount("/products", product.MakeHttpHandler(productService))

	http.ListenAndServe(":3000", r)
}
