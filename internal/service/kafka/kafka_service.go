package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	myconfig "rigon-chat-server/internal/config"
	"rigon-chat-server/pkg/zlog"
	"time"
)

var ctx = context.Background()

type kafkaService struct {
	ChatWriter *kafka.Writer
	ChatReader *kafka.Reader
	KafkaConn  *kafka.Conn
}

var KafkaService = new(kafkaService)

// KafkaInit 初始化 kafka
func (k *kafkaService) KafkaInit() {
	kafkaConfig := myconfig.GetConfig().KafkaConfig
	k.ChatWriter = &kafka.Writer{
		Addr:                   kafka.TCP(kafkaConfig.HostPort),
		Topic:                  kafkaConfig.ChatTopic,
		Balancer:               &kafka.Hash{},
		WriteTimeout:           kafkaConfig.Timeout * time.Second,
		RequiredAcks:           kafka.RequireNone,
		AllowAutoTopicCreation: false,
	}
	k.ChatReader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{kafkaConfig.HostPort},
		Topic:          kafkaConfig.ChatTopic,
		CommitInterval: kafkaConfig.Timeout * time.Second,
		GroupID:        "chat",
		StartOffset:    kafka.LastOffset,
	})
}

func (k *kafkaService) KafkaClose() {
	if err := k.ChatWriter.Close(); err != nil {
		zlog.Error(err.Error())
	}
	if err := k.ChatReader.Close(); err != nil {
		zlog.Error(err.Error())
	}
}

// CreateTopic 创建 kafka topic
func (k *kafkaService) CreateTopic() {
	kafkaConfig := myconfig.GetConfig().KafkaConfig
	chatTopic := kafkaConfig.ChatTopic

	// 连接至任意 kafka 节点
	var err error
	k.KafkaConn, err = kafka.Dial("tcp", kafkaConfig.HostPort)
	if err != nil {
		zlog.Fatal(err.Error())
	}

	topicConfigs := []kafka.TopicConfig{
		{
			Topic:             chatTopic,
			NumPartitions:     kafkaConfig.Partition,
			ReplicationFactor: 1,
		},
	}
	if err = k.KafkaConn.CreateTopics(topicConfigs...); err != nil {
		zlog.Fatal(err.Error())
	}
}
