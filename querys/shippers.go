package querys

import (
	"fmt"
	"golang-crud-rest-api/database"
	types "golang-crud-rest-api/type"
)

func CreateShippers(s types.Shippers) (types.Shippers, error) {
	var shipper types.Shippers
	res, err := database.DB.Exec(`INSERT INTO shippers 
	(CompanyName,
	Phone) VALUES (?,?)
	`, s.CompanyName,
		s.Phone)
	if err != nil {
		return shipper, err
	}
	rowID, err := res.LastInsertId()
	if err != nil {
		return shipper, err
	}

	shipper.ShipperID = int64(rowID)

	// find  by id
	result, err := GetShippersById(shipper.ShipperID)
	if err != nil {
		return shipper, err
	}
	return result, nil
}

func DeleteShippers(id int64) (err error) {
	r, err := database.DB.Exec("DELETE FROM shippers WHERE ShipperID = ? ", id)

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

func UpdateShippers(s types.Shippers) (err error) {
	r, err := database.DB.Exec(`UPDATE shippers SET  CompanyName = ? , Phone= ? WHERE ShipperID = ?`, s.CompanyName, s.Phone, s.ShipperID)

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

func GetShippers() ([]types.Shippers, error) {
	shipper := []types.Shippers{}

	rows, err := database.DB.Query(`SELECT  
	ShipperID,
	CompanyName,
	Phone FROM shippers`)
	if err != nil {
		return shipper, err
	}

	for rows.Next() {
		var s types.Shippers
		err = rows.Scan(
			&s.ShipperID,
			&s.CompanyName,
			&s.Phone)
		if err != nil {
			return shipper, err
		}
		shipper = append(shipper, s)
	}
	return shipper, nil

}

func GetShippersById(id int64) (types.Shippers, error) {

	var s types.Shippers

	row := database.DB.QueryRow(`
	SELECT  ShipperID,
	CompanyName,
	Phone FROM shippers WHERE ShipperID = ? `, id)
	err := row.Scan(
		&s.ShipperID,
		&s.CompanyName,
		&s.Phone)
	if err != nil {
		return s, err
	}
	return s, err
}
