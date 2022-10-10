package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserID   int64  `gorm:"column:user_id;type:bigint(20)" json:"user_id"`     // 用户id
	Username string `gorm:"column:username;type:varchar(50);" json:"username"` // 用户名
	Password string `gorm:"column:password;type:varchar(50);" json:"password"` // 密码
	// UUID     string `gorm:"column:uuid;type:varchar(36);" json:"uudi"`         // 唯一标识
	Roles []Role `gorm:"many2many:user_role" json:"roles"`
}

func (table *User) TableName() string {
	return "user"
}
