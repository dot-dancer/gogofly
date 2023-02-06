package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"size:64;not null"`
	RealName string `gorm:"size:128"`
	Avatar   string `gorm:"size:255"`
	Mobile   string `gorm:"size:11"`
	Email    string `gorm:"size:128"`
	Password string `gorm:"128;not null"`
}
