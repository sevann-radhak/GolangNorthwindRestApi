package order

type OrderItem struct {
	Id         int               `json:"id"`
	Address    string            `json:"address"`
	City       string            `json:"city"`
	Company    string            `json:"company"`
	Customer   string            `json:"customer"`
	CustomerId string            `json:"customer_id"`
	Data       []*OrderDetaiItem `json:"data"`
	OrderDate  string            `json:"order_date"`
	Phone      string            `json:"phone"`
	StatusId   string            `json:"status_id"`
	StatusName string            `json:"status_name"`
}

type OrderItemDeleted struct {
	Data         *OrderItem `json:"data"`
	RowsAffected int        `json:"rows_affected"`
}

type OrderDetaiItem struct {
	Id          int     `json:"id"`
	OrderId     int     `json:"order_id"`
	ProductId   int     `json:"product_id"`
	ProductName string  `json:"product_name"`
	Quantity    float64 `json:"quantity"`
	UnitPrice   float64 `json:"unit_price"`
}

type OrderList struct {
	Data         []*OrderItem `json:"data"`
	TotalRecords int          `json:"total_records"`
}
