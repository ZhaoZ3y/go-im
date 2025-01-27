package mq

import (
	"github.com/IBM/sarama"
	"strings"
)

var consumer sarama.Consumer

type ConsumerCallback func(msg []byte)

// InitConsumer 初始化消费者
func InitConsumer(hosts string) {
	config := sarama.NewConfig()
	client, err := sarama.NewClient(strings.Split(hosts, ","), config)
	if err != nil {
		panic(err)
	}

	consumer, err = sarama.NewConsumerFromClient(client)
	if err != nil {
		panic(err)
	}
}

// ConsumerMessage 消费消息
func ConsumerMessage(callback ConsumerCallback) {
	PartitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}
	defer PartitionConsumer.Close()
	for {
		msg := <-PartitionConsumer.Messages()
		if msg != nil {
			callback(msg.Value)
		}
	}
}

// CloseConsumer 关闭消费者
func CloseConsumer() {
	if consumer != nil {
		consumer.Close()
	}
}
