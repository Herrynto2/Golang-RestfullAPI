package model

// DataClient ...
type DataClient struct {
	ID             uint64 `json:"_id"`
	CompanyCode    int64  `form:"company_code" json:"company_code"`
	SalesOrg       string `form:"sales_org" json:"sales_org"`
	DistChannel    string `form:"dist_channel" json:"dist_channel"`
	DivisionClient string `form:"division_client" json:"division_client"`
	CustomerNumber string `form:"sap_customer_number" json:"sap_customer_number"`
	CustomerName   string `form:"customer_name" json:"customer_name"`
	Address        string `form:"address" json:"address"`
	Telephone      string `form:"telephone" json:"telephone"`
	MobilePhone    string `form:"mobile_phone" json:"mobile_phone"`
	Fax            string `form:"fax" json:"fax"`
	Email          string `form:"email" json:"email"`
	Status         string `form:"status" json:"status"`
	Attachment     string `form:"attachment" json:"attachment"`
	CreatedAt      string `json:"created_at"`
	CreatedBy      string `json:"created_by"`
	ModifiedAt     string `json:"modified_at"`
	ModifiedBy     string `json:"modified_by"`
	DeleteAt       string `json:"delete_at"`
	DeleteBy       string `json:"delete_by"`
}
