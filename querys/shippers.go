package querys

import (
	"golang-crud-rest-api/database"
	types "golang-crud-rest-api/type"
)

func CreateShippers(s types.Shippers) (err error) {
	_, err = database.DB.Exec(`INSERT INTO shippers 
	(CompanyName,
	Phone) VALUES (?,?)
	`, s.CompanyName,
		s.Phone)

	return err
}

func DeleteShippers(id int64) (err error) {
	_, err = database.DB.Exec("DELETE FROM shippers WHERE ShipperID = ? ", id)
	return err
}

func UpdateShippers(s types.Shippers) (err error) {
	_, err = database.DB.Exec(`UPDATE shippers SET  CompanyName = ? , Phone= ? WHERE ShipperID = ?`, s.CompanyName, s.Phone, s.ShipperID)
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
