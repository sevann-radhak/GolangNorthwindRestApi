package order

import (
	"context"

	"github.com/GolangNorthwindRestApi/helper"
	"github.com/go-kit/kit/endpoint"
)

type addOrderDetailRequest struct {
	OrderId   int64
	ProductId int
	Quantity  float64
	UnitPrice float64
}

type addOrderRequest struct {
	CustomerId   int
	OrderDate    string
	OrderDetails []addOrderDetailRequest
}

type deleteOrderDetailRequest struct {
	OrderDetailId string
}

type deleteOrderRequest struct {
	OrderId int
}

type getOrderItemByIdRequest struct {
	OrderId int
}

type getOrdersRequest struct {
	Limit    int
	Offset   int
	Status   interface{}
	DateFrom interface{}
	DateTo   interface{}
}

func makeAddOrderEndPoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addOrderRequest)
		result, err := s.AddOrder(&req)
		helper.Catch(err)
		return result, nil
	}
}

func makeDeleteOrderDetailEndPoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteOrderDetailRequest)
		result, err := s.DeleteOrderDetail(&req)
		helper.Catch(err)
		return result, nil
	}
}

func makeDeleteOrderEndPoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteOrderRequest)
		result, err := s.DeleteOrderById(&req)
		helper.Catch(err)
		return result, nil
	}
}

func makeGetOrderItemByIdEndPoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getOrderItemByIdRequest)
		result, err := s.GetOrderItemById(&req)
		helper.Catch(err)
		return result, nil
	}
}

func makeGetOrdersEndPoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getOrdersRequest)
		result, err := s.GetOrders(&req)
		helper.Catch(err)
		return result, nil
	}
}
