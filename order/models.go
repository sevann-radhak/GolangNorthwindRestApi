package order

type OrderItem struct {
	Id         int               `json:"id"`
	Address    string            `json:"address"`
	City       string            `json:"city"`
	Company    string            `json:"company"`
	Customer   string            `json:"customer"`
	CustomerId string            `json:"customerId"`
	Data       []*OrderDetaiItem `json:"data"`
	OrderDate  string            `json:"orderDate"`
	Phone      string            `json:"phone"`
	StatusId   string            `json:"statusId"`
	StatusName string            `json:"statusName"`
}

type OrderItemDeleted struct {
	Data         *OrderItem `json:"data"`
	RowsAffected int        `json:"rowsAffected"`
}

type OrderDetaiItem struct {
	Id          int     `json:"id"`
	OrderId     int     `json:"orderId"`
	ProductId   int     `json:"productId"`
	ProductName string  `json:"productName"`
	Quantity    float64 `json:"quantity"`
	UnitPrice   float64 `json:"unitPrice"`
}

type OrderList struct {
	Data         []*OrderItem `json:"data"`
	TotalRecords int          `json:"totalRecords"`
}
