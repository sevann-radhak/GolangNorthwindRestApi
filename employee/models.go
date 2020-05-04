package employee

type Employee struct {
	Id            int    `json:"id"`
	Address       string `json:"address"`
	BusinessPhone string `json:"bussiness_phone"`
	Company       string `json:"company"`
	EmailAddress  string `json:"email_address"`
	FaxNumber     string `json:"fax_number"`
	FirstName     string `json:"first_name"`
	HomePhone     string `json:"home_phone"`
	JobTitle      string `json:"job_title"`
	LastName      string `json:"last_name"`
	MobilePhone   string `json:"mobile_phone"`
}

type EmployeesList struct {
	Data         []*Employee `json:"data"`
	TotalRecords int         `json:"total_records"`
}

type EmployeeTop struct {
	Id            int    `json:"id"`
	EmailAddress  string `json:"email_address"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	TotalSellings string `json:"total_sellings"`
}
