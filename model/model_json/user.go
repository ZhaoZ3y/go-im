package model_json

type User struct {
	Uuid     string `json:"uuid"`
	UserName string `json:"username"`
	NickName string `json:"nickname"`
	Email    string `json:"email"`
	PassWord string `json:"password"`
	Avatar   string `json:"avatar"`
}

type LoginReq struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

type RefreshTokenReq struct {
	RefreshToken string `json:"refresh_token"`
}
