package querys

import (
	"fmt"
	"golang-crud-rest-api/database"
	types "golang-crud-rest-api/type"
)

func CreateCustomer(c types.Customers) (types.Customers, error) {
	var customer types.Customers
	res, err := database.DB.Exec(`INSERT INTO customers 
	(CompanyName,
	ContactName,
	ContactTitle,
	Address,
	City,
	Region,
	PostalCode,
	Country,
	Phone,
	Fax,
	Image, 
	ImageThumbnail) 
	VALUES (?,?,?,?,?,?,?,?,?,?,?,?) `,
		c.CompanyName,
		c.ContactName,
		c.ContactTitle,
		c.Address,
		c.City,
		c.Region,
		c.PostalCode,
		c.Country,
		c.Phone,
		c.Fax,
		c.Image,
		c.ImageThumbnail)

	if err != nil {
		return customer, err
	}
	rowID, err := res.LastInsertId()
	if err != nil {
		return customer, err
	}

	customer.ID = int64(rowID)

	// find  by id
	result, err := GetCustomerById(customer.ID)
	if err != nil {
		return customer, err
	}

	return result, nil
}

func DeleteCustomer(id int64) error {

	r, err := database.DB.Exec("DELETE FROM customers where CustomerID=?", id)

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

func UpdateCustomer(c types.Customers) error {

	r, err := database.DB.Exec(`UPDATE customers SET 
	CompanyName =?,
	ContactName=?,
	ContactTitle=?,
	Address=?,
	City=?,
	Region=?,
	PostalCode=?,
	Country=?,
	Phone=?,
	Fax=?,
	Image=?,
	ImageThumbnail=? WHERE CustomerID= ?`,
		c.CompanyName,
		c.ContactName,
		c.ContactTitle,
		c.Address,
		c.City,
		c.Region,
		c.PostalCode,
		c.Country,
		c.Phone,
		c.Fax,
		c.Image,
		c.ImageThumbnail,
		c.ID)
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

func GetCustomer(start, count int) ([]types.Customers, error) {
	if count == 0 {
		count = 10
	}

	customer := []types.Customers{}
	rows, err := database.DB.Query(`SELECT 
	CustomerID,
	CompanyName,
	ContactName,
	ContactTitle,
	Address,City,
	Region,
	PostalCode,
	Country,
	Phone,
	Fax,
	Image, 
	ImageThumbnail FROM customers  LIMIT ? OFFSET ?`, count, start)
	if err != nil {
		return customer, err
	}

	for rows.Next() {
		var c types.Customers
		err = rows.Scan(&c.ID,
			&c.CompanyName,
			&c.ContactName,
			&c.ContactTitle,
			&c.Address,
			&c.City,
			&c.Region,
			&c.PostalCode,
			&c.Country,
			&c.Phone,
			&c.Fax,
			&c.Image,
			&c.ImageThumbnail)
		if err != nil {
			return customer, err
		}
		customer = append(customer, c)
	}
	return customer, nil

}

func GetCustomerById(id int64) (types.Customers, error) {
	var c types.Customers

	rows := database.DB.QueryRow(`SELECT 
	CustomerID,
	CompanyName,
	ContactName,
	ContactTitle,
	Address,
	City,
	Region,
	PostalCode,
	Country,
	Phone,
	Fax,
	Image, 
	ImageThumbnail FROM customers WHERE CustomerID =? `, id)
	err := rows.Scan(&c.ID,
		&c.CompanyName,
		&c.ContactName,
		&c.ContactTitle,
		&c.Address,
		&c.City,
		&c.Region,
		&c.PostalCode,
		&c.Country,
		&c.Phone,
		&c.Fax,
		&c.Image,
		&c.ImageThumbnail)

	if err != nil {
		return c, err
	}
	return c, nil
}
