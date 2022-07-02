package types

import "database/sql"

type Employee struct {
	EmployeeID      int64          `json:"EmployeeID,string,omitempty"`
	LastName        string         `json:"LastName"`
	FirstName       string         `json:"FirstName"`
	Title           string         `json:"Title"`
	TitleOfCourtesy sql.NullString `json:"TitleOfCourtesy"`
	BirthDate       sql.NullString `json:"BirthDate"`
	Address         sql.NullString `json:"Address"`
	HireDate        sql.NullString `json:"HireDate"`
	City            sql.NullString `json:"City"`
	Region          sql.NullString `json:"Region"`
	PostalCode      sql.NullString `json:"PostalCode"`
	Country         sql.NullString `json:"Country"`
	HomePhone       sql.NullString `json:"HomePhone"`
	Extension       sql.NullString `json:"Extension"`
	Photo           any            `json:"Photo"`
	Notes           string         `json:"Notes"`
	ReportsTo       sql.NullInt64  `json:"ReportsTo"`
}

// EmployeeID
// LastName
// FirstName
// Title
// TitleOfCourtesy
// BirthDate
// HireDate
// Address
// City
// Region
// PostalCode
// Country
// HomePhone
// Extension
// Photo
// Notes
// ReportsTo
