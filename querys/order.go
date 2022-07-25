package querys

import (
	"fmt"
	"golang-crud-rest-api/database"
	types "golang-crud-rest-api/type"
)

func CreateOrders(o types.Orders) (types.Orders, error) {
	var order types.Orders
	res, err := database.DB.Exec(`INSERT INTO orders 
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

	if err != nil {
		return order, err
	}
	rowID, err := res.LastInsertId()
	if err != nil {
		return order, err
	}

	order.OrderID = int64(rowID)

	// find  by id
	result, err := GetOrdersByID(order.OrderID)
	if err != nil {
		return order, err
	}
	return result, nil
}

func DeleteOrders(id int64) (err error) {
	r, err := database.DB.Exec("DELETE FROM orders WHERE OrderID = ?", id)
	ar, e := r.RowsAffected()
	var msg string
	if e != nil {
		msg = e.Error()
	}
	if ar == 0 {
		msg += "The Id Entered does not exist"
		err = fmt.Errorf(msg)
	}
	return err
}

func UpdateOrders(o types.Orders) (err error) {
	r, err := database.DB.Exec(`UPDATE orders SET 
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

	if err != nil {
		return err
	}

	ar, e := r.RowsAffected()
	var msg string
	if e != nil {
		msg = e.Error()
	}
	if ar == 0 {
		msg += " Enter Valid Id to update"
		err = fmt.Errorf(msg)
	}
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
