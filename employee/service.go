package employee

import "github.com/GolangNorthwindRestApi/helper"

type Service interface {
	AddEmploye(params *addEmployeeRequest) (*Employee, error)
	GetEmployeeById(params *getEmployeeByIdRequest) (*Employee, error)
	GetEmployees(params *getEmployeesRequest) (*EmployeesList, error)
	GetEmployeeTop() (*EmployeeTop, error)
	UpdateEmployee(params *updateEmployeeRequest) (*Employee, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) AddEmploye(params *addEmployeeRequest) (*Employee, error) {
	idEmployee, err := s.repo.AddEmploye(params)
	helper.Catch(err)

	return s.repo.GetEmployeeById(&getEmployeeByIdRequest{EmployeId: int(idEmployee)})
}

func (s *service) GetEmployeeById(params *getEmployeeByIdRequest) (*Employee, error) {
	return s.repo.GetEmployeeById(params)
}

func (s *service) GetEmployees(params *getEmployeesRequest) (*EmployeesList, error) {
	employees, err := s.repo.GetEmployees(params)
	helper.Catch(err)
	totalEmployees, err := s.repo.GetTotalEmployees()
	helper.Catch(err)
	return &EmployeesList{Data: employees, TotalRecords: totalEmployees}, nil
}

func (s *service) GetEmployeeTop() (*EmployeeTop, error) {
	employee, err := s.repo.GetEmployeeTop()
	helper.Catch(err)
	return employee, nil
}

func (s *service) UpdateEmployee(params *updateEmployeeRequest) (*Employee, error) {
	idEmployee, err := s.repo.UpdateEmployee(params)
	helper.Catch(err)
	return s.repo.GetEmployeeById(&getEmployeeByIdRequest{EmployeId: int(idEmployee)})
}
