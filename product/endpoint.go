package product

import (
	"context"

	"github.com/GolangNorthwindRestApi/helper"
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

// @Summary Create new Product
// @Tags Product
// @Accept json
// @Produce json
// @Param addProductRequest body product.getAddProductRequest true "User data"
// @Success 200 {object} product.Product "New product created"
// @Router /products/ [post]
func makeAddProductEndPoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getAddProductRequest)
		result, err := s.InsertProduct(&req)
		helper.Catch(err)
		return result, nil
	}
}

// @Summary Delete Product by Id
// @Tags Product
// @Accept json
// @Produce json
// @Param id path int true "User data"
// @Success 200 {object} product.Product "Product deleted"
// @Router /products/{id} [delete]
func makeDeleteProductEndPoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteProductRequest)
		result, err := s.DeleteProductById(&req)
		helper.Catch(err)
		return result, nil
	}
}

// @Summary Get best selling Products
// @Tags Product
// @Accept json
// @Produce json
// @Success 200 {object} product.ProductTopResponse "Top ten best selling products"
// @Router /products/bestselling [get]
func makeGetBestSellingProductsEndPoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := s.GetBestSellingProducts()
		helper.Catch(err)
		return result, nil
	}
}

// @Summary Get Product by Id
// @Tags Product
// @Accept json
// @Produce json
// @Param id path int true "User data"
// @Success 200 {object} product.Product "Product"
// @Router /products/{id} [get]
func makeGetProductByIdEndPoint(s Service) endpoint.Endpoint {
	getProductByIdEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getProductByIDRequest)
		product, err := s.GetProductById(&req)
		helper.Catch(err)
		return product, nil
	}

	return getProductByIdEndpoint
}

// @Summary Products list
// @Tags Product
// @Accept json
// @Produce json
// @Param getProductsRequest body product.getProductsRequest true "User data"
// @Success 200 {object} product.ProductsList "Products paginated"
// @Router /products/paginated [post]
func makeGetProductsEndPoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getProductsRequest)
		result, err := s.GetProducts(&req)
		helper.Catch(err)
		return result, nil
	}
}

// @Summary Update Product
// @Tags Product
// @Accept json
// @Produce json
// @Param updateProductRequest body product.getUpdateProductRequest true "User data"
// @Success 200 {object} product.Product "Product updated"
// @Router /products [put]
func makeUpdateProductEndPoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getUpdateProductRequest)
		result, err := s.UpdateProduct(&req)
		helper.Catch(err)
		return result, nil
	}
}
