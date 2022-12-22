package user

import (
	"go_code/gintest/app/dao"
	daoUser "go_code/gintest/app/dao/user"
	"go_code/gintest/app/models"
	validator "go_code/gintest/app/validators/user"
	"go_code/gintest/pkg/jwt"
	"go_code/gintest/pkg/snowflake"
	"time"
)

// SignUp 注册用户
func SignUp(request *validator.SignUpRequest) error {
	user, err := daoUser.GetUserByUsername(request.Username)

	// 判断用户是否已存在
	if user != nil {
		return dao.UserExists
	}
	if err != nil {
		return err
	}

	// 注册用户
	// 生成 userID
	user = &models.User{
		UserID: snowflake.GenID(),
		Username: request.Username,
		Password: request.Password,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	return daoUser.InsertUser(user)
}

// SignIn 登录
func SignIn(request *validator.SignInRequest) (map[string]string, error) {
	user, err := daoUser.GetUserByUsername(request.Username)
	if err != nil {
		return nil, err
	}
	// 判断用户是否已存在
	if user == nil {
		return nil, dao.UserNotExists
	}
	// 验证密码
	if user.Password != daoUser.EncryptPassword(request.Password) {
		return nil, dao.PasswordError
	}

	return jwt.GenToken(user.UserID, user.Username), nil
}

// RefreshToken 刷新 token
func RefreshToken(request *validator.RefreshTokenRequest) (map[string]string, error) {
	token, err := jwt.RefreshToken(request.RefreshToken)
	if err != nil {
		return nil, err
	}

	if token == nil {
		return nil, dao.TokenExpiredError
	}

	return token, err
}

// GetUserByID 查询用户
func GetUserByID(userID int64) (*models.User, error) {
	user, err := daoUser.GetUserByID(userID)

	if err != nil {
		return nil, err
	}
	// 判断用户是否已存在
	if user == nil {
		return nil, dao.UserNotExists
	}

	return user, err
}