package types

import (
	"gopkg.in/guregu/null.v3"
)

type Supplier struct {
	SupplierID   int64       `json:"SupplierID"`
	CompanyName  string      `json:"CompanyName"`
	ContactName  string      `json:"ContactName"`
	ContactTitle string      `json:"ContactTitle"`
	Address      string      `json:"Address"`
	City         string      `json:"City"`
	Region       null.String `json:"Region"`
	PostalCode   string      `json:"PostalCode"`
	Country      string      `json:"Country"`
	Phone        string      `json:"Phone"`
	Fax          null.String `json:"Fax"`
	HomePage     null.String `json:"HomePage"`
}
