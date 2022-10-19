package test

type Channel struct {
	Message Message
}

func (c *Channel) GetMsg() Message {
	return c.Message
}

func NewChannel(m Message) Channel {
	return Channel{Message: m}
}
