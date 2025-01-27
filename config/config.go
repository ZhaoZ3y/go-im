package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	AppName    string
	MySQL      MySQLConfig
	Redis      RedisConfig
	Path       PathConfig
	MsgChannel MsgChannelConfig
}

// MySQLConfig mysql配置
type MySQLConfig struct {
	Host        string
	Port        int
	Username    string
	Password    string
	Database    string
	TablePrefix string
}

// RedisConfig redis配置
type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
}

// PathConfig 文件路径配置
type PathConfig struct {
	FilePath string
}

// MsgChannelConfig 消息通道配置
type MsgChannelConfig struct {
	ChannelType string
	KafkaHost   string
	KafkaTopic  string
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
