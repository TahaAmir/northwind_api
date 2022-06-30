package types

import "database/sql"

type NullString sql.NullString

type Customers struct {
	ID             int64      `json:"CustomerID,string,omitempty"`
	CompanyName    string     `json:"CompanyName"`
	ContactName    string     `json:"ContactName"`
	ContactTitle   string     `json:"ContactTiltle"`
	Address        string     `json:"Address"`
	City           string     `json:"City"`
	Region         NullString `json:"Region"`
	PostalCode     string     `json:"PostalCode"`
	Country        string     `json:"Country"`
	Phone          string     `json:"Phone"`
	Fax            string     `json:"Fax"`
	Image          any        `json:"Image"`
	ImageThumbnail any        `json:"ImageThumbnail"`
}
