package main

import (
	"database/sql"
	"net/http"

	"github.com/GolangNorthwindRestApi/database"
	"github.com/GolangNorthwindRestApi/product"

	"github.com/go-chi/chi"

	_ "github.com/go-sql-driver/mysql"
)

var databaaseConnection *sql.DB

func main() {
	databaaseConnection = database.InitDB()
	defer databaaseConnection.Close()

	var productRepository = product.NewRepository(databaaseConnection)
	var productService product.Service
	productService = product.NewService(productRepository)

	r := chi.NewRouter()
	r.Mount("/products", product.MakeHttpHandler(productService))

	http.ListenAndServe(":3000", r)
}
