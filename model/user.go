package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"size:64;not null"`
	RealName string `json:"real_name" gorm:"size:128"`
	Avatar   string `json:"avatar" gorm:"size:255"`
	Mobile   string `json:"mobile" gorm:"size:11"`
	Email    string `json:"email" gorm:"size:128"`
	Password string `json:"-" gorm:"128;not null"`
}
