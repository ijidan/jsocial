package test

type Message string

func NewMessage() Message {
	return Message("hello wire!")
}
