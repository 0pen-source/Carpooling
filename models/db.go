package models

// User _
type User struct {
	Id                 string  `db:"id"`
	Guid               string  `db:"guid"`
	Phone              string  `db:"phone"`
	IDCardsURL         string  `db:"id_cards_url"`
	DriverURL          string  `db:"driver_url"`
	Username           string  `db:"username"`
	Password           string  `db:"password"`
	Nickname           string  `db:"nickname"`
	PortraitURL        string  `db:"portrait_url"`
	Sex                int64   `db:"sex"`
	Balance            int64   `db:"balance"`
	LastLat            float64 `db:"last_lat"`
	LastLon            float64 `db:"last_lon"`
	RealNameAuthStatus int64   `db:"real_name_auth_status"`
	DriverAuthStatus   int64   `db:"driver_auth_status"`
}

// PassengersTrip _
type PassengersTrip struct {
	Guid                       string  `db:"guid"`
	UserName                   string  `db:"username"`
	NickName                   string  `db:"nickname"`
	PortraitURL                string  `db:"portrait_url"`
	Phone                      string  `db:"phone"`
	Msg                        string  `db:"msg"`
	CreateTime                 int64   `db:"create_time"`
	TravelTime                 string  `db:"travel_time"`
	TravelTimeTitle            string  `db:"travel_time_title"`
	From                       string  `db:"From"`
	FromRegion                 string  `db:"from_region"`
	FromCity                   string  `db:"from_city"`
	FromDistrict               string  `db:"from_district"`
	FromAccurateAddress        string  `db:"from_accurate_address"`
	FromVagueAddress           string  `db:"from_vague_address"`
	FromLon                    float64 `db:"from_lon"`
	FromLat                    float64 `db:"from_lat"`
	Destination                string  `db:"destination"`
	DestinationRegion          string  `db:"destination_region"`
	DestinationCity            string  `db:"destination_city"`
	DestinationDistrict        string  `db:"destination_district"`
	DestinationAccurateAddress string  `db:"destination_accurate_address"`
	DestinationVagueAddress    string  `db:"destination_vague_address"`
	Source                     string  `db:"source"`
	DestinationLon             float64 `db:"destination_lon"`
	DestinationLat             float64 `db:"destination_lat"`
	Mileage                    float64 `db:"mileage"`
	PayPrice                   int64   `db:"pay_price"`
	Surplus                    int64   `db:"surplus"`
	SeatNum                    int64   `db:"seat_num"`
	Complete                   int64   `db:"complete"`
}

// AlreadyConnDriver _
type AlreadyConnDriver struct {
	Guid       string `db:"guid"`
	Phone      string `db:"phone"`
	UpdateTime int64  `db:"update_time"`
}

// DriverTrip _
type DriverTrip struct {
	Guid                       string  `db:"guid"`
	UserName                   string  `db:"username"`
	NickName                   string  `db:"nickname"`
	PortraitURL                string  `db:"portrait_url"`
	Phone                      string  `db:"phone"`
	Msg                        string  `db:"msg"`
	CreateTime                 int64   `db:"create_time"`
	TravelTime                 string  `db:"travel_time"`
	TravelTimeTitle            string  `db:"travel_time_title"`
	From                       string  `db:"From"`
	FromRegion                 string  `db:"from_region"`
	FromCity                   string  `db:"from_city"`
	FromDistrict               string  `db:"from_district"`
	FromAccurateAddress        string  `db:"from_accurate_address"`
	FromVagueAddress           string  `db:"from_vague_address"`
	FromLon                    float64 `db:"from_lon"`
	FromLat                    float64 `db:"from_lat"`
	Destination                string  `db:"destination"`
	DestinationRegion          string  `db:"destination_region"`
	DestinationCity            string  `db:"destination_city"`
	DestinationDistrict        string  `db:"destination_district"`
	DestinationAccurateAddress string  `db:"destination_accurate_address"`
	DestinationVagueAddress    string  `db:"destination_vague_address"`
	Source                     string  `db:"source"`
	DestinationLon             float64 `db:"destination_lon"`
	DestinationLat             float64 `db:"destination_lat"`
	Mileage                    float64 `db:"mileage"`
	PayPrice                   int64   `db:"pay_price"`
	Surplus                    int64   `db:"surplus"`
	SeatNum                    int64   `db:"seat_num"`
	Complete                   int64   `db:"complete"`
}
