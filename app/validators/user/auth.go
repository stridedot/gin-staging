package user

// SignUpRequest 注册参数 validator
type SignUpRequest struct {
	Username string `json:"username" binding:"required" label:"用户名"`
	Password string `json:"password" binding:"required,min=6" label:"密码"`
}

// SignInRequest 登录 validator
type SignInRequest struct {
	Username string `json:"username" binding:"required" label:"用户名"`
	Password string `json:"password" binding:"required" label:"密码"`
}

// RefreshTokenRequest 刷新 access token validator
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}