package types

type Product struct {
	ID              int64   `json:"ProductID,string,omitempty"`
	Name            string  `json:"ProductName"`
	SupID           int     `json:"SupplierID"`
	CatoID          int     `json:"CategoryID"`
	Price           float64 `json:"UnitPrice"`
	UnitsInStock    int     `json:"UnitsInStockk"`
	UnitsnOrder     int     `json:"UnitsOnOrder"`
	ReorderLevel    int     `json:"ReorderLevel"`
	Discontinued    int     `json:"Discontinued"`
	QuantityPerUnit string  `json:"QuantityPerUnit "`
}

// ProductID , ProductName , SupplierID , CategoryID  ,QuantityPerUnit , UnitPrice , UnitsInStock ,UnitsOnOrder  ,ReorderLevel
// Discontinued
