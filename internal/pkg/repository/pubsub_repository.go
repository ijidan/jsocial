package repository

import (
	"github.com/ijidan/jsocial/internal/pkg/global"
	"github.com/ijidan/jsocial/internal/pkg/pubsub"
	"github.com/spf13/cast"
)

const (
	TopicNewUser = "user.new"
)

//func ProduceNewUser(userId uint64) error {
//	kafka := pkg.NewKafkaS(pkg.Conf.PubSub.Brokers)
//	content := cast.ToString(userId)
//	return kafka.PublishMessage(TopicNewUser, content)
//}
//
//func ConsumerNewUser() error {
//	f := func(msg *sarama.ConsumerMessage) error {
//		userId := cast.ToUint64(msg.Value)
//		extra := fmt.Sprintf("received message,key:%s value:%s", cast.ToString(msg.Key), cast.ToString(msg.Value))
//		return SendNewUserEmail(userId, extra)
//	}
//
//	kafka := pkg.NewKafkaS(pkg.Conf.PubSub.Brokers)
//	return kafka.Subscribe(TopicNewUser, TopicNewUser, f)
//}

func PublishNewUser(userId uint64) {
	pubSub := pubsub.NewPubSub(global.Conf.PubSub.Brokers)
	content := []byte(cast.ToString(userId))
	err := pubSub.PublishMessage(TopicNewUser, content)
	if err != nil {
		_ = SendErrorEmail(err)
	}
}
func SubscribeNewUser() {
	pubSub := pubsub.NewPubSub(global.Conf.PubSub.Brokers)
	message, err := pubSub.Subscribe(TopicNewUser)
	if err != nil {
		_ = SendErrorEmail(err)
	}
	for msg := range message {
		msg.Ack()
		//userId := cast.ToUint64(string(msg.Payload))
		//extra := fmt.Sprintf("received message: %s, payload: %s", msg.UUID, string(msg.Payload))
		//err = SendNewUserEmail(userId, extra)
		//if err==nil{
		//	msg.Ack()
		//}
	}
}
