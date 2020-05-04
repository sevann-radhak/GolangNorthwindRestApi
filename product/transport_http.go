package product

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeHttpHandler(s Service) http.Handler {
	r := chi.NewRouter()

	getProductByIdHandler := kithttp.NewServer(
		makeGetProductByIdEndPoint(s),
		getProductByIdRequestDecoder,
		kithttp.EncodeJSONResponse)

	r.Method(http.MethodGet, "/{id}", getProductByIdHandler)

	getProductsHandler := kithttp.NewServer(
		makeGetProductsEndPoint(s),
		getProductsRequestDecoder,
		kithttp.EncodeJSONResponse)

	r.Method(http.MethodPost, "/paginated", getProductsHandler)

	addProductHandler := kithttp.NewServer(
		makeUpdateProductEndPoint(s),
		addProductRequestDecoder,
		kithttp.EncodeJSONResponse)

	r.Method(http.MethodPost, "/", addProductHandler)

	updateProductHandler := kithttp.NewServer(
		makeUpdateProductEndPoint(s),
		updateProductRequestDecoder,
		kithttp.EncodeJSONResponse)

	r.Method(http.MethodPut, "/", updateProductHandler)

	return r
}

func addProductRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := getAddProductRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		panic(err)
	}

	return request, nil
}

func getProductByIdRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	productId, _ := strconv.Atoi(chi.URLParam(r, "id"))
	return getProductByIDRequest{
		ProductID: productId,
	}, nil
}

func getProductsRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := getProductsRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		panic(err)
	}

	return request, nil
}

func updateProductRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := getUpdateProductRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		panic(err)
	}

	return request, nil
}
