package types

import (
	"gopkg.in/guregu/null.v3"
)

type Employee struct {
	EmployeeID      int64       `json:"EmployeeID,string,omitempty"`
	LastName        string      `json:"LastName"`
	FirstName       string      `json:"FirstName"`
	Title           string      `json:"Title"`
	TitleOfCourtesy null.String `json:"TitleOfCourtesy"`
	BirthDate       null.String `json:"BirthDate"`
	Address         null.String `json:"Address"`
	HireDate        null.String `json:"HireDate"`
	City            null.String `json:"City"`
	Region          null.String `json:"Region"`
	PostalCode      null.String `json:"PostalCode"`
	Country         null.String `json:"Country"`
	HomePhone       null.String `json:"HomePhone"`
	Extension       null.String `json:"Extension"`
	Notes           string      `json:"Notes"`
	ReportsTo       null.Int    `json:"ReportsTo"`
}
