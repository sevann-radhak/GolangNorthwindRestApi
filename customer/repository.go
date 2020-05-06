package customer

import (
	"database/sql"

	"github.com/GolangNorthwindRestApi/helper"
)

type Repository interface {
	GetCustomers(params *getCustomersRequest) ([]*Customer, error)
	GetTotalCustomers() (int, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(databaseConnection *sql.DB) Repository {
	return &repository{db: databaseConnection}
}

func (repo *repository) GetCustomers(params *getCustomersRequest) ([]*Customer, error) {
	const sql = `
		SELECT 
			id,
			first_name,
			last_name,
			address,
			business_phone,
			city,
			company
		FROM northwind.customers
		ORDER BY last_name, first_name 
		LIMIT ? 
		OFFSET ?;`

	results, err := repo.db.Query(sql, params.Limit, params.Offset)
	helper.Catch(err)

	var customers []*Customer
	for results.Next() {
		customer := &Customer{}
		err = results.Scan(
			&customer.Id,
			&customer.FirstName,
			&customer.LastName,
			&customer.Address,
			&customer.BusinessPhone,
			&customer.City,
			&customer.Company)

		helper.Catch(err)
		customers = append(customers, customer)
	}

	return customers, nil
}

func (repo *repository) GetTotalCustomers() (int, error) {
	const sql = `SELECT COUNT(*) FROM northwind.customers;`

	var total int
	row := repo.db.QueryRow(sql)
	err := row.Scan(&total)
	helper.Catch(err)
	return total, nil
}
