package querys

import (
	"golang-crud-rest-api/database"
	types "golang-crud-rest-api/type"
)

func CreateOrderDetails(od types.OrderDetails) (err error) {
	_, err = database.DB.Exec(`INSERT INTO order_details 
	(OrderID,
	ProductID, 
	UnitPrice, 
	Quantity,  
	Discount) VALUES (?,?,?,?,?) 
	`, od.OrderID,
		od.ProductID,
		od.UnitPrice,
		od.Quantity,
		od.Discount)

	return err
}

func DeleteOrderDetails(id int64) (err error) {

	_, err = database.DB.Exec("DELETE FROM order_details WHERE OrderID=?", id)
	return err
}

func UpdateOrderDetails(od types.OrderDetails) (err error) {
	_, err = database.DB.Exec(`UPDATE order_details SET
	ProductID=?,
	UnitPrice=?,
	Quantity=?,
	Discount=? WHERE 
	OrderID=?
	`, od.ProductID,
		od.UnitPrice,
		od.Quantity,
		od.Discount,
		od.OrderID)

	return err
}

func GetOrderDetails() ([]types.OrderDetails, error) {
	orderdetails := []types.OrderDetails{}

	rows, err := database.DB.Query(`SELECT  
	OrderID,
	ProductID,
	UnitPrice,
	Quantity,
	Discount FROM order_details`)
	if err != nil {
		return orderdetails, err
	}

	for rows.Next() {
		var od types.OrderDetails
		err = rows.Scan(
			&od.OrderID,
			&od.ProductID,
			&od.UnitPrice,
			&od.Quantity,
			&od.Discount)
		if err != nil {
			return orderdetails, err
		}
		orderdetails = append(orderdetails, od)
	}
	return orderdetails, nil

}

func GetOrderDetailsById(id int64) (types.OrderDetails, error) {

	var od types.OrderDetails

	row := database.DB.QueryRow(`SELECT  
	OrderID,
	ProductID,
	UnitPrice,
	Quantity,
	Discount FROM order_details WHERE OrderID = ? `, id)
	err := row.Scan(
		&od.OrderID,
		&od.ProductID,
		&od.UnitPrice,
		&od.Quantity,
		&od.Discount)
	if err != nil {
		return od, err
	}
	return od, err
}
