package main

import (
	"database/sql"
	"net/http"

	"github.com/GolangNorthwindRestApi/customer"
	"github.com/GolangNorthwindRestApi/database"
	"github.com/GolangNorthwindRestApi/employee"
	"github.com/GolangNorthwindRestApi/order"
	"github.com/GolangNorthwindRestApi/product"

	"github.com/go-chi/chi"

	_ "github.com/go-sql-driver/mysql"
)

var databaaseConnection *sql.DB

func main() {
	databaaseConnection = database.InitDB()
	defer databaaseConnection.Close()

	var (
		customerRepository = customer.NewRepository(databaaseConnection)
		employeeRepository = employee.NewRepository(databaaseConnection)
		orderRepository    = order.NewRepository(databaaseConnection)
		productRepository  = product.NewRepository(databaaseConnection)
	)

	var (
		customerService customer.Service
		employeeService employee.Service
		orderService    order.Service
		productService  product.Service
	)

	customerService = customer.NewService(customerRepository)
	employeeService = employee.NewService(employeeRepository)
	orderService = order.NewService(orderRepository)
	productService = product.NewService(productRepository)

	r := chi.NewRouter()

	r.Mount("/customers", customer.MakeHttpHandler(customerService))
	r.Mount("/employees", employee.MakeHttpHandler(employeeService))
	r.Mount("/orders", order.MakeHttpHandler(orderService))
	r.Mount("/products", product.MakeHttpHandler(productService))

	http.ListenAndServe(":3000", r)
}
