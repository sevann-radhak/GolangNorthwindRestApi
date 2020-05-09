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

// @Summary Create new Order
// @Tags Order
// @Accept json
// @Produce json
// @Param addOrderRequest body order.addOrderRequest true "User data"
// @Success 200 {object} order.OrderItem "New order created"
// @Router /orders/ [post]
func makeAddOrderEndPoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addOrderRequest)
		result, err := s.AddOrder(&req)
		helper.Catch(err)
		return result, nil
	}
}

// @Summary Delete OrderDetail
// @Tags Order
// @Accept json
// @Produce json
// @Param orderDetailId path int true "User data"
// @Success 200  {integer} int "ok"
// @Router /orders/{id}/detail/{orderDetailId} [delete]
func makeDeleteOrderDetailEndPoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteOrderDetailRequest)
		result, err := s.DeleteOrderDetail(&req)
		helper.Catch(err)
		return result, nil
	}
}

// @Summary Delete Order
// @Tags Order
// @Accept json
// @Produce json
// @Param id path int true "User data"
// @Success 200  {object} order.OrderItemDeleted "Order deleted and rows affected"
// @Router /orders/{id} [delete]
func makeDeleteOrderEndPoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteOrderRequest)
		result, err := s.DeleteOrderById(&req)
		helper.Catch(err)
		return result, nil
	}
}

// @Summary Get Order by Id
// @Tags Order
// @Accept json
// @Produce json
// @Param id path int true "User data"
// @Success 200  {object} order.OrderItem "Order with details"
// @Router /orders/{id} [get]
func makeGetOrderItemByIdEndPoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getOrderItemByIdRequest)
		result, err := s.GetOrderItemById(&req)
		helper.Catch(err)
		return result, nil
	}
}

// @Summary Get Orders
// @Tags Order
// @Accept json
// @Produce json
// @Param getOrdersRequest body order.getOrdersRequest true "User data"
// @Success 200  {object} order.OrderList "Orders list paginated"
// @Router /orders/paginated [post]
func makeGetOrdersEndPoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getOrdersRequest)
		result, err := s.GetOrders(&req)
		helper.Catch(err)
		return result, nil
	}
}
