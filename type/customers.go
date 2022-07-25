package types

import (
	"gopkg.in/guregu/null.v3"
)

type Customers struct {
	ID             int64       `json:"CustomerID"`
	CompanyName    string      `json:"CompanyName"`
	ContactName    string      `json:"ContactName"`
	ContactTitle   string      `json:"ContactTiltle"`
	City           null.String `json:"City"`
	Address        null.String `json:"Address"`
	Region         null.String `json:"Region"`
	PostalCode     null.String `json:"PostalCode"`
	Country        null.String `json:"Country"`
	Phone          null.String `json:"Phone"`
	Fax            null.String `json:"Fax"`
	Image          any         `json:"Image"`
	ImageThumbnail any         `json:"ImageThumbnail"`
}

// func (s null.String) MarshalJSON() ([]byte, error) {
// 	if !s.Valid {
// 		return []byte("null"), nil
// 	}
// 	return json.Marshal(s.String)
// }

// func (s *null.String) UnmarshalJSON(data []byte) error {
// 	if string(data) == "null" {
// 		s.String, s.Valid = "", false
// 		return nil
// 	}
// 	s.String, s.Valid = string(data), true
// 	return nil
// }
