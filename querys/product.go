package querys

import (
	"golang-crud-rest-api/database"
	products "golang-crud-rest-api/type"
)

func CreateProduct(p products.Product) error {
	_, err := database.DB.Exec("INSERT INTO products ( ProductID, ProductName , SupplierID , CategoryID  ,QuantityPerUnit , UnitPrice , UnitsInStock ,UnitsOnOrder  ,ReorderLevel,Discontinued) VALUES (?,?,?,?,?,?,?,?,?,?)  ",
		p.ID, p.Name, p.SupID, p.CatoID, p.QuantityPerUnit, p.Price, p.UnitsInStock, p.UnitsnOrder, p.ReorderLevel, p.Discontinued)
	return err

}

func DeleteProduct(id int64) error {

	_, err := database.DB.Exec("DELETE FROM products WHERE ProductID = ?", id)
	return err
}

func UpdateProduct(p products.Product) error {

	_, err := database.DB.Exec("UPDATE products SET ProductName = ? , SupplierID = ?, CategoryID = ?  ,QuantityPerUnit = ? , UnitPrice = ?, UnitsInStock = ?,UnitsOnOrder =?  ,ReorderLevel= ? , Discontinued=? WHERE ProductID = ? ", p.Name, p.SupID, p.CatoID, p.QuantityPerUnit, p.Price, p.UnitsInStock, p.UnitsnOrder, p.ReorderLevel, p.Discontinued, p.ID)

	return err
}

func GetProduct() ([]products.Product, error) {
	od := []products.Product{}

	rows, err := database.DB.Query("SELECT ProductID , ProductName,Discontinued , SupplierID , CategoryID  ,QuantityPerUnit , UnitPrice , UnitsInStock ,UnitsOnOrder ,ReorderLevel FROM products")
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

	row := database.DB.QueryRow("SELECT ProductID, ProductName , SupplierID ,Discontinued, CategoryID  ,QuantityPerUnit , UnitPrice , UnitsInStock ,UnitsOnOrder  ,ReorderLevel  FROM products WHERE ProductID = ?", id)
	err := row.Scan(&od.ID, &od.Name, &od.SupID, &od.Discontinued, &od.CatoID, &od.QuantityPerUnit, &od.Price, &od.UnitsInStock, &od.UnitsnOrder, &od.ReorderLevel)
	if err != nil {
		return od, err
	}
	return od, nil

}
