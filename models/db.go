package models

// ABTest _
type User struct {
	Id           string `db:"id"`
	Guid         int64  `db:"guid"`
	Phone        string `db:"phone"`
	Username     string `db:"username"`
	Nickname     string `db:"nickname"`
	Sex          string `db:"sex"`
	balance      string `db:"balance"`
	LastLocation string `db:"LastLocation"`
}
