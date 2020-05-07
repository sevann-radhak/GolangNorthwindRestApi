package order

import (
	"database/sql"
	"fmt"

	"github.com/GolangNorthwindRestApi/helper"
)

type Repository interface {
	AddOrder(params *addOrderRequest) (int64, error)
	AddOrderDetail(params *addOrderDetailRequest) (int64, error)
	DeleteOrderById(param *deleteOrderRequest) (int64, error)
	DeleteOrderDetail(param *deleteOrderDetailRequest) (int64, error)
	DeleteOrderDetailByOrderId(param *deleteOrderRequest) (int64, error)
	GetOrderItemById(param *getOrderItemByIdRequest) (*OrderItem, error)
	GetOrders(params *getOrdersRequest) ([]*OrderItem, error)
	GetTotalOrders(params *getOrdersRequest) (int, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(databaseConnection *sql.DB) Repository {
	return &repository{db: databaseConnection}
}

func (repo *repository) AddOrder(params *addOrderRequest) (int64, error) {
	const sql = `
		INSERT INTO orders (
			customer_id,
			order_date )
		VALUES( ?, ? )`

	result, err := repo.db.Exec(sql, params.CustomerId, params.OrderDate)
	helper.Catch(err)

	id, err := result.LastInsertId()
	helper.Catch(err)

	return id, nil
}

func (repo *repository) AddOrderDetail(params *addOrderDetailRequest) (int64, error) {
	const sql = `
		INSERT INTO order_details(
			order_id,
			product_id,
			quantity,
			unit_price )
		VALUES ( ?, ?, ?, ? )`

	result, err := repo.db.Exec(sql, params.OrderId, params.ProductId, params.Quantity, params.UnitPrice)
	helper.Catch(err)

	id, err := result.LastInsertId()
	helper.Catch(err)

	return id, nil
}

func (repo *repository) DeleteOrderById(param *deleteOrderRequest) (int64, error) {
	const sql = `DELETE FROM orders WHERE id = ?`
	result, err := repo.db.Exec(sql, param.OrderId)
	helper.Catch(err)

	count, err := result.RowsAffected()
	helper.Catch(err)
	return count, nil
}

func (repo *repository) DeleteOrderDetail(param *deleteOrderDetailRequest) (int64, error) {
	const sql = `DELETE FROM order_details WHERE id = ?`
	result, err := repo.db.Exec(sql, param.OrderDetailId)
	helper.Catch(err)

	count, err := result.RowsAffected()
	helper.Catch(err)
	return count, nil
}

func (repo *repository) DeleteOrderDetailByOrderId(param *deleteOrderRequest) (int64, error) {
	const sql = `DELETE FROM order_details WHERE order_id = ?`
	result, err := repo.db.Exec(sql, param.OrderId)
	helper.Catch(err)

	count, err := result.RowsAffected()
	helper.Catch(err)
	return count, nil
}

func (repo *repository) GetOrderItemById(param *getOrderItemByIdRequest) (*OrderItem, error) {
	const sql = `
		SELECT 
			o.id,
			c.address,
			c.city,
			c.company,
			CONCAT( c.first_name, ' ', c.last_name ),
			o.customer_id,
			o.order_date,
			c.business_phone,
			o.status_id,
			os.status_name
		FROM northwind.orders o
		INNER JOIN northwind.orders_status os ON o.status_id = os.id
		INNER JOIN customers c ON o.customer_id = c.id
		WHERE o.id = ?;`

	row := repo.db.QueryRow(sql, param.OrderId)
	order := &OrderItem{}

	err := row.Scan(
		&order.Id,
		&order.Address,
		&order.City,
		&order.Company,
		&order.Customer,
		&order.CustomerId,
		&order.OrderDate,
		&order.Phone,
		&order.StatusId,
		&order.StatusName)

	helper.Catch(err)

	orderDetail, err := GetOrderDetailsItemById(repo, param)
	helper.Catch(err)
	order.Data = orderDetail

	return order, nil
}

func GetOrderDetailsItemById(repo *repository, param *getOrderItemByIdRequest) ([]*OrderDetaiItem, error) {
	const sql = `
		SELECT 
			od.id,
			order_id,
			product_id,
			p.product_name,
			quantity,
			unit_price
		FROM northwind.order_details od
		INNER JOIN products p ON od.product_id = p.id
		WHERE od.order_id = ?;`

	results, err := repo.db.Query(sql, param.OrderId)
	helper.Catch(err)

	var details []*OrderDetaiItem
	for results.Next() {
		detail := &OrderDetaiItem{}
		err = results.Scan(
			&detail.Id,
			&detail.OrderId,
			&detail.ProductId,
			&detail.ProductName,
			&detail.Quantity,
			&detail.UnitPrice)

		helper.Catch(err)
		details = append(details, detail)
	}

	return details, nil
}

func (repo *repository) GetOrders(params *getOrdersRequest) ([]*OrderItem, error) {
	filter := GetFilter(params)

	var sql = `
		SELECT
			o.id,
			CONCAT( c.first_name, ' ', c.last_name ) as customer,
			o.customer_id,
			o.order_date,
			o.status_id,
			os.status_name
		FROM northwind.orders o
		INNER JOIN northwind.orders_status os ON o.status_id = os.id
		INNER JOIN northwind.customers c ON o.customer_id = c.id
		WHERE 1 = 1 ` + filter + `LIMIT ? OFFSET ?;`

	results, err := repo.db.Query(sql, params.Limit, params.Offset)
	helper.Catch(err)

	var orders []*OrderItem
	for results.Next() {
		order := &OrderItem{}
		err = results.Scan(
			&order.Id,
			&order.Customer,
			&order.CustomerId,
			&order.OrderDate,
			&order.StatusId,
			&order.StatusName)

		helper.Catch(err)

		orderDetailRequest := &getOrderItemByIdRequest{OrderId: order.Id}

		orderDetail, err := GetOrderDetailsItemById(repo, orderDetailRequest)
		helper.Catch(err)

		order.Data = orderDetail

		orders = append(orders, order)
	}

	return orders, nil
}

func (repo *repository) GetTotalOrders(params *getOrdersRequest) (int, error) {
	filter := GetFilter(params)
	var sql = `SELECT COUNT(*) FROM northwind.orders o WHERE 1 = 1 ` + filter + `;`

	var total int
	row := repo.db.QueryRow(sql)
	err := row.Scan(&total)
	helper.Catch(err)
	return total, nil
}

func GetFilter(params *getOrdersRequest) string {
	var filter string

	if params.Status != nil {
		filter += fmt.Sprintf(" AND o.status_id = %v ", params.Status.(float64))
	}

	if params.DateFrom != nil && params.DateTo == nil {
		filter += fmt.Sprintf(" AND o.order_date >= '%v' ", params.DateFrom.(string))
	}

	if params.DateFrom == nil && params.DateTo != nil {
		filter += fmt.Sprintf(" AND o.order_date <= '%v' ", params.DateTo.(string))
	}

	if params.DateFrom != nil && params.DateTo != nil {
		filter += fmt.Sprintf(
			" AND o.order_date between '%v' and '%v' ",
			params.DateFrom.(string),
			params.DateFrom.(string))
	}

	return filter
}
