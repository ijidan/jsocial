package test

import "fmt"

type BroadCast struct {
	Channel Channel
}

func (b *BroadCast) Start() {
	msg := b.Channel.GetMsg()
	fmt.Println(msg)
}

func NewBroadCast(c Channel) BroadCast {
	return BroadCast{Channel: c}
}
