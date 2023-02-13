package api

import (
	"github.com/gin-gonic/gin"
)

type UserApi struct {
}

func NewUserApi() UserApi {
	return UserApi{}
}

// @Tag 用户管理
// @Summary 用户登录
// @Description 用户登录详情描述
// @Param name formData string true "用户名"
// @Param password formData string true "密码"
// @Success 200 {string} string "登录成功"
// @Failure 401 {string} string "登录失败"
// @Router /api/v1/public/user/login [post]
func (m UserApi) Login(ctx *gin.Context) {
	//fmt.Println("Login 执行了")
	//ctx.JSON(http.StatusOK, gin.H{
	//	"msg": "Login Success",
	//})

	//OK(ctx, ResponseJson{
	//	Msg: "Login Success",
	//})

	Fail(ctx, ResponseJson{
		Code: 9001,
		Msg:  "Login Failed",
	})
}
