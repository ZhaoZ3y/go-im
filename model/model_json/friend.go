package model_json

type Friend struct {
	Uuid     string `json:"uuid"`
	Username string `json:"username"`
	NickName string `json:"nick_name"`
	Avatar   string `json:"avatar"`
}
