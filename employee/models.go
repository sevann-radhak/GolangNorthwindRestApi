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
	Id           int    `json:"id"`
	EmailAddress string `json:"emailAddress"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	TotalSales   string `json:"totalSales"`
}
