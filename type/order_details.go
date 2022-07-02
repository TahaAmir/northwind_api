package types

type OrderDetails struct {
	OrderID   int64   `json:"OrderID"`
	ProductID int64   `json:"ProductID"`
	UnitPrice float64 `json:"UnitPrice"`
	Quantity  int64   `json:"Quantity"`
	Discount  float64 `json:"Discount"`
}
