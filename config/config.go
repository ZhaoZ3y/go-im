package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	AppName string
	MySQL   MySQLConfig
}

type MySQLConfig struct {
	Host        string
	Port        int
	Username    string
	Password    string
	Database    string
	TablePrefix string
}

var c Config

func ConfigInit() {
	//设置文件名
	viper.SetConfigName("config")
	//设置文件类型
	viper.SetConfigType("yaml")
	//设置文件路径
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	viper.Unmarshal(&c)
}

func GetConfig() Config {
	return c
}
