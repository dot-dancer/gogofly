package model

import (
	"github.com/dotdancer/gogofly/utils"
	"gorm.io/gorm"
)

// ===============================================================================
// = 用户信息
type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"size:64;not null"`
	RealName string `json:"real_name" gorm:"size:128"`
	Avatar   string `json:"avatar" gorm:"size:255"`
	Mobile   string `json:"mobile" gorm:"size:11"`
	Email    string `json:"email" gorm:"size:128"`
	Password string `json:"-" gorm:"128;not null"`
}

func (m *User) Encrypt() error {
	stHash, err := utils.Encrypt(m.Password)
	if err == nil {
		m.Password = stHash
	}

	return err
}

func (m *User) BeforeCreate(orm *gorm.DB) error {
	return m.Encrypt()
}

// ===============================================================================
// = 用户登录信息
type LoginUser struct {
	ID   uint
	Name string
}
