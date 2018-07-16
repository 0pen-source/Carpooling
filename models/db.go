package models

// User _
type User struct {
	Id                 string `db:"id"`
	Guid               string `db:"guid"`
	Phone              string `db:"phone"`
	IDCardsURL         string `db:"id_cards_url"`
	DriverURL          string `db:"driver_url"`
	Username           string `db:"username"`
	Password           string `db:"password"`
	Nickname           string `db:"nickname"`
	Sex                int    `db:"sex"`
	Balance            int64  `db:"balance"`
	LastLocation       string `db:"last_location"`
	RealNameAuthStatus int    `db:"real_name_auth_status"`
	DriverAuthStatus   int    `db:"driver_auth_status"`
}

// PassengersTrip _
type PassengersTrip struct {
	Guid            string  `db:"guid"`
	Username        string  `db:"username"`
	Nickname        string  `db:"nickname"`
	Phone           string  `db:"phone"`
	CreateTime      int64   `db:"create_time"`
	TravelTime      string  `db:"travel_time"`
	TravelTimeTitle string  `db:"travel_time_title"`
	From            string  `db:"From"`
	FromLon         float64 `db:"from_lon"`
	FromLat         float64 `db:"from_lat"`
	Destination     string  `db:"destination"`
	DestinationLon  float64 `db:"destination_lon"`
	DestinationLat  float64 `db:"destination_lat"`
	PayPrice        int64   `db:"pay_price"`
	Surplus         int     `db:"surplus"`
}
