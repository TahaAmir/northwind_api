package types

import "gopkg.in/guregu/null.v3"

type Catogories struct {
	ID          int64       `json:"CategoryID,string,omitempty"`
	Name        string      `json:"CategoryName"`
	Description string      `json:"Description"`
	Picture     null.String `json:"Picture"`
}

//  int AI PK
// CategoryName varchar(15)
// Description longtext
// Picture
