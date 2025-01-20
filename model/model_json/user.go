package model_json

type User struct {
	Uuid     string `json:"uuid"`
	UserName string `json:"username"`
	NickName string `json:"nickname"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
}

type ChangePasswordReq struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
