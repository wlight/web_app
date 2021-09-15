package models

type User struct {
	UserId   int64  `db:"userid"`
	Username string `db:"username"`
	Password string `db:"password"`
}
