package config

import (
	"github.com/kr/pretty"
	"gopkg.in/validator.v2"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"sync"
)

var (
	conf *Config
	once sync.Once
)

// Config 配置结构体
type Config struct {
	AppName    string           `yaml:"appName"`
	MySQL      MySQLConfig      `yaml:"mysql"`
	Redis      RedisConfig      `yaml:"redis"`
	Path       PathConfig       `yaml:"path"`
	MsgChannel MsgChannelConfig `yaml:"msgChannel"`
}

// MySQLConfig MySQL 配置
type MySQLConfig struct {
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
	Database    string `yaml:"database"`
	TablePrefix string `yaml:"table_prefix"`
}

// RedisConfig redis配置
type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

// PathConfig 文件路径配置
type PathConfig struct {
	FilePath string `yaml:"filePath"`
}

// MsgChannelConfig 消息通道配置
type MsgChannelConfig struct {
	ChannelType string `yaml:"channelType"`
	KafkaHost   string `yaml:"kafkaHost"`
	KafkaTopic  string `yaml:"kafkaTopic"`
}

func ConfigInit() {
	// 在工作目录下查找 config.yaml
	confFileRelPath := "config.yaml"
	content, err := ioutil.ReadFile(confFileRelPath)
	if err != nil {
		panic(err)
	}
	conf = new(Config)
	err = yaml.Unmarshal(content, conf)
	if err != nil {
		panic(err)
	}
	if err := validator.Validate(conf); err != nil {
		panic(err)
	}
	pretty.Printf("%+v\n", conf)
}

// GetConfig 获取配置实例
func GetConfig() *Config {
	once.Do(ConfigInit)
	return conf
}
