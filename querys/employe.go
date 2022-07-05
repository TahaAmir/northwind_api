package querys

import (
	"golang-crud-rest-api/database"
	types "golang-crud-rest-api/type"
)

func CreateEmployee(e types.Employee) (err error) {
	_, err = database.DB.Exec(`INSERT INTO employees 
   ( LastName
	,FirstName
	,Title
	,TitleOfCourtesy
	,BirthDate
	,HireDate
	,Address
	,City
	,Region
	,PostalCode
	,Country
	,HomePhone
	,Extension
	,Photo
	,Notes
	,ReportsTo ) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
		e.LastName,
		e.FirstName,
		e.Title,
		e.TitleOfCourtesy,
		e.BirthDate,
		e.HireDate,
		e.Address,
		e.City,
		e.Region,
		e.PostalCode,
		e.Country,
		e.HomePhone,
		e.Extension,
		e.Photo,
		e.Notes,
		e.ReportsTo)
	return
}

func DeleteEmployee(id int64) (err error) {
	_, err = database.DB.Exec("DELETE FROM employees WHERE EmployeeID = ?", id)

	return
}

func UpdateEmployee(e types.Employee) (err error) {

	_, err = database.DB.Exec(`UPDATE employees SET 
     LastName =?
	,FirstName=?
	,Title=?
	,TitleOfCourtesy=?
	,BirthDate=?
	,HireDate=?
	,Address=?
	,City=?
	,Region=?
	,PostalCode=?
	,Country=?
	,HomePhone=?
	,Extension=?
	,Photo=?
	,Notes=?
	,ReportsTo=?
	WHERE EmployeeID=?`,
		e.LastName,
		e.FirstName,
		e.Title,
		e.TitleOfCourtesy,
		e.BirthDate,
		e.HireDate,
		e.Address,
		e.City,
		e.Region,
		e.PostalCode,
		e.Country,
		e.HomePhone,
		e.Extension,
		e.Photo,
		e.Notes,
		e.ReportsTo,
		e.EmployeeID)

	return
}

func GetEmployee(start, count int) ([]types.Employee, error) {

	if count == 0 {
		count = 10
	}
	employee := []types.Employee{}
	rows, err := database.DB.Query(`SELECT 
	 EmployeeID
	,LastName
	,FirstName
	,Title
	,TitleOfCourtesy
	,BirthDate
	,HireDate
	,Address
	,City
	,Region
	,PostalCode
	,Country
	,HomePhone
	,Extension
	,Photo
	,Notes
	,ReportsTo FROM employees  LIMIT ? OFFSET ?
	`, count, start)
	if err != nil {
		return employee, err
	}
	for rows.Next() {
		var e types.Employee
		err = rows.Scan(
			&e.EmployeeID,
			&e.LastName,
			&e.FirstName,
			&e.Title,
			&e.TitleOfCourtesy,
			&e.BirthDate,
			&e.HireDate,
			&e.Address,
			&e.City,
			&e.Region,
			&e.PostalCode,
			&e.Country,
			&e.HomePhone,
			&e.Extension,
			&e.Photo,
			&e.Notes,
			&e.ReportsTo)
		if err != nil {
			return employee, err
		}
		employee = append(employee, e)
	}
	return employee, nil

}

func GetEmployeeByID(id int64) (types.Employee, error) {

	var e types.Employee

	row := database.DB.QueryRow(`SELECT 
	EmployeeID
   ,LastName
   ,FirstName
   ,Title
   ,TitleOfCourtesy
   ,BirthDate
   ,HireDate
   ,Address
   ,City
   ,Region
   ,PostalCode
   ,Country
   ,HomePhone
   ,Extension
   ,Photo
   ,Notes
   ,ReportsTo FROM employees WHERE EmployeeID=?`, id)

	err := row.Scan(
		&e.EmployeeID,
		&e.LastName,
		&e.FirstName,
		&e.Title,
		&e.TitleOfCourtesy,
		&e.BirthDate,
		&e.HireDate,
		&e.Address,
		&e.City,
		&e.Region,
		&e.PostalCode,
		&e.Country,
		&e.HomePhone,
		&e.Extension,
		&e.Photo,
		&e.Notes,
		&e.ReportsTo)
	if err != nil {
		return e, err
	}
	return e, nil
}
