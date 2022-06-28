package main

func CreateProduct(od Product) error {
	_, err := DB.Exec("INSERT INTO products (  ProductName , SupplierID , CategoryID  ,QuantityPerUnit , UnitPrice , UnitsInStock ,UnitsOnOrder  ,ReorderLevel,Discontinued  ",
		od.Name, od.SupID, od.CatoID, od.QuantityPerUnit, od.Price, od.UnitsInStock, od.UnitsnOrder, od.ReorderLevel, od.Discontinued)
	return err

}

func DeleteProduct(id int64) error {

	_, err := DB.Exec("DELETE FROM products WHERE ProductID = ?", id)
	return err
}

func UpdateProduct(od Product) error {
	// bd, err := getDB()
	// if err != nil {
	// 	return err
	// }
	_, err := DB.Exec("UPDATE products SET ProductName = ? , SupplierID = ?, CategoryID = ?  ,QuantityPerUnit = ? , UnitPrice = ?, UnitsInStock = ?,UnitsOnOrder =?  ,ReorderLevel= ? , Discontinued=? WHERE ProductID = ? ", od.Name, od.SupID, od.CatoID, od.QuantityPerUnit, od.Price, od.UnitsInStock, od.UnitsnOrder, od.ReorderLevel, od.Discontinued, od.ID)

	return err
}

func GetProduct() ([]Product, error) {
	od := []Product{}

	rows, err := DB.Query("SELECT ProductID , ProductName,Discontinued , SupplierID , CategoryID  ,QuantityPerUnit , UnitPrice , UnitsInStock ,UnitsOnOrder ,ReorderLevel FROM products")
	if err != nil {
		return od, err
	}
	for rows.Next() {
		var order Product
		err = rows.Scan(&order.ID, &order.Name, &order.Discontinued, &order.SupID, &order.CatoID, &order.QuantityPerUnit, &order.Price, &order.UnitsInStock, &order.UnitsnOrder, &order.ReorderLevel)
		if err != nil {
			return od, err
		}
		od = append(od, order)
	}
	return od, nil
}

func GetProductById(id int64) (Product, error) {
	var od Product

	row := DB.QueryRow("SELECT ProductName , SupplierID ,Discontinued, CategoryID  ,QuantityPerUnit , UnitPrice , UnitsInStock ,UnitsOnOrder  ,ReorderLevel  FROM products WHERE ProductID = ?", id)
	err := row.Scan(&od.Name, &od.SupID, &od.Discontinued, &od.CatoID, &od.QuantityPerUnit, &od.Price, &od.UnitsInStock, &od.UnitsnOrder, &od.ReorderLevel)
	if err != nil {
		return od, err
	}
	return od, nil

}
