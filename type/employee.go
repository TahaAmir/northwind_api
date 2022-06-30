package types

type Employee struct {
	EmployeeID      int64  `json:"EmployeeID,string,omitempty"`
	LastName        string `json:"LastName"`
	FirstName       string `json:"FirstName"`
	Title           string `json:"Title"`
	TitleOfCourtesy string `json:"TitleOfCourtesy"`
	BirthDate       string `json:"BirthDate"`
	HireDate        string `json:"HireDate"`
	Address         string `json:"Address"`
	City            string `json:"City"`
	Region          string `json:"Region"`
	PostalCode      string `json:"PostalCode"`
	Country         string `json:"Country"`
	HomePhone       string `json:"HomePhone"`
	Extension       string `json:"Extension"`
	Photo           any    `json:"Photo"`
	Notes           string `json:"Notes"`
	ReportsTo       int    `json:"ReportsTo"`
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
