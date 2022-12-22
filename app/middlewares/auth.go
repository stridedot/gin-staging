package middlewares

import (
	"github.com/gin-gonic/gin"
	"go_code/gintest/app/controllers"
	"go_code/gintest/pkg/jwt"
	"strings"
)

// JWTAuth 登录认证
func JWTAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 获取 Authorization 的值，格式 Bearer abc123abc123abc
		authorization := c.Request.Header.Get("Authorization")
		if authorization == "" {
			controllers.ErrorRes(c, controllers.CodeUnauthorized, nil)
			c.Abort()
			return
		}
		// token = parts[1]
		parts := strings.SplitN(authorization, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			controllers.ErrorRes(c, controllers.CodeInvalidToken, nil)
			c.Abort()
			return
		}
		claims, err := jwt.ParseToken(parts[1])
		if err != nil {
			controllers.ErrorRes(c, controllers.CodeInvalidToken, nil)
			c.Abort()
			return
		}
		c.Set(controllers.CtxUserIDKey, claims.UserID)
		c.Next()
	}
}
