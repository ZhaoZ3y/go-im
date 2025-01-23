package constant

const (
	UserExisted   = "用户已存在"
	UserLoginErr  = "账号或密码错误"
	RefreshFailed = "该令牌在10分钟内已被刷新，无法再次刷新"
	TokenExpired  = "令牌已过期"
	GroupNotExits = "群组不存在"
	FriendExisted = "已经是好友了"
)

const (
	HEAT_BEAT = "heatbeat"
	PONG      = "pong"

	// 消息类型，单聊或者群聊
	MESSAGE_TYPE_USER  = 1
	MESSAGE_TYPE_GROUP = 2

	// 消息内容类型
	TEXT         = 1
	FILE         = 2
	IMAGE        = 3
	AUDIO        = 4
	VIDEO        = 5
	AUDIO_ONLINE = 6
	VIDEO_ONLINE = 7

	// 消息队列类型
	GO_CHANNEL = "gochannel"
	KAFKA      = "kafka"
)
