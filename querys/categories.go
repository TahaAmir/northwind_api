package querys

import (
	"golang-crud-rest-api/database"
	types "golang-crud-rest-api/type"
)

func CreateCatogorry(c types.Catogories) error {

	_, err := database.DB.Exec("INSERT INTO categories (CategoryName,Description,Picture) VALUES (?,?,?) ", c.Name, c.Description, c.Picture)
	return err
}

func DeleteCatogory(id int64) error {
	_, err := database.DB.Exec("DELETE FROM categories WHERE CategoryID =?", id)
	return err
}

func UpdateCategory(c types.Catogories) error {
	_, err := database.DB.Exec(" UPDATE categories SET CategoryName= ?, Description=?,Picture=? WHERE CategoryID=?", c.Name, c.Description, c.Picture, c.ID)
	return err
}

func GetCategory() ([]types.Catogories, error) {
	category := []types.Catogories{}

	rows, err := database.DB.Query("SELECT CategoryID,CategoryName,Description,Picture FROM categories")
	if err != nil {
		return category, err
	}

	for rows.Next() {
		var c types.Catogories
		err = rows.Scan(&c.ID, &c.Name, &c.Description, &c.Picture)
		if err != nil {
			return category, err
		}
		category = append(category, c)
	}
	return category, nil
}

func GetCategoryById(id int64) (types.Catogories, error) {

	var category types.Catogories

	row := database.DB.QueryRow("SELECT CategoryID,CategoryName,Description,Picture FROM categories WHERE CategoryID = ?", id)
	err := row.Scan(&category.ID, &category.Name, &category.Description, &category.Picture)
	if err != nil {
		return category, err
	}
	return category, nil
}
