package main

import (
	"goim/config"
	"goim/dao"
)

func main() {
	config.ConfigInit()
	dao.MysqlInit()
}
