package dao

import "github.com/dotdancer/gogofly/model"

var userDao *UserDao

type UserDao struct {
	BaseDao
}

func NewUserDao() *UserDao {
	if userDao == nil {
		userDao = &UserDao{NewBaseDao()}
	}

	return userDao
}

func (m *UserDao) GetUserByNameAndPassword(stUserName, stPassword string) model.User {
	var iUser model.User
	m.Orm.Model(&iUser).Where("name=? and password=?", stUserName, stPassword).Find(&iUser)
	return iUser
}
