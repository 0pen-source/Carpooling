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
	RealtimeOrder  []ResponseTrip   `json:"realtime_order,omitempty"`
	RecommendOrder []ResponseTrip   `json:"recommend_order,omitempty"`
	Banner         []ResponseBanner `json:"banner,omitempty"`
}

type ResponseTrip struct {
	Guid                       string  `json:"guid" db:"guid"`
	UserName                   string  `json:"username" db:"username"`
	NickName                   string  `json:"nickname" db:"nickname"`
	Phone                      string  `json:"phone" db:"phone"`
	CreateTime                 int     `json:"create_time" db:"create_time"`
	TravelTime                 int     `json:"travel_time" db:"travel_time"`
	TravelTimeTitle            string  `json:"travel_time_title" db:"travel_time_title"`
	From                       string  `json:"from" db:"from"`
	FromRegion                 string  `json:"from_region" db:"from_region"`
	FromCity                   string  `json:"from_city" db:"from_city"`
	FromAccurateAddress        string  `json:"from_accurate_address" db:"from_accurate_address"`
	FromVagueAddress           string  `json:"from_vague_address" db:"from_vague_address"`
	FromLon                    float64 `json:"from_lon" db:"from_lon"`
	FromLat                    float64 `json:"from_lat" db:"from_lat"`
	Destination                string  `json:"destination" db:"destination"`
	DestinationRegion          string  `json:"destination_region" db:"destination_region"`
	DestinationCity            string  `json:"destination_city" db:"destination_city"`
	DestinationAccurateAddress string  `json:"destination_accurate_address" db:"destination_accurate_address"`
	DestinationVagueAddress    string  `json:"destination_vague_address" db:"destination_vague_address"`
	Distance                   float64 `json:"distance" db:"distance"`
	DestinationLon             float64 `json:"destination_lon" db:"destination_lon"`
	DestinationLat             float64 `json:"destination_lat" db:"destination_lat"`
	PayPrice                   int     `json:"pay_price" db:"pay_price"`
	Surplus                    int     `json:"surplus" db:"surplus"`
}

type ResponseBanner struct {
	Guid      string `json:"guid" db:"guid"`
	ImageType int    `json:"type" db:"type"`
	Image     string `json:"image" db:"image"`
	Click     string `json:"click" db:"click"`
}

type Upload struct {
	URL string `json:"url,omitempty" structs:"url"`
}

type Commonresponse struct {
	Code    int    `json:"code" structs:"code"`
	Message string `json:"message" structs:"message"`
}
