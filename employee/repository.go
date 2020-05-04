package employee

import (
	"database/sql"

	"github.com/GolangNorthwindRestApi/helper"
)

type Repository interface {
	GetEmployeeById(param *getEmployeeByIdRequest) (*Employee, error)
	GetEmployees(params *getEmployeesRequest) ([]*Employee, error)
	GetTotalEmployees() (int, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(databaseConnection *sql.DB) Repository {
	return &repository{db: databaseConnection}
}

func (repo *repository) GetEmployeeById(param *getEmployeeByIdRequest) (*Employee, error) {
	const sql = `
		SELECT 
			id,
			company,
			last_name,
			first_name,
			email_address,
			job_title,
			business_phone,
			home_phone,
			COALESCE(mobile_phone,''),
			fax_number,
			address
		FROM northwind.employees
		WHERE id = ?;`

	row := repo.db.QueryRow(sql, param.EmployeId)
	employee := &Employee{}

	err := row.Scan(
		&employee.Id,
		&employee.Company,
		&employee.LastName,
		&employee.FirstName,
		&employee.EmailAddress,
		&employee.JotTitle,
		&employee.BusinessPhone,
		&employee.HomePhone,
		&employee.MobilePhone,
		&employee.FaxNumber,
		&employee.Address)

	return employee, err
}

func (repo *repository) GetEmployees(params *getEmployeesRequest) ([]*Employee, error) {
	const sql = `
		SELECT 
			id,
			company,
			last_name,
			first_name,
			email_address,
			job_title,
			business_phone,
			home_phone,
			COALESCE(mobile_phone,''),
			fax_number,
			address
		FROM northwind.employees
		ORDER BY last_name, first_name 
		LIMIT ? 
		OFFSET ?;`

	results, err := repo.db.Query(sql, params.Limit, params.Offset)
	helper.Catch(err)

	var employees []*Employee
	for results.Next() {
		employee := &Employee{}
		err = results.Scan(
			&employee.Id,
			&employee.Company,
			&employee.LastName,
			&employee.FirstName,
			&employee.EmailAddress,
			&employee.JotTitle,
			&employee.BusinessPhone,
			&employee.HomePhone,
			&employee.MobilePhone,
			&employee.FaxNumber,
			&employee.Address)

		helper.Catch(err)
		employees = append(employees, employee)
	}

	return employees, nil
}

func (repo *repository) GetTotalEmployees() (int, error) {
	const sql = `SELECT COUNT(*) FROM northwind.employees;`

	var total int
	row := repo.db.QueryRow(sql)
	err := row.Scan(&total)
	helper.Catch(err)
	return total, nil
}
