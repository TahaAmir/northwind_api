package types

import "database/sql"

type Supplier struct {
	SupplierID   int64          `json:"SupplierID"`
	CompanyName  string         `json:"CompanyName"`
	ContactName  string         `json:"ContactName"`
	ContactTitle string         `json:"ContactTitle"`
	Address      string         `json:"Address"`
	City         string         `json:"City"`
	Region       sql.NullString `json:"Region"`
	PostalCode   string         `json:"PostalCode"`
	Country      string         `json:"Country"`
	Phone        string         `json:"Phone"`
	Fax          sql.NullString `json:"Fax"`
	HomePage     sql.NullString `json:"HomePage"`
}
