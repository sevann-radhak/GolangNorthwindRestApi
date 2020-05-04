package product

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type getAddProductRequest struct {
	Category     string
	Description  string
	ListPrice    float32
	StandardCost float32
	ProductCode  string
	ProductName  string
}

type getBestSellingsReqest struct{}

type getProductByIDRequest struct {
	ProductID int
}

type getProductsRequest struct {
	Limit  int
	Offset int
}

type getUpdateProductRequest struct {
	Id           int64
	Category     string
	Description  string
	ListPrice    float32
	StandardCost float32
	ProductCode  string
	ProductName  string
}

type deleteProductRequest struct {
	Id int
}

func makeDeleteProductEndPoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteProductRequest)
		result, err := s.DeleteProductById(&req)
		if err != nil {
			panic(err)
		}

		return result, nil
	}
}

func makeAddProductEndPoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getAddProductRequest)
		result, err := s.InsertProduct(&req)
		if err != nil {
			panic(err)
		}

		return result, nil
	}
}

func makeGetBestSellingProductsEndPoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := s.GetBestSellingProducts()
		if err != nil {
			panic(err)
		}

		return result, nil
	}
}

func makeGetProductByIdEndPoint(s Service) endpoint.Endpoint {
	getProductByIdEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getProductByIDRequest)
		product, err := s.GetProductById(&req)
		if err != nil {
			panic(err)
		}
		return product, nil
	}

	return getProductByIdEndpoint
}

func makeGetProductsEndPoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getProductsRequest)
		result, err := s.GetProducts(&req)
		if err != nil {
			panic(err)
		}

		return result, nil
	}
}

func makeUpdateProductEndPoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getUpdateProductRequest)
		result, err := s.UpdateProduct(&req)
		if err != nil {
			panic(err)
		}

		return result, nil
	}
}
