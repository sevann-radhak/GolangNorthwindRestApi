package order

import (
	"github.com/GolangNorthwindRestApi/helper"
)

type Service interface {
	AddOrder(params *addOrderRequest) (*OrderItem, error)
	DeleteOrderById(param *deleteOrderRequest) (*OrderItemDeleted, error)
	DeleteOrderDetail(param *deleteOrderDetailRequest) (int64, error)
	GetOrderItemById(param *getOrderItemByIdRequest) (*OrderItem, error)
	GetOrders(params *getOrdersRequest) (*OrderList, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) AddOrder(params *addOrderRequest) (*OrderItem, error) {
	orderId, err := s.repo.AddOrder(params)
	helper.Catch(err)

	for _, detail := range params.OrderDetails {
		detail.OrderId = orderId
		_, err := s.repo.AddOrderDetail(&detail)
		helper.Catch(err)
	}

	orderIdItem := &getOrderItemByIdRequest{OrderId: int(orderId)}
	order, err := s.repo.GetOrderItemById(orderIdItem)
	helper.Catch(err)

	return order, nil
}

func (s *service) DeleteOrderById(param *deleteOrderRequest) (*OrderItemDeleted, error) {
	totalRows := 0

	orderIdItem := &getOrderItemByIdRequest{OrderId: param.OrderId}
	order, err := s.repo.GetOrderItemById(orderIdItem)
	helper.Catch(err)

	rows, err := s.repo.DeleteOrderDetailByOrderId(param)
	helper.Catch(err)
	totalRows += int(rows)

	rows, err = s.repo.DeleteOrderById(param)
	helper.Catch(err)
	totalRows += int(rows)

	orderDeleted := &OrderItemDeleted{Data: order, RowsAffected: totalRows}

	return orderDeleted, nil
}

func (s *service) DeleteOrderDetail(param *deleteOrderDetailRequest) (int64, error) {
	return s.repo.DeleteOrderDetail(param)
}

func (s *service) GetOrderItemById(param *getOrderItemByIdRequest) (*OrderItem, error) {
	return s.repo.GetOrderItemById(param)
}

func (s *service) GetOrders(params *getOrdersRequest) (*OrderList, error) {
	orders, err := s.repo.GetOrders(params)
	helper.Catch(err)
	totalOrders, err := s.repo.GetTotalOrders(params)
	helper.Catch(err)

	return &OrderList{Data: orders, TotalRecords: totalOrders}, nil
}
