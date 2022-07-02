package types

import "database/sql"

type Orders struct {
	OrderID        int64          `json:"OrderID"`
	CustomerID     int64          `json:"CustomerID"`
	EmployeeID     int64          `json:"EmployeeID"`
	OrderDate      string         `json:"OrderDate"`
	RequiredDate   string         `json:"RequiredDate"`
	ShippedDate    sql.NullString `json:"ShippedDate"`
	ShipVia        int            `json:"ShipVia"`
	Freight        float64        `json:"Freight"`
	ShipName       string         `json:"ShipName"`
	ShipAddress    string         `json:"ShipAddress"`
	ShipCity       string         `json:"ShipCity"`
	ShipRegion     sql.NullString `json:"ShipRegion"`
	ShipPostalCode sql.NullString `json:"ShipPostalCode"`
	ShipCountry    sql.NullString `json:"ShipCountry"`
}
