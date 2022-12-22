package user

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"go_code/gintest/app/models"
	"go_code/gintest/bootstrap/gdb"
)

const salt = "gin"

// GetUserByID 查询指定用户名的用户
func GetUserByID(userID int64) (*models.User, error) {
	user := new(models.User)
	sqlStr := "SELECT user_id, username, `password`, created_at " +
		"FROM `user` WHERE user_id = ?"
	err := gdb.DB.Get(user, sqlStr, userID)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return user, err
}

// GetUserByUsername 查询指定用户名的用户
func GetUserByUsername(username string) (*models.User, error) {
	user := new(models.User)
	sqlStr := "SELECT user_id, username, `password`, created_at " +
		"FROM `user` WHERE username = ?"
	err := gdb.DB.Get(user, sqlStr, username)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return user, err
}

// InsertUser 注册用户
func InsertUser(user *models.User) error {
	sqlStr := "INSERT INTO `user` " +
		"(user_id, username, password, created_at) VALUES (?, ?, ?, ?)"
	_, err := gdb.DB.Exec(
		sqlStr,
		user.UserID,
		user.Username,
		EncryptPassword(user.Password),
		user.CreatedAt,
	)

	return err
}

// EncryptPassword 密码加密
func EncryptPassword(password string) string {
	h := md5.New()
	h.Write([]byte(salt))
	return hex.EncodeToString(h.Sum([]byte(password)))
}
