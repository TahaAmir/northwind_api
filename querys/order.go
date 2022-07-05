package querys

import (
	"golang-crud-rest-api/database"
	types "golang-crud-rest-api/type"
)

func CreateOrders(o types.Orders) (err error) {
	_, err = database.DB.Exec(`INSERT INTO orders 
(CustomerID,
EmployeeID,
OrderDate,
RequiredDate,
ShippedDate,
ShipVia,
Freight,
ShipName,
ShipAddress,
ShipCity,
ShipRegion,
ShipPostalCode,
ShipCountry)  VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)  
`, o.CustomerID,
		o.EmployeeID,
		o.OrderDate,
		o.RequiredDate,
		o.ShippedDate,
		o.ShipVia,
		o.Freight,
		o.ShipName,
		o.ShipAddress,
		o.ShipCity,
		o.ShipRegion,
		o.ShipPostalCode,
		o.ShipCountry)
	return err

}

func DeleteOrders(id int64) (err error) {
	_, err = database.DB.Exec("DELETE FROM orders WHERE OrderID = ?", id)
	return err
}

func UpdateOrders(o types.Orders) (err error) {
	_, err = database.DB.Exec(`UPDATE orders SET 
	CustomerID = ?,
	EmployeeID =?,
	OrderDate=?,
	RequiredDate =?,
	ShippedDate =?,
	ShipVia =?,
	Freight =?,
	ShipName =?,
	ShipAddress  =?,
	ShipCity =?,
	ShipRegion =?,
	ShipPostalCode =?,
	ShipCountry =? WHERE OrderID =?`,
		o.CustomerID,
		o.EmployeeID,
		o.OrderDate,
		o.RequiredDate,
		o.ShippedDate,
		o.ShipVia,
		o.Freight,
		o.ShipName,
		o.ShipAddress,
		o.ShipCity,
		o.ShipRegion,
		o.ShipPostalCode,
		o.ShipCountry,
		o.OrderID)
	return err
}

func GetOrders(start, count int) ([]types.Orders, error) {

	if count == 0 {
		count = 10
	}
	orders := []types.Orders{}

	rows, err := database.DB.Query(`SELECT 
	OrderID,
	CustomerID,
	EmployeeID,
	OrderDate,
	RequiredDate,
	ShippedDate,
	ShipVia,
	Freight,
	ShipName,
	ShipAddress,
	ShipCity,
	ShipRegion,
	ShipPostalCode,
	ShipCountry 
	FROM orders LIMIT ? OFFSET ? `, start, count)

	if err != nil {
		return orders, err
	}
	for rows.Next() {
		var o types.Orders

		err = rows.Scan(
			&o.OrderID,
			&o.CustomerID,
			&o.EmployeeID,
			&o.OrderDate,
			&o.RequiredDate,
			&o.ShippedDate,
			&o.ShipVia,
			&o.Freight,
			&o.ShipName,
			&o.ShipAddress,
			&o.ShipCity,
			&o.ShipRegion,
			&o.ShipPostalCode,
			&o.ShipCountry)

		if err != nil {
			return orders, err
		}
		orders = append(orders, o)

	}
	return orders, nil

}

func GetOrdersByID(id int64) (types.Orders, error) {

	var orders types.Orders

	row := database.DB.QueryRow(`SELECT 
	OrderID,
	CustomerID,
	EmployeeID,
	OrderDate,
	RequiredDate,
	ShippedDate,
	ShipVia,
	Freight,
	ShipName,
	ShipAddress,
	ShipCity,
	ShipRegion,
	ShipPostalCode,
	ShipCountry 
	FROM orders WHERE OrderID = ?`, id)

	err := row.Scan(
		&orders.OrderID,
		&orders.CustomerID,
		&orders.EmployeeID,
		&orders.OrderDate,
		&orders.RequiredDate,
		&orders.ShippedDate,
		&orders.ShipVia,
		&orders.Freight,
		&orders.ShipName,
		&orders.ShipAddress,
		&orders.ShipCity,
		&orders.ShipRegion,
		&orders.ShipPostalCode,
		&orders.ShipCountry)
	if err != nil {
		return orders, err
	}
	return orders, nil
}
