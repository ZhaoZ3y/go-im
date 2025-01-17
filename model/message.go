package model

type Message struct {
	ID          uint64 `gorm:"primaryKey;autoIncrement"`
	CreatedAt   int64  `gorm:"type:datetime(3);default:null"`
	UpdatedAt   int64  `gorm:"type:datetime(3);default:null"`
	DeletedAt   int64  `gorm:"index:idx_messages_deleted_at;default:null"`
	FromUserID  uint   `gorm:"index:idx_messages_from_user_id;" comment:"发送者ID"`
	ToUserID    uint   `gorm:"index:idx_messages_to_user_id;" comment:"接收者ID"`
	Content     string `gorm:"type:text;not null" comment:"消息内容"`
	Url         string `gorm:"type:varchar(255);default null" comment:"图片地址"`
	Picture     string `gorm:"type:text;default null" comment:"图片"`
	MessageType *int16 `gorm:"type:smallint;default 0" comment:"消息类型：1私聊消息，2群聊消息"`
	ContentType *int16 `gorm:"type:smallint;default 0" comment:"内容类型：1文本，2图片，3语言，4视频"`
}

func (m *Message) TableName() string {
	return "messages"
}
