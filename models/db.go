package models

// ABTest _
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
