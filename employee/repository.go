package employee

import (
	"database/sql"

	"github.com/GolangNorthwindRestApi/helper"
)

type Repository interface {
	AddEmploye(params *addEmployeeRequest) (int64, error)
	GetEmployeeById(param *getEmployeeByIdRequest) (*Employee, error)
	GetEmployees(params *getEmployeesRequest) ([]*Employee, error)
	GetTotalEmployees() (int, error)
	GetEmployeeTop() (*EmployeeTop, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(databaseConnection *sql.DB) Repository {
	return &repository{db: databaseConnection}
}

func (repo *repository) AddEmploye(params *addEmployeeRequest) (int64, error) {
	const sql = `
		INSERT INTO employees (
			address,
			business_phone,
			company,
			email_address,
			fax_number,
			first_name,
			home_phone,
			job_title,
			last_name,
			mobile_phone )
		VALUES (?,?,?,?,?,?,?,?,?,?)`

	result, err := repo.db.Exec(
		sql,
		params.Address,
		params.BusinessPhone,
		params.Company,
		params.EmailAddress,
		params.FaxNumber,
		params.FirstName,
		params.HomePhone,
		params.JobTitle,
		params.LastName,
		params.MobilePhone)

	helper.Catch(err)
	id, _ := result.LastInsertId()
	return id, nil
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
		&employee.JobTitle,
		&employee.BusinessPhone,
		&employee.HomePhone,
		&employee.MobilePhone,
		&employee.FaxNumber,
		&employee.Address)

	helper.Catch(err)

	return employee, nil
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
			&employee.JobTitle,
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

func (repo *repository) GetEmployeeTop() (*EmployeeTop, error) {
	const sql = `
		SELECT
			e.id,
			COUNT( e.id ) as totalSellings,
			e.first_name,
			e.last_name,
			e.email_address
		FROM northwind.orders o
		INNER JOIN northwind.employees e ON o.employee_id = e.id
		GROUP BY o.employee_id
		ORDER BY totalSellings desc
		limit 1;`

	row := repo.db.QueryRow(sql)
	employee := &EmployeeTop{}

	err := row.Scan(
		&employee.Id,
		&employee.TotalSellings,
		&employee.FirstName,
		&employee.LastName,
		&employee.EmailAddress)

	helper.Catch(err)

	return employee, nil
}

func (repo *repository) GetTotalEmployees() (int, error) {
	const sql = `SELECT COUNT(*) FROM northwind.employees;`

	var total int
	row := repo.db.QueryRow(sql)
	err := row.Scan(&total)
	helper.Catch(err)
	return total, nil
}
