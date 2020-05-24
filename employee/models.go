package employee

type Employee struct {
	Id            int    `json:"id"`
	Address       string `json:"address"`
	BusinessPhone string `json:"bussinessPhone"`
	Company       string `json:"company"`
	EmailAddress  string `json:"emailAddress"`
	FaxNumber     string `json:"faxNumber"`
	FirstName     string `json:"firstName"`
	HomePhone     string `json:"homePhone"`
	JobTitle      string `json:"jobTitle"`
	LastName      string `json:"lastName"`
	MobilePhone   string `json:"mobilePhone"`
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
