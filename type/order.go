package types

import "gopkg.in/guregu/null.v3"

type Orders struct {
	OrderID        int64       `json:"OrderID"`
	CustomerID     int64       `json:"CustomerID"`
	EmployeeID     int64       `json:"EmployeeID"`
	OrderDate      string      `json:"OrderDate"`
	RequiredDate   string      `json:"RequiredDate"`
	ShippedDate    null.String `json:"ShippedDate"`
	ShipVia        int         `json:"ShipVia"`
	Freight        float64     `json:"Freight"`
	ShipName       string      `json:"ShipName"`
	ShipAddress    string      `json:"ShipAddress"`
	ShipCity       string      `json:"ShipCity"`
	ShipRegion     null.String `json:"ShipRegion"`
	ShipPostalCode null.String `json:"ShipPostalCode"`
	ShipCountry    null.String `json:"ShipCountry"`
}
