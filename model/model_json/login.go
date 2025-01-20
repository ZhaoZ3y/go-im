package model_json

type LoginReq struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

type RefreshTokenReq struct {
	RefreshToken string `json:"refresh_token"`
}
