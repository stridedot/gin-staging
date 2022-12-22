package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"go_code/gintest/bootstrap/glog"
	"time"
)

type Claims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

var secret = []byte("隆冬强")

// access token 过期时间
const accessTokenExpireDuration = time.Hour * 24
// refresh token 过期时间
const refreshTokenExpireDuration = time.Hour * 24 * 30

// GenToken 生成JWT
func GenToken(userID int64, username string) map[string]string {
	var accessToken, refreshToken string
	var err error
	// 生成 access token
	accessToken, err = createToken(userID, username, accessTokenExpireDuration)
	if err != nil {
		glog.SL.Error("生成 access token 失败", err)
		return nil
	}

	// 生成 refresh token
	refreshToken, err = createToken(userID, username, refreshTokenExpireDuration)
	if err != nil {
		glog.SL.Error("生成 refresh token 失败", err)
		return nil
	}

	return map[string]string{
		"access_token": accessToken,
		"refresh_token": refreshToken,
	}
}

// RefreshToken 刷新 token
func RefreshToken(refreshToken string) (map[string]string, error) {
	// 判断 refresh token 是否过期，过期则返回重新登录
	claims, err := ParseToken(refreshToken)
	if err != nil {
		return nil, err
	}

	if claims.ExpiresAt.Before(time.Now()) {
		return nil, nil
	}

	accessToken, err := createToken(claims.UserID, claims.Username, accessTokenExpireDuration)
	if err != nil {
		glog.SL.Error("生成 access token 失败", err)
		return nil, err
	}

	m := map[string]string{
		"access_token": accessToken,
		"refresh_token": refreshToken,
	}

	return m, nil
}

// createToken 生成一个 token
func createToken(userID int64, username string, duration time.Duration) (string, error) {
	// 创建一个我们自己的声明
	claims := Claims{
		userID,
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			Issuer:    "隆冬强", // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(secret)
}

// ParseToken 解析 token
func ParseToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// 直接使用标准的Claim则可以直接使用Parse方法
		// token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
