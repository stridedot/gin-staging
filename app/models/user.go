package models

type User struct {
	UserID int64 `db:"user_id" json:"user_id,string"`
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"password"`
	CreatedAt string `db:"created_at" json:"created_at"`
}
