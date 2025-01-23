package main

import (
	"goim/api/router"
	"goim/config"
	"goim/dao"
	"goim/infra/mq"
	constant "goim/utils/const"
	"goim/utils/redi"
)

func main() {
	config.ConfigInit()
	if config.GetConfig().MsgChannel.ChannelType == constant.KAFKA {
		mq.InitProducer(config.GetConfig().MsgChannel.KafkaTopic, config.GetConfig().MsgChannel.KafkaHost)
		mq.InitConsumer(config.GetConfig().MsgChannel.KafkaHost)
	}

	dao.MysqlInit()
	redi.RedisInit()
	router.RouterInit()
}
