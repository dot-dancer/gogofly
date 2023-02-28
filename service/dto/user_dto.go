package dto

import "github.com/dotdancer/gogofly/model"

type UserLoginDTO struct {
	Name     string `json:"name" binding:"required" message:"用户名不能为空"`
	Password string `json:"password" binding:"required" message:"密码不能为空"`
}

// ===============================================================================
// = 添加用户相关DTO
type UserAddDTO struct {
	ID       uint
	Name     string `json:"name" form:"name" binding:"required" message:"用户名不能为空"`
	RealName string `json:"real_name" form:"real_name"`
	Avatar   string `json:"avatar"`
	Mobile   string `json:"mobile" form:"mobile"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password,omitempty" form:"password" binding:"required" message:"密码不能为空"`
}

func (m *UserAddDTO) ConvertToModel(iUser *model.User) {
	iUser.Name = m.Name
	iUser.RealName = m.RealName
	iUser.Avatar = m.Avatar
	iUser.Mobile = m.Mobile
	iUser.Email = m.Email
	iUser.Password = m.Password
}

// ===============================================================================
// = 更新用户相关DTO
type UserUpdateDTO struct {
	ID       uint   `json:"id" form:"id" uri:"id"`
	Name     string `json:"name" form:"name"`
	RealName string `json:"real_name" form:"real_name"`
	Mobile   string `json:"mobile" form:"mobile"`
	Email    string `json:"email" form:"email"`
}

func (m *UserUpdateDTO) ConvertToModel(iUser *model.User) {
	iUser.ID = m.ID
	iUser.Name = m.Name
	iUser.RealName = m.RealName
	iUser.Mobile = m.Mobile
	iUser.Email = m.Email
}

// ===============================================================================
// = 用户列表相关DTO
type UserListDTO struct {
	Paginate
}
