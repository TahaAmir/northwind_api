package querys

import (
	"fmt"
	"golang-crud-rest-api/database"

	types "golang-crud-rest-api/type"
)

func CreateProduct(p types.Product) (types.Product, error) {
	var product types.Product
	res, err := database.DB.Exec(`INSERT INTO products 
	(  ProductName ,
	   SupplierID , 
	   CategoryID  ,
	   QuantityPerUnit , 
	   UnitPrice ,
	   UnitsInStock ,
	   UnitsOnOrder  ,
	   ReorderLevel,
	   Discontinued) VALUES (?,?,?,?,?,?,?,?,?)  `,
		p.Name,
		p.SupID,
		p.CatoID,
		p.QuantityPerUnit,
		p.Price,
		p.UnitsInStock,
		p.UnitsnOrder,
		p.ReorderLevel,
		p.Discontinued)
	if err != nil {
		return product, err
	}
	rowID, err := res.LastInsertId()
	if err != nil {
		return product, err
	}

	product.ID = int64(rowID)

	// find  by id
	result, err := GetProductById(product.ID)
	if err != nil {
		return product, err
	}
	return result, nil

}

func DeleteProduct(id int64) error {

	r, err := database.DB.Exec("DELETE FROM products WHERE ProductID = ?", id)
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

func UpdateProduct(p types.Product) error {

	r, err := database.DB.Exec(`UPDATE products SET 
	ProductName = ? , 
	SupplierID = ?, 
	CategoryID = ?  ,
	QuantityPerUnit = ? , 
	UnitPrice = ?, 
	UnitsInStock = ?,
	UnitsOnOrder =?  ,
	ReorderLevel= ? , 
	Discontinued=? WHERE ProductID = ? `,
		p.Name,
		p.SupID,
		p.CatoID,
		p.QuantityPerUnit,
		p.Price,
		p.UnitsInStock,
		p.UnitsnOrder,
		p.ReorderLevel,
		p.Discontinued,
		p.ID)
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

func GetProduct(start, count int) ([]types.Product, error) {

	if count == 0 {
		count = 10
	}
	od := []types.Product{}

	rows, err := database.DB.Query(`SELECT 
	ProductID , 
	ProductName,
	Discontinued , 
	SupplierID , 
	CategoryID  ,
	QuantityPerUnit , 
	UnitPrice , 
	UnitsInStock ,
	UnitsOnOrder ,
	ReorderLevel FROM products LIMIT ? OFFSET ? `, count, start)
	if err != nil {
		return od, err
	}
	for rows.Next() {
		var p types.Product
		err = rows.Scan(
			&p.ID,
			&p.Name,
			&p.Discontinued,
			&p.SupID,
			&p.CatoID,
			&p.QuantityPerUnit,
			&p.Price,
			&p.UnitsInStock,
			&p.UnitsnOrder,
			&p.ReorderLevel)
		if err != nil {
			return od, err
		}
		od = append(od, p)
	}
	return od, nil
}

func GetProductById(id int64) (types.Product, error) {
	var od types.Product

	row := database.DB.QueryRow(`SELECT 
	ProductID, 
	ProductName , 
	SupplierID ,
	Discontinued, 
	CategoryID  ,
	QuantityPerUnit , 
	UnitPrice , 
	UnitsInStock ,
	UnitsOnOrder  ,
	ReorderLevel  FROM products WHERE ProductID = ?`, id)
	err := row.Scan(
		&od.ID,
		&od.Name,
		&od.SupID,
		&od.Discontinued,
		&od.CatoID,
		&od.QuantityPerUnit,
		&od.Price,
		&od.UnitsInStock,
		&od.UnitsnOrder,
		&od.ReorderLevel)
	if err != nil {
		return od, err
	}
	return od, nil

}
