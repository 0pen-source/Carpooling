package models

type Response struct {
	Commonresponse   `structs:",flatten"`
	Data interface{} `json:"data" structs:"data"`
}
type PhoneTestResponse struct {
	Exit bool `json:"exit" structs:"exit"`
}

type LoginResponse struct {
	Token         string `json:"token,omitempty" structs:"token"`
	Uid           string `json:"uid,omitempty" structs:"uid"`
	Sex           int    `json:"sex,omitempty" structs:"sex"`
	Balance       int64  `json:"balance,omitempty" structs:"balance"`
	UserName      string `json:"username,omitempty" structs:"username"`
	NickName      string `json:"nickname,omitempty" structs:"nickname"`
	Last_location string `json:"last_location,omitempty" structs:"last_location"`
	Exit          bool   `json:"exit,omitempty" structs:"exit"`
}

type Commonresponse struct {
	Code    int    `json:"code" structs:"code"`
	Message string `json:"message" structs:"message"`
}
