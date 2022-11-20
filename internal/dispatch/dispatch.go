package dispatch

import (
	"context"
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/fatih/color"
	"github.com/ijidan/jsocial/internal/app/message/repository"
	"github.com/ijidan/jsocial/internal/global"
	"github.com/ijidan/jsocial/internal/pkg/kafka"
)

const (
	TopicUserNew    = "topic.user.new"
	TopicMessageNew = "topic.message.new"
)

func PublishMessage(senderId uint64, receiverId uint64, groupId uint64, messageContentId int64, messageType string, messageBody []byte, messageExtra []byte) {
	k := kafka.NewKafka(global.GR.Conf.PubSub.Brokers)
	messageContent := repository.MessageContent{
		SenderId:         senderId,
		ReceiverId:       receiverId,
		GroupId:          groupId,
		MessageContentId: messageContentId,
		MessageType:      MessageTypeToDbType[messageType],
		MessageBody:      messageBody,
		MessageExtra:     messageExtra,
	}
	content, _ := json.Marshal(messageContent)
	err := k.Publish(TopicMessageNew, nil, content)
	if err != nil {
		//_ = repository.SendErrorEmail(err)
	}

}

func SubscribeMessage(ctx context.Context) error {
	f := func(message *sarama.ConsumerMessage) error {
		color.Yellow("kafka message received:%s", string(message.Value))

		var messageContent repository.MessageContent
		_ = json.Unmarshal(message.Value, &messageContent)
		_, err := repository.CreateMessageContent(global.GR.Db, messageContent)
		if err != nil {
			color.Red("save message content error:%s", err.Error())
		}
		return nil
	}
	k := kafka.NewKafka(global.GR.Conf.PubSub.Brokers)
	client, err := k.Subscribe(TopicMessageNew, TopicUserNew, f)
	if err != nil {
		return err
	}
	go func(client sarama.Client) {
		for {
			select {
			case <-ctx.Done():
				color.Red("close kafka client")
				_ = client.Close()
				return
			}
		}
	}(client)
	return nil
}
