package customer

import "github.com/GolangNorthwindRestApi/helper"

type Service interface {
	GetCustomers(params *getCustomersRequest) (*CustomersList, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetCustomers(params *getCustomersRequest) (*CustomersList, error) {
	customers, err := s.repo.GetCustomers(params)
	helper.Catch(err)
	totalCusomers, err := s.repo.GetTotalCustomers()
	helper.Catch(err)
	return &CustomersList{Data: customers, TotalRecords: totalCusomers}, nil
}
