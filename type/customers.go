package types

import "database/sql"

type Customers struct {
	ID             int64          `json:"CustomerID,string,omitempty"`
	CompanyName    string         `json:"CompanyName"`
	ContactName    string         `json:"ContactName"`
	ContactTitle   string         `json:"ContactTiltle"`
	City           sql.NullString `json:"City"`
	Address        sql.NullString `json:"Address"`
	Region         sql.NullString `json:"Region"`
	PostalCode     sql.NullString `json:"PostalCode"`
	Country        sql.NullString `json:"Country"`
	Phone          sql.NullString `json:"Phone"`
	Fax            sql.NullString `json:"Fax"`
	Image          any            `json:"Image"`
	ImageThumbnail any            `json:"ImageThumbnail"`
}
