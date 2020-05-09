package employee

import (
	"context"

	"github.com/GolangNorthwindRestApi/helper"
	"github.com/go-kit/kit/endpoint"
)

type addEmployeeRequest struct {
	Address       string
	BusinessPhone string
	Company       string
	EmailAddress  string
	FaxNumber     string
	FirstName     string
	HomePhone     string
	JobTitle      string
	LastName      string
	MobilePhone   string
}

type deleteEmployeeByIdRequest struct {
	EmployeId int
}

type getEmployeeByIdRequest struct {
	EmployeId int
}

type getEmployeesRequest struct {
	Limit  int
	Offset int
}

type getEmployeeTopRequest struct{}

type updateEmployeeRequest struct {
	Address       string
	BusinessPhone string
	Company       string
	EmailAddress  string
	FaxNumber     string
	FirstName     string
	HomePhone     string
	Id            int
	JobTitle      string
	LastName      string
	MobilePhone   string
}

// @Summary Create new Emplyee
// @Tags Employee
// @Accept json
// @Produce json
// @Param addEmployeeRequest body employee.addEmployeeRequest true "User data"
// @Success 200 {object} employee.Employee "New employee created"
// @Router /employees/ [post]
func makeAddEmployeeEndPoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addEmployeeRequest)
		result, err := s.AddEmploye(&req)
		helper.Catch(err)
		return result, nil
	}
}

// @Summary Delete Emplyee by Id
// @Tags Employee
// @Accept json
// @Produce json
// @Param id path int true "Employee Id"
// @Success 200 {object} employee.Employee "Employee"
// @Router /employees/{id} [delete]
func makeDeleteEmployeeEndPoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteEmployeeByIdRequest)
		result, err := s.DeleteEmployeeById(&req)
		helper.Catch(err)
		return result, nil
	}
}

// @Summary Emplyee by Id
// @Tags Employee
// @Accept json
// @Produce json
// @Param id path int true "Employee Id"
// @Success 200 {object} employee.Employee "Employee"
// @Router /employees/{id} [get]
func makeGetEmployeeByIdEndPoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getEmployeeByIdRequest)
		result, err := s.GetEmployeeById(&req)
		helper.Catch(err)
		return result, nil
	}
}

// @Summary Emplyees list
// @Tags Employee
// @Accept json
// @Produce json
// @Param getEmployees body employee.getEmployeesRequest true "User data"
// @Success 200 {object} employee.EmployeesList "Paginate list of employees"
// @Router /employees/paginated [post]
func makeGetEmployeesEndPoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getEmployeesRequest)
		result, err := s.GetEmployees(&req)
		helper.Catch(err)
		return result, nil
	}
}

// @Summary Best Emplyee
// @Tags Employee
// @Accept json
// @Produce json
// @Success 200 {object} employee.EmployeeTop "Best Employee"
// @Router /employees/best-seller [get]
func makeGetEmployeeTopEndPoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := s.GetEmployeeTop()
		helper.Catch(err)
		return result, nil
	}
}

// @Summary Update an Emplyee
// @Tags Employee
// @Accept json
// @Produce json
// @Param updateEmployeeRequest body employee.updateEmployeeRequest true "User data"
// @Success 200 {object} employee.Employee "Employee updated"
// @Router /employees/ [put]
func makeUpdateEmployeeEndPoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateEmployeeRequest)
		result, err := s.UpdateEmployee(&req)
		helper.Catch(err)
		return result, nil
	}
}
