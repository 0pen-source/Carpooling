package models

// ABTest _
type User struct {
	Id           string `db:"id"`
	Guid         string `db:"guid"`
	Phone        string `db:"phone"`
	Username     string `db:"username"`
	Password     string `db:"password"`
	Nickname     string `db:"nickname"`
	Sex          int    `db:"sex"`
	Balance      int64  `db:"balance"`
	LastLocation string `db:"LastLocation"`
}
