package controllers

import "github.com/gin-gonic/gin"

const CtxUserIDKey = "userID"

// GetLoggedUserID 获取登录的用户 ID
func GetLoggedUserID(c *gin.Context) (interface{}, bool) {
	return c.Get(CtxUserIDKey)
}