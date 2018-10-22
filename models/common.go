package models

type Position struct {
	Status int    `json:"status,omitempty" structs:"status"`
	Result Result `json:"result,omitempty" structs:"result"`
}
type Result struct {
	AddressComponent   AddressComponent `json:"addressComponent,omitempty" structs:"addressComponent"`
	SematicDescription string           `json:"sematic_description,omitempty" structs:"sematic_description"`
}
type AddressComponent struct {
	Country  string `json:"country,omitempty" structs:"country"`
	Province string `json:"province,omitempty" structs:"province"`
	City     string `json:"city,omitempty" structs:"city"`
	District string `json:"district,omitempty" structs:"district"`
}
