package user

import (
	"github.com/gin-gonic/gin"
	"go_code/gintest/app/controllers"
	"go_code/gintest/app/dao"
	serviceUser "go_code/gintest/app/services/user"
	"go_code/gintest/app/validators"
	validatorUser "go_code/gintest/app/validators/user"
	"go_code/gintest/bootstrap/glog"
)

func SignUp(c *gin.Context) {
	// 验证参数
	request := new(validatorUser.SignUpRequest)
	if err := c.ShouldBindJSON(request); err != nil {
		s := validators.Error(err, request)
		controllers.ErrorRes(c, controllers.CodeUnexpectedParam, s)
		return
	}

	err := serviceUser.SignUp(request)
	if err == dao.UserExists {
		controllers.ErrorRes(c, controllers.CodeUserExists, nil)
		return
	}
	if err != nil {
		glog.SL.Error("注册用户", err)
		controllers.ErrorRes(c, controllers.CodeInternalError, nil)
		return
	}
	controllers.SuccessRes(c, nil)
	return
}

// SignIn 登录
func SignIn(c *gin.Context) {
	request := new(validatorUser.SignInRequest)
	if err := c.ShouldBindJSON(request); err != nil {
		s := validators.Error(err, request)
		controllers.ErrorRes(c, controllers.CodeUnexpectedParam, s)
		return
	}

	token, err := serviceUser.SignIn(request)
	if err == dao.UserNotExists {
		controllers.ErrorRes(c, controllers.CodeUserNotExists, nil)
		return
	}
	if err == dao.PasswordError {
		controllers.ErrorRes(c, controllers.CodePasswordError, nil)
		return
	}
	if err != nil {
		controllers.ErrorRes(c, controllers.CodeInternalError, nil)
		return
	}

	data := &controllers.Data{Data: token}
	controllers.SuccessRes(c, data)
	return
}

// RefreshToken 刷新 access token
func RefreshToken(c *gin.Context) {
	request := new(validatorUser.RefreshTokenRequest)
	if err := c.ShouldBindJSON(request); err != nil {
		s := validators.Error(err, request)
		controllers.ErrorRes(c, controllers.CodeInvalidToken, s)
		return
	}

	token, err := serviceUser.RefreshToken(request)
	if err == dao.TokenExpiredError {
		controllers.ErrorRes(c, controllers.CodeTokenExpired, nil)
		return
	}
	if err != nil {
		controllers.ErrorRes(c, controllers.CodeInternalError, nil)
		return
	}

	data := &controllers.Data{Data: token}
	controllers.SuccessRes(c, data)
	return
}
