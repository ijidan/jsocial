package message

import (
	"github.com/ijidan/jsocial/api/proto_build"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

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

func CreateUserTextMessage(db *gorm.DB, senderId uint64, toUserId uint64, text *proto_build.TextMessage) (*MessageContent, error) {
	messageContent := &MessageContent{
		ID:    0,
		Type:  MessageTypeText,
		Body:  text.Content,
		Extra: "",
	}
	err := db.Create(messageContent).Error
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return messageContent, nil
}

func CreateUserLocationMessage(db *gorm.DB, senderId uint64, toUserId uint64, location *proto_build.LocationMessage) (*MessageContent, error) {
	messageContent := &MessageContent{
		ID:    0,
		Type:  MessageTypeText,
		Body:  text.Content,
		Extra: "",
	}
	err := db.Create(messageContent).Error
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return messageContent, nil
}

func CreateUserFaceMessage(db *gorm.DB, senderId uint64, toUserId uint64, face *proto_build.FaceMessage) (*Message, error) {
	message := &Message{
		SenderID:     senderId,
		SenderType:   SenderTypeUser,
		ReceiverID:   toUserId,
		ReceiverType: ReceiverTypeUser,
		MessageType:  MessageTypeFace,
		Data:         face.String(),
	}
	err := db.Create(message).Error
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return message, nil
}

func CreateUserSoundMessage(db *gorm.DB, senderId uint64, toUserId uint64, sound *proto_build.SoundMessage) (*Message, error) {
	message := &Message{
		SenderID:     senderId,
		SenderType:   SenderTypeUser,
		ReceiverID:   toUserId,
		ReceiverType: ReceiverTypeUser,
		MessageType:  MessageTypeSound,
		Data:         sound.String(),
	}
	err := db.Create(message).Error
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return message, nil
}

func CreateUserVideoMessage(db *gorm.DB, senderId uint64, toUserId uint64, video *proto_build.VideoMessage) (*Message, error) {
	message := &Message{
		SenderID:     senderId,
		SenderType:   SenderTypeUser,
		ReceiverID:   toUserId,
		ReceiverType: ReceiverTypeUser,
		MessageType:  MessageTypeVideo,
		Data:         video.String(),
	}
	err := db.Create(message).Error
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return message, nil
}

func CreateUserImageMessage(db *gorm.DB, senderId uint64, toUserId uint64, image *proto_build.ImageMessageItem) (*Message, error) {
	message := &Message{
		SenderID:     senderId,
		SenderType:   SenderTypeUser,
		ReceiverID:   toUserId,
		ReceiverType: ReceiverTypeUser,
		MessageType:  MessageTypeImage,
		Data:         image.String(),
	}
	err := db.Create(message).Error
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return message, nil
}

func CreateUserFileMessage(db *gorm.DB, senderId uint64, toUserId uint64, file *proto_build.FileMessage) (*Message, error) {
	message := &Message{
		SenderID:     senderId,
		SenderType:   SenderTypeUser,
		ReceiverID:   toUserId,
		ReceiverType: ReceiverTypeUser,
		MessageType:  MessageTypeFile,
		Data:         file.String(),
	}
	err := db.Create(message).Error
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return message, nil
}

func CreateGroupTextMessage(db *gorm.DB, senderId uint64, toGroupId uint64, atUserId []uint64, text *proto_build.TextMessage) (*Message, error) {
	message := &Message{
		SenderID:     senderId,
		SenderType:   SenderTypeGroup,
		ReceiverID:   toGroupId,
		AtUserID:     pkg.ConvertUintSlice2SepString(atUserId),
		ReceiverType: ReceiverTypeGroup,
		MessageType:  MessageTypeText,
		Data:         text.Content,
	}
	err := db.Create(message).Error
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return message, nil
}

func CreateGroupLocationMessage(db *gorm.DB, senderId uint64, toGroupId uint64, atUserId []uint64, location *proto_build.LocationMessage) (*Message, error) {
	message := &Message{
		SenderID:     senderId,
		SenderType:   SenderTypeGroup,
		ReceiverID:   toGroupId,
		AtUserID:     pkg.ConvertUintSlice2SepString(atUserId),
		ReceiverType: ReceiverTypeGroup,
		MessageType:  MessageTypeLocation,
		Data:         location.String(),
	}
	err := db.Create(message).Error
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return message, nil
}

func CreateGroupFaceMessage(db *gorm.DB, senderId uint64, toGroupId uint64, atUserId []uint64, face *proto_build.FaceMessage) (*Message, error) {
	message := &Message{
		SenderID:     senderId,
		SenderType:   SenderTypeGroup,
		ReceiverID:   toGroupId,
		AtUserID:     pkg.ConvertUintSlice2SepString(atUserId),
		ReceiverType: ReceiverTypeGroup,
		MessageType:  MessageTypeFace,
		Data:         face.String(),
	}
	err := db.Create(message).Error
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return message, nil
}

func CreateGroupSoundMessage(db *gorm.DB, senderId uint64, toGroupId uint64, atUserId []uint64, sound *proto_build.SoundMessage) (*Message, error) {
	message := &Message{
		SenderID:     senderId,
		SenderType:   SenderTypeGroup,
		ReceiverID:   toGroupId,
		AtUserID:     pkg.ConvertUintSlice2SepString(atUserId),
		ReceiverType: ReceiverTypeGroup,
		MessageType:  MessageTypeSound,
		Data:         sound.String(),
	}
	err := db.Create(message).Error
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return message, nil
}

func CreateGroupVideoMessage(db *gorm.DB, senderId uint64, toGroupId uint64, atUserId []uint64, video *proto_build.VideoMessage) (*Message, error) {
	message := &Message{
		SenderID:     senderId,
		SenderType:   SenderTypeGroup,
		ReceiverID:   toGroupId,
		AtUserID:     pkg.ConvertUintSlice2SepString(atUserId),
		ReceiverType: ReceiverTypeGroup,
		MessageType:  MessageTypeVideo,
		Data:         video.String(),
	}
	err := db.Create(message).Error
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return message, nil
}

func CreateGroupImageMessage(db *gorm.DB, senderId uint64, toGroupId uint64, atUserId []uint64, image *proto_build.ImageMessageItem) (*Message, error) {
	message := &Message{
		SenderID:     senderId,
		SenderType:   SenderTypeGroup,
		ReceiverID:   toGroupId,
		AtUserID:     pkg.ConvertUintSlice2SepString(atUserId),
		ReceiverType: ReceiverTypeGroup,
		MessageType:  MessageTypeImage,
		Data:         image.String(),
	}
	err := db.Create(message).Error
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return message, nil
}

func CreateGroupFileMessage(db *gorm.DB, senderId uint64, toGroupId uint64, atUserId []uint64, file *proto_build.FileMessage) (*Message, error) {
	message := &Message{
		SenderID:     senderId,
		SenderType:   SenderTypeGroup,
		ReceiverID:   toGroupId,
		AtUserID:     pkg.ConvertUintSlice2SepString(atUserId),
		ReceiverType: ReceiverTypeGroup,
		MessageType:  MessageTypeFile,
		Data:         file.String(),
	}
	err := db.Create(message).Error
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return message, nil
}
