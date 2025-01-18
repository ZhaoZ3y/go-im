package main

import (
	"goim/api/router"
	"goim/config"
	"goim/dao"
	"goim/utils/redi"
)

func main() {
	config.ConfigInit()
	dao.MysqlInit()
	redi.RedisInit()
	router.RouterInit()
}
