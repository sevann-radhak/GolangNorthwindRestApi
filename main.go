package main

// httpSwagger "github.com/swaggo/http-swagger"

import (
	"net/http"

	"github.com/GolangNorthwindRestApi/customer"
	"github.com/GolangNorthwindRestApi/database"
	"github.com/GolangNorthwindRestApi/employee"
	"github.com/GolangNorthwindRestApi/helper"
	"github.com/GolangNorthwindRestApi/order"
	"github.com/GolangNorthwindRestApi/product"

	"github.com/go-chi/chi"

	_ "github.com/go-sql-driver/mysql"

	_ "github.com/GolangNorthwindRestApi/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

// var databaaseConnection *sql.DB

// @title Products with Golang - API
// @version 1.0
// @description This is a sample server celler developed with Golang and MariaDB
// @contact.name Sevann Radhak
// @contact.url http://github.com/sevannr
// @contact.email sevann.radhak@gmail.com

func main() {
	databaaseConnection := database.InitDB()
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
	r.Use(helper.GetCors().Handler)

	r.Mount("/customers", customer.MakeHttpHandler(customerService))
	r.Mount("/employees", employee.MakeHttpHandler(employeeService))
	r.Mount("/orders", order.MakeHttpHandler(orderService))
	r.Mount("/products", product.MakeHttpHandler(productService))

	r.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL("../swagger/doc.json")))

	http.ListenAndServe(":3000", r)
}
