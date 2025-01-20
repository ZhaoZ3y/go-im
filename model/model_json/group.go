package model_json

import "time"

type Group struct {
	ID        int64     `json:"id"`
	Uuid      string    `json:"uuid"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	Notice    string    `json:"notice"`
	Avatar    string    `json:"avatar"`
}
