package constant

const (
	SenderTypeUser  = 1
	SenderTypeGroup = 2

	ReceiverTypeUser  = 1
	ReceiverTypeGroup = 2

	MessageTypeText     = 0
	MessageTypeLocation = 1
	MessageTypeFace     = 2
	MessageTypeSound    = 3
	MessageTypeImage    = 4
	MessageTypeFile     = 5
	MessageTypeVideo    = 6
)

const (
	MessageIndexStatusToSend    = 1
	MessageIndexStatusSent      = 2
	MessageIndexStatusConfirmed = 3
	MessageIndexStatusDelete    = 0
)
