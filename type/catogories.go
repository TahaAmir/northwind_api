package types

type Catogories struct {
	ID          int64  `json:"CategoryID "`
	Name        string `json:"CategoryName"`
	Description string `json:"Description"`
	Picture     string `json:"Picture"`
}

// CategoryID int AI PK
// CategoryName varchar(15)
// Description longtext
// Picture
