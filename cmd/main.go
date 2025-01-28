package main

import (
	"goim/api/router"
	"goim/config"
	"goim/dao"
	"goim/infra/mq"
	"goim/infra/socket"
	constant "goim/utils/const"
	"goim/utils/redi"
	"net/http"
	"time"
)

func main() {
	if config.GetConfig().MsgChannel.ChannelType == constant.KAFKA {
		mq.InitProducer(config.GetConfig().MsgChannel.KafkaTopic, config.GetConfig().MsgChannel.KafkaHost)
		mq.InitConsumer(config.GetConfig().MsgChannel.KafkaHost)
		go mq.ConsumerMessage(socket.ConsumerKafkaMessage)
	}
	dao.MysqlInit()
	redi.RedisInit()
	r := router.RouterInit()

	go socket.WsSever.Run()

	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
