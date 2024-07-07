package model

// PersonRequestModel struct
type PersonRequestModel struct {
	Id          uint   `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}
