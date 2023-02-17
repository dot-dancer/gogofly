package dto

type UserLoginDTO struct {
	Name     string `json:"name" binding:"required,email,first_is_a" message:"用户名填写错误" required_err:"用户名不能为空"`
	Password string `json:"password" binding:"required" message:"密码不能为空"`
}
