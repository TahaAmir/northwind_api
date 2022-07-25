package querys

import (
	"fmt"
	"golang-crud-rest-api/database"
	types "golang-crud-rest-api/type"
)

func CreateSuppliers(s types.Supplier) (types.Supplier, error) {
	var suupplier types.Supplier
	res, err := database.DB.Exec(`INSERT INTO suppliers  
  ( CompanyName,
	ContactName,
	ContactTitle,
	Address,
	City,
	Region,
	PostalCode,
	Country,
	Phone,
	Fax,
	HomePage ) VALUES (?,?,?,?,?,?,?,?,?,?,?)
	`, s.CompanyName,
		s.ContactName,
		s.ContactTitle,
		s.Address,
		s.City,
		s.Region,
		s.PostalCode,
		s.Country,
		s.Phone,
		s.Fax,
		s.HomePage)

	if err != nil {
		return suupplier, err
	}
	rowID, err := res.LastInsertId()
	if err != nil {
		return suupplier, err
	}

	suupplier.SupplierID = int64(rowID)

	// find  by id
	result, err := GetSupplierByID(suupplier.SupplierID)
	if err != nil {
		return suupplier, err
	}
	return result, nil
}

func DeleteSupplier(id int64) (err error) {

	r, err := database.DB.Exec("DELETE  FROM suppliers WHERE SupplierID = ? ", id)
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

func UpdateSuppliers(s types.Supplier) (err error) {

	r, err := database.DB.Exec(`UPDATE suppliers SET 
	CompanyName = ?,
	ContactName  = ?,
	ContactTitle = ?,
	Address = ?,
	City = ?,
	Region = ?,
	PostalCode =?,
	Country =?,
	Phone = ?,
	Fax = ?,
	HomePage = ?  WHERE  SupplierID = ?
    `, s.CompanyName,
		s.ContactName,
		s.ContactTitle,
		s.Address,
		s.City,
		s.Region,
		s.PostalCode,
		s.Country,
		s.Phone,
		s.Fax,
		s.HomePage,
		s.SupplierID)

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

func GetSupplier(start, count int) ([]types.Supplier, error) {

	supplier := []types.Supplier{}

	if count == 0 {
		count = 10
	}
	rows, err := database.DB.Query(`SELECT 
	SupplierID,
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
	HomePage FROM suppliers LIMIT ? OFFSET ?`, count, start)

	if err != nil {
		return supplier, err
	}

	defer rows.Close()

	for rows.Next() {
		var s types.Supplier

		err = rows.Scan(
			&s.SupplierID,
			&s.CompanyName,
			&s.ContactName,
			&s.ContactTitle,
			&s.Address,
			&s.City,
			&s.Region,
			&s.PostalCode,
			&s.Country,
			&s.Phone,
			&s.Fax,
			&s.HomePage)

		if err != nil {
			return supplier, err
		}
		supplier = append(supplier, s)

	}
	return supplier, nil

}

func GetSupplierByID(id int64) (types.Supplier, error) {

	var s types.Supplier

	row := database.DB.QueryRow(`SELECT 
	SupplierID,
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
	HomePage
	FROM suppliers WHERE SupplierID = ?`, id)

	err := row.Scan(
		&s.SupplierID,
		&s.CompanyName,
		&s.ContactName,
		&s.ContactTitle,
		&s.Address,
		&s.City,
		&s.Region,
		&s.PostalCode,
		&s.Country,
		&s.Phone,
		&s.Fax,
		&s.HomePage)
	if err != nil {
		return s, err
	}
	return s, nil
}
