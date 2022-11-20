package repository

import (
	"github.com/ijidan/jsocial/internal/global"
	"github.com/ijidan/jsocial/internal/pkg/pubsub"
	"github.com/spf13/cast"
)

const (
	TopicNewUser = "user.new"
)

func PublishNewUser(userId uint64) {
	pubSub := pubsub.NewPubSub(global.GR.Conf.PubSub.Brokers)
	content := []byte(cast.ToString(userId))
	err := pubSub.PublishMessage(TopicNewUser, content)
	if err != nil {
		_ = SendErrorEmail(err)
	}
}
func SubscribeNewUser() {
	pubSub := pubsub.NewPubSub(global.GR.Conf.PubSub.Brokers)
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
