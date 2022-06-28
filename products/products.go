package products

type Product struct {
	ID              int     `json:"ProductID"`
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
