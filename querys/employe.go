package querys

import (
	"fmt"
	"golang-crud-rest-api/database"
	types "golang-crud-rest-api/type"
)

func CreateEmployee(e types.Employee) (types.Employee, error) {
	var employee types.Employee
	res, err := database.DB.Exec(`INSERT INTO employees 
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
	,Notes
	,ReportsTo ) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
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
		e.Notes,
		e.ReportsTo)
	if err != nil {
		return employee, err
	}

	rowID, err := res.LastInsertId()
	if err != nil {
		return employee, err
	}

	employee.EmployeeID = int64(rowID)

	// find user by id
	result, err := GetEmployeeByID(int64(employee.EmployeeID))
	if err != nil {
		return employee, err
	}

	return result, nil
}

func DeleteEmployee(id int64) (err error) {
	r, err := database.DB.Exec("DELETE FROM employees WHERE EmployeeID = ?", id)

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

func UpdateEmployee(e types.Employee) (err error) {

	r, err := database.DB.Exec(`UPDATE employees SET 
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
		e.Notes,
		e.ReportsTo,
		e.EmployeeID)

	ar, er := r.RowsAffected()
	var msg string
	if er != nil {
		msg = er.Error()
	}
	if ar == 0 {
		msg += "The Id Entered does not exist"
		err = fmt.Errorf(msg)
	}
	return err

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
		&e.Notes,
		&e.ReportsTo)
	if err != nil {
		return e, err
	}
	return e, nil
}
