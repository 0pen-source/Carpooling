package models

type Response struct {
	Commonresponse   `structs:",flatten"`
	Data interface{} `json:"data" structs:"data"`
}
type PhoneTestResponse struct {
	Exit   bool   `json:"exit,omitempty" structs:"exit"`
	Status bool   `json:"status,omitempty" structs:"status"`
	ID     string `json:"id,omitempty" structs:"id"`
}

type LoginResponse struct {
	Token    string  `json:"token,omitempty" structs:"token"`
	Uid      string  `json:"uid,omitempty" structs:"uid"`
	Sex      int     `json:"sex,omitempty" structs:"sex"`
	Balance  int64   `json:"balance,omitempty" structs:"balance"`
	UserName string  `json:"username,omitempty" structs:"username"`
	NickName string  `json:"nickname,omitempty" structs:"nickname"`
	LastLon  float64 `json:"last_lon,omitempty" structs:"last_location"`
	LastLat  float64 `json:"last_lat,omitempty" structs:"last_location"`
	Exit     bool    `json:"exit,omitempty" structs:"exit"`
}
type IndexResponse struct {
	RealtimeOrder  []ResponseTrip `json:"realtime_order"`
	RecommendOrder []ResponseTrip `json:"recommend_order"`
	banner         []ResponseTrip `json:"banner"`
}

type ResponseTrip struct {
	Guid            string  `json:"guid,omitempty" structs:"guid"`
	UserName        string  `json:"username,omitempty" structs:"username"`
	NickName        string  `json:"nickname,omitempty" structs:"nickname"`
	Phone           string  `json:"phone,omitempty" structs:"phone"`
	CreateTime      int64   `json:"create_time,omitempty" structs:"create_time"`
	TravelTime      string  `json:"travel_time,omitempty" structs:"travel_time"`
	TravelTimeTitle string  `json:"travel_time_title,omitempty" structs:"travel_time_title"`
	From            string  `json:"from,omitempty" structs:"from"`
	FromLon         float64 `json:"from_lon,omitempty" structs:"from_lon"`
	FromLat         float64 `json:"from_lat,omitempty" structs:"from_lat"`
	Destination     string  `json:"destination,omitempty" structs:"destination"`
	DestinationLon  float64 `json:"destination_lon,omitempty" structs:"destination_lon"`
	DestinationLat  float64 `json:"destination_lat,omitempty" structs:"destination_lat"`
	PayPrice        int64   `json:"pay_price,omitempty" structs:"pay_price"`
	Surplus         int     `json:"surplus,omitempty" structs:"surplus"`
}

type Upload struct {
	URL string `json:"url,omitempty" structs:"url"`
}

type Commonresponse struct {
	Code    int    `json:"code" structs:"code"`
	Message string `json:"message" structs:"message"`
}
