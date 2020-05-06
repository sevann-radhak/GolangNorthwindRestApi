package customer

type Customer struct {
	Id            int    `json:"id"`
	Address       string `json:"address"`
	BusinessPhone string `json:"bussiness_phone"`
	City          string `json:"city"`
	Company       string `json:"company"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
}

type CustomersList struct {
	Data         []*Customer `json:"data"`
	TotalRecords int         `json:"total_records"`
}
