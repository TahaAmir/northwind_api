package querys

import (
	"golang-crud-rest-api/database"
	types "golang-crud-rest-api/type"
)

func CreateCustomer(c types.Customers) error {

	_, err := database.DB.Exec("INSERT INTO customers (CompanyName,ContactName,ContactTitle,Address,City,Region,PostalCode,Country,Phone,Fax,Image, ImageThumbnail)VALUES (?,?,?,?,?,?,?,?,?,?,?,?)",
		c.CompanyName, c.ContactName, c.ContactTitle, c.Address, c.City, c.Region, c.PostalCode, c.Country, c.Phone, c.Fax, c.Image, c.ImageThumbnail)
	return err
}

func DeleteCustomer(id int64) error {

	_, err := database.DB.Exec("DELETE FROM customers where CustomerID=?", id)
	return err
}

func UpdateCustomer(c types.Customers) error {

	_, err := database.DB.Exec("UPDATE customers SET CompanyName =?,ContactName=?,ContactTitle=?,Address=?,City=?,Region=?,PostalCode=?,Country=?,Phone=?,Fax=?,Image=?, ImageThumbnail=? WHERE CustomerID= ?",
		c.CompanyName, c.ContactName, c.ContactTitle, c.Address, c.City, c.Region, c.PostalCode, c.Country, c.Phone, c.Fax, c.Image, c.ImageThumbnail, c.ID)
	return err
}

func GetCustomer() ([]types.Customers, error) {

	customer := []types.Customers{}
	rows, err := database.DB.Query("SELECT CustomerID,CompanyName,ContactName,ContactTitle,Address,City,Region,PostalCode,Country,Phone,Fax,Image, ImageThumbnail FROM customers ")
	if err != nil {
		return customer, err
	}

	for rows.Next() {
		var c types.Customers
		err = rows.Scan(&c.ID, &c.CompanyName, &c.ContactName, &c.ContactTitle, &c.Address, &c.City, &c.Region, &c.PostalCode, &c.Country, &c.Phone, &c.Fax, &c.Image, &c.ImageThumbnail)
		if err != nil {
			return customer, err
		}
		customer = append(customer, c)
	}
	return customer, nil

}

func GetCustomerById(id int64) (types.Customers, error) {
	var c types.Customers

	rows := database.DB.QueryRow("SELECT CustomerID,CompanyName,ContactName,ContactTitle,Address,City,Region,PostalCode,Country,Phone,Fax,Image, ImageThumbnail FROM customers WHERE CustomerID =? ", id)
	err := rows.Scan(&c.ID, &c.CompanyName, &c.ContactName, &c.ContactTitle, &c.Address, &c.City, &c.Region, &c.PostalCode, &c.Country, &c.Phone, &c.Fax, &c.Image, &c.ImageThumbnail)

	if err != nil {
		return c, err
	}
	return c, nil
}
