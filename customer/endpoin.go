package customer

import (
	"context"

	"github.com/GolangNorthwindRestApi/helper"
	"github.com/go-kit/kit/endpoint"
)

type getCustomersRequest struct {
	Limit  int
	Offset int
}

// @Summary Customers list
// @Tags Customer
// @Accept json
// @Produce json
// @Param getCustomers body customer.getCustomersRequest true "User data"
// @Success 200 {object} customer.CustomersList "Paginate list of customers"
// @Router /customers/paginated [post]
func makeGetCustomersEndPoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getCustomersRequest)
		result, err := s.GetCustomers(&req)
		helper.Catch(err)
		return result, nil
	}
}
