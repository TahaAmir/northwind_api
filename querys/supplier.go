package querys

import (
	"golang-crud-rest-api/database"
	types "golang-crud-rest-api/type"
)

func CreateSuppliers(s types.Supplier) (err error) {

	_, err = database.DB.Exec(`INSERT INTO suppliers  
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

	return err
}

func DeleteSupplier(id int64) (err error) {

	_, err = database.DB.Exec("DELETE  FROM suppliers WHERE SupplierID = ? ", id)
	return err
}

func UpdateSuppliers(s types.Supplier) (err error) {

	_, err = database.DB.Exec(`UPDATE suppliers SET 
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
	return err
}

func GetSupplier() ([]types.Supplier, error) {

	supplier := []types.Supplier{}

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
	HomePage FROM suppliers`)

	if err != nil {
		return supplier, err
	}
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
