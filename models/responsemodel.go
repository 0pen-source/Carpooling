package models

type Response struct {
	Commonresponse   `structs:",flatten"`
	Data interface{} `json:"data" structs:"data"`
}
type PhoneTestResponse struct {
	Exit bool `json:"exit" structs:"exit"`
}

type Commonresponse struct {
	Code    int    `json:"code" structs:"code"`
	Message string `json:"message" structs:"message"`
}
