package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name     string `gorm:"column:name;type:varchar(50)" json:"name"`
	Descript string `gorm:"column:descript;type:varchar(200)" json:"descript"`
	Users    []User `gorm:"many2many:user_role" json:"users"`
}

func (table *Role) TableName() string {
	return "role"
}
