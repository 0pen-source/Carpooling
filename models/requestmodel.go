package models

// Phonetest _
type Phonetest struct {
	CommonPayload `structs:",flatten"`

	Phone            string `binding:"required,omitempty,len=11,numeric" form:"phone" structs:"phone"`
	VerificationCode string `binding:"omitempty,len=6,numeric" form:"verification_code" structs:"verification_code"`
}

// UserMessage _
type UserMessage struct {
	CommonPayload `structs:",flatten"`

	Phone     string  `binding:"required,omitempty,len=11,numeric" form:"phone" structs:"phone"`
	Password  string  `binding:"-" form:"password" structs:"password"`
	ImageType string  `binding:"-" form:"image_type" structs:"image_type"`
	Nickname  string  `binding:"-" form:"nickname" structs:"nickname"`
	LastLon   float64 `binding:"-" form:"last_lon" structs:"last_lon"`
	LastLat   float64 `binding:"-" form:"last_lat" structs:"last_lat"`
	Sex       int     `binding:"-" form:"sex" structs:"sex"`
	Username  string  `binding:"-" form:"username" structs:"username"`
}

// UserMessage _
type TripMessage struct {
	CommonPayload `structs:",flatten"`

	Username        string  `binding:"-" form:"username" structs:"username"`
	Nickname        string  `binding:"-" form:"nickname" structs:"nickname"`
	Phone           string  `binding:"required,omitempty,len=11,numeric" form:"phone" structs:"phone"`
	TravelTime      string  `binding:"-" form:"travel_time" structs:"travel_time"`
	TravelTimeTitle string  `binding:"-" form:"travel_time_title" structs:"travel_time_title"`
	From            string  `binding:"required" form:"from" structs:"from"`
	FromLon         float64 `binding:"-" form:"from_lon" structs:"from_lon"`
	FromLat         float64 `binding:"-" form:"from_lat" structs:"from_lat"`
	Destination     string  `binding:"required" form:"destination" structs:"destination"`
	DestinationLon  float64 `binding:"-" form:"destination_lon" structs:"destination_lon"`
	DestinationLat  float64 `binding:"-" form:"destination_lat" structs:"destination_lat"`
	PayPrice        int64   `binding:"-" form:"pay_price" structs:"pay_price"`
	Surplus         int     `binding:"-" form:"surplus" structs:"surplus"`
	Distance        int64   `binding:"omitempty,numeric" form:"distance" structs:"distance"`
}

type Code struct {
	Code   string `binding:"-" form:"code" structs:"code"`
	Msg    string `binding:"-" form:"msg" structs:"msg"`
	SmUuid string `binding:"-" form:"smUuid" structs:"smUuid"`
}
type Index struct {
	CommonPayload `structs:",flatten"`
	Phone string  `binding:"required,omitempty,len=11,numeric" form:"phone" structs:"phone"`
}

// CommonPayload _
// it shouldn't be exported but it fails to bind, not familiar with go 1.8.3 features
type CommonPayload struct {
	IDFA          string  `binding:"omitempty,ne=00000000-0000-0000-0000-000000000000" form:"idfa" structs:"idfa"`
	MACAddress    string  `binding:"omitempty,len=32" form:"mac_address" structs:"mac_address"`
	IMEI          string  `binding:"omitempty,len=32" form:"imei" structs:"imei"`
	IMSI          string  `binding:"omitempty,len=32" form:"imsi" structs:"imsi"`
	AndroidID     string  `binding:"omitempty,len=32" form:"android_id" structs:"android_id"`
	AdvertisingID string  `binding:"-" form:"advertising_id" structs:"advertising_id"`
	Lon           float64 `binding:"-" form:"lon" structs:"lon"`
	Lat           float64 `binding:"-" form:"lat" structs:"lat"`
}
