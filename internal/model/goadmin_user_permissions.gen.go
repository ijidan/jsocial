// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameGoadminUserPermission = "goadmin_user_permissions"

// GoadminUserPermission mapped from table <goadmin_user_permissions>
type GoadminUserPermission struct {
	UserID       int32     `gorm:"column:user_id;primaryKey" json:"user_id"`
	PermissionID int32     `gorm:"column:permission_id;primaryKey" json:"permission_id"`
	CreatedAt    time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName GoadminUserPermission's table name
func (*GoadminUserPermission) TableName() string {
	return TableNameGoadminUserPermission
}
