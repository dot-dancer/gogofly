package dao

import (
	"github.com/dotdancer/gogofly/model"
	"github.com/dotdancer/gogofly/service/dto"
)

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

func (m *UserDao) GetUserByName(stUserName string) (model.User, error) {
	var iUser model.User
	err := m.Orm.Model(&iUser).Where("name = ?", stUserName).Find(&iUser).Error

	return iUser, err
}

func (m *UserDao) GetUserByNameAndPassword(stUserName, stPassword string) model.User {
	var iUser model.User
	m.Orm.Model(&iUser).Where("name=? and password=?", stUserName, stPassword).Find(&iUser)
	return iUser
}

func (m *UserDao) CheckUserNameExist(stUserName string) bool {
	var nTotal int64
	m.Orm.Model(&model.User{}).Where("name = ? ", stUserName).
		Count(&nTotal)

	return nTotal > 0
}

func (m *UserDao) AddUser(iUserAddDTO *dto.UserAddDTO) error {
	var iUser model.User
	iUserAddDTO.ConvertToModel(&iUser)

	err := m.Orm.Save(&iUser).Error
	if err == nil {
		iUserAddDTO.ID = iUser.ID
		iUserAddDTO.Password = ""
	}

	return err
}

func (m *UserDao) GetUserById(id uint) (model.User, error) {
	var iUser model.User
	err := m.Orm.First(&iUser, id).Error
	return iUser, err
}

func (m *UserDao) GetUserList(iUserListDTO *dto.UserListDTO) ([]model.User, int64, error) {
	var giUserList []model.User
	var nTotal int64

	err := m.Orm.Model(&model.User{}).
		Scopes(Paginate(iUserListDTO.Paginate)).
		Find(&giUserList).
		Offset(-1).Limit(-1).
		Count(&nTotal).
		Error

	return giUserList, nTotal, err
}

func (m *UserDao) UpdateUser(iUserUpdateDTO *dto.UserUpdateDTO) error {
	var iUser model.User

	m.Orm.First(&iUser, iUserUpdateDTO.ID)
	iUserUpdateDTO.ConvertToModel(&iUser)

	return m.Orm.Save(&iUser).Error
}

func (m *UserDao) DeleteUserById(id uint) error {
	return m.Orm.Delete(&model.User{}, id).Error
}
