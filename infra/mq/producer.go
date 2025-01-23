package mq

import (
	"github.com/IBM/sarama"
	"strings"
)

var producer sarama.AsyncProducer
var topic = "default_message"

// InitProducer 初始化生产者
func InitProducer(topicInput, hosts string) {
	topic = topicInput
	config := sarama.NewConfig()
	config.Producer.Compression = sarama.CompressionGZIP

	client, err := sarama.NewClient(strings.Split(hosts, ","), config)
	if err != nil {
		panic(err)
	}

	producer, err = sarama.NewAsyncProducerFromClient(client)
	if err != nil {
		panic(err)
	}
}

// SendMessage 发送消息
func SendMessage(msg []byte) {
	be := sarama.ByteEncoder(msg)
	producer.Input() <- &sarama.ProducerMessage{Topic: topic, Key: nil, Value: be}
}

// CloseProducer 关闭生产者
func CloseProducer() {
	if producer != nil {
		producer.Close()
	}
}
