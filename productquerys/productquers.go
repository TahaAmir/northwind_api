package productquerys

import (
	"golang-crud-rest-api/databaseconnection"
	"golang-crud-rest-api/products"
)

var product *products.Product

func CreateProduct(p products.Product) error {
	_, err := databaseconnection.DB.Exec("INSERT INTO products (  ProductName , SupplierID , CategoryID  ,QuantityPerUnit , UnitPrice , UnitsInStock ,UnitsOnOrder  ,ReorderLevel,Discontinued  ",
		p.Name, p.SupID, p.CatoID, p.QuantityPerUnit, p.Price, p.UnitsInStock, p.UnitsnOrder, p.ReorderLevel, p.Discontinued)
	return err

}

func DeleteProduct(id int64) error {

	_, err := databaseconnection.DB.Exec("DELETE FROM products WHERE ProductID = ?", id)
	return err
}

func UpdateProduct(p products.Product) error {

	_, err := databaseconnection.DB.Exec("UPDATE products SET ProductName = ? , SupplierID = ?, CategoryID = ?  ,QuantityPerUnit = ? , UnitPrice = ?, UnitsInStock = ?,UnitsOnOrder =?  ,ReorderLevel= ? , Discontinued=? WHERE ProductID = ? ", p.Name, p.SupID, p.CatoID, p.QuantityPerUnit, p.Price, p.UnitsInStock, p.UnitsnOrder, p.ReorderLevel, p.Discontinued, p.ID)

	return err
}

func GetProduct() ([]products.Product, error) {
	od := []products.Product{}

	rows, err := databaseconnection.DB.Query("SELECT ProductID , ProductName,Discontinued , SupplierID , CategoryID  ,QuantityPerUnit , UnitPrice , UnitsInStock ,UnitsOnOrder ,ReorderLevel FROM products")
	if err != nil {
		return od, err
	}
	for rows.Next() {
		var p products.Product
		err = rows.Scan(&p.ID, &p.Name, &p.Discontinued, &p.SupID, &p.CatoID, &p.QuantityPerUnit, &p.Price, &p.UnitsInStock, &p.UnitsnOrder, &p.ReorderLevel)
		if err != nil {
			return od, err
		}
		od = append(od, p)
	}
	return od, nil
}

func GetProductById(id int64) (products.Product, error) {
	var od products.Product

	row := databaseconnection.DB.QueryRow("SELECT ProductName , SupplierID ,Discontinued, CategoryID  ,QuantityPerUnit , UnitPrice , UnitsInStock ,UnitsOnOrder  ,ReorderLevel  FROM products WHERE ProductID = ?", id)
	err := row.Scan(&od.Name, &od.SupID, &od.Discontinued, &od.CatoID, &od.QuantityPerUnit, &od.Price, &od.UnitsInStock, &od.UnitsnOrder, &od.ReorderLevel)
	if err != nil {
		return od, err
	}
	return od, nil

}
