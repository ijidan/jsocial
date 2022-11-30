package message

import (
	"encoding/json"
	"github.com/ijidan/jsocial/api/proto_build"
	"github.com/ijidan/jsocial/internal/constant"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func CreateUserTextMessage(db *gorm.DB, senderId uint64, toUserId uint64, text *proto_build.TextMessage) (*MessageContent, error) {
	body := text.Content
	messageContent := &MessageContent{
		ID:    0,
		Type:  constant.MessageTypeText,
		Body:  body,
		Extra: "",
	}
	err := db.Create(messageContent).Error
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return messageContent, nil
}

func CreateUserLocationMessage(db *gorm.DB, senderId uint64, toUserId uint64, location *proto_build.LocationMessage) (*MessageContent, error) {
	body, _ := json.Marshal(location)
	messageContent := &MessageContent{
		ID:    0,
		Type:  constant.MessageTypeLocation,
		Body:  string(body),
		Extra: "",
	}
	err := db.Create(messageContent).Error
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return messageContent, nil
}

func CreateUserFaceMessage(db *gorm.DB, senderId uint64, toUserId uint64, face *proto_build.FaceMessage) (*MessageContent, error) {
	body, _ := json.Marshal(face)
	messageContent := &MessageContent{
		ID:    0,
		Type:  constant.MessageTypeFace,
		Body:  string(body),
		Extra: "",
	}
	err := db.Create(messageContent).Error
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return messageContent, nil
}

func CreateUserSoundMessage(db *gorm.DB, senderId uint64, toUserId uint64, sound *proto_build.SoundMessage) (*MessageContent, error) {
	body, _ := json.Marshal(sound)
	messageContent := &MessageContent{
		ID:    0,
		Type:  constant.MessageTypeSound,
		Body:  string(body),
		Extra: "",
	}
	err := db.Create(messageContent).Error
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return messageContent, nil
}

func CreateUserVideoMessage(db *gorm.DB, senderId uint64, toUserId uint64, video *proto_build.VideoMessage) (*MessageContent, error) {
	body, _ := json.Marshal(video)
	messageContent := &MessageContent{
		ID:    0,
		Type:  constant.MessageTypeVideo,
		Body:  string(body),
		Extra: "",
	}
	err := db.Create(messageContent).Error
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return messageContent, nil
}

func CreateUserImageMessage(db *gorm.DB, senderId uint64, toUserId uint64, image *proto_build.ImageMessageItem) (*MessageContent, error) {
	body, _ := json.Marshal(image)
	messageContent := &MessageContent{
		ID:    0,
		Type:  constant.MessageTypeImage,
		Body:  string(body),
		Extra: "",
	}
	err := db.Create(messageContent).Error
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return messageContent, nil
}

func CreateUserFileMessage(db *gorm.DB, senderId uint64, toUserId uint64, file *proto_build.FileMessage) (*MessageContent, error) {
	body, _ := json.Marshal(file)
	messageContent := &MessageContent{
		ID:    0,
		Type:  constant.MessageTypeFile,
		Body:  string(body),
		Extra: "",
	}
	err := db.Create(messageContent).Error
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return messageContent, nil
}

func CreateGroupTextMessage(db *gorm.DB, senderId uint64, toGroupId uint64, atUserId []uint64, text *proto_build.TextMessage) (*MessageContent, error) {
	body, _ := json.Marshal(text)
	messageContent := &MessageContent{
		ID:    0,
		Type:  constant.MessageTypeText,
		Body:  string(body),
		Extra: "",
	}
	err := db.Create(messageContent).Error
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return messageContent, nil
}

func CreateGroupLocationMessage(db *gorm.DB, senderId uint64, toGroupId uint64, atUserId []uint64, location *proto_build.LocationMessage) (*MessageContent, error) {
	body, _ := json.Marshal(location)
	messageContent := &MessageContent{
		ID:    0,
		Type:  constant.MessageTypeLocation,
		Body:  string(body),
		Extra: "",
	}
	err := db.Create(messageContent).Error
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return messageContent, nil
}

func CreateGroupFaceMessage(db *gorm.DB, senderId uint64, toGroupId uint64, atUserId []uint64, face *proto_build.FaceMessage) (*MessageContent, error) {
	body, _ := json.Marshal(face)
	messageContent := &MessageContent{
		ID:    0,
		Type:  constant.MessageTypeFace,
		Body:  string(body),
		Extra: "",
	}
	err := db.Create(messageContent).Error
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return messageContent, nil
}

func CreateGroupSoundMessage(db *gorm.DB, senderId uint64, toGroupId uint64, atUserId []uint64, sound *proto_build.SoundMessage) (*MessageContent, error) {
	body, _ := json.Marshal(sound)
	messageContent := &MessageContent{
		ID:    0,
		Type:  constant.MessageTypeSound,
		Body:  string(body),
		Extra: "",
	}
	err := db.Create(messageContent).Error
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return messageContent, nil
}

func CreateGroupVideoMessage(db *gorm.DB, senderId uint64, toGroupId uint64, atUserId []uint64, video *proto_build.VideoMessage) (*MessageContent, error) {
	body, _ := json.Marshal(video)
	messageContent := &MessageContent{
		ID:    0,
		Type:  constant.MessageTypeVideo,
		Body:  string(body),
		Extra: "",
	}
	err := db.Create(messageContent).Error
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return messageContent, nil
}

func CreateGroupImageMessage(db *gorm.DB, senderId uint64, toGroupId uint64, atUserId []uint64, image *proto_build.ImageMessageItem) (*MessageContent, error) {
	body, _ := json.Marshal(image)
	messageContent := &MessageContent{
		ID:    0,
		Type:  constant.MessageTypeImage,
		Body:  string(body),
		Extra: "",
	}
	err := db.Create(messageContent).Error
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return messageContent, nil
}

func CreateGroupFileMessage(db *gorm.DB, senderId uint64, toGroupId uint64, atUserId []uint64, file *proto_build.FileMessage) (*MessageContent, error) {
	body, _ := json.Marshal(file)
	messageContent := &MessageContent{
		ID:    0,
		Type:  constant.MessageTypeFile,
		Body:  string(body),
		Extra: "",
	}
	err := db.Create(messageContent).Error
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return messageContent, nil
}
