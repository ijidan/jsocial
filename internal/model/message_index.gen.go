// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameMessageIndex = "message_index"

// MessageIndex mapped from table <message_index>
type MessageIndex struct {
	ID               int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	SenderID         int64          `gorm:"column:sender_id;not null" json:"sender_id"`                   // 发送人ID
	ReceiverID       int64          `gorm:"column:receiver_id;not null" json:"receiver_id"`               // 接收人ID
	MessageContentID int64          `gorm:"column:message_content_id;not null" json:"message_content_id"` // 消息ID
	GroupID          int64          `gorm:"column:group_id;not null" json:"group_id"`                     // 群ID
	Status           int32          `gorm:"column:status;not null;default:1" json:"status"`               // 状态：1-待送达，2-已送到，3-已确认，0-已删除
	CreatedAt        time.Time      `gorm:"column:created_at" json:"created_at"`                          // 创建时间
	UpdatedAt        time.Time      `gorm:"column:updated_at" json:"updated_at"`                          // 更新时间
	DeletedAt        gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`                          // 删除时间
}

// TableName MessageIndex's table name
func (*MessageIndex) TableName() string {
	return TableNameMessageIndex
}
