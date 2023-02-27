package api

import (
	"github.com/dotdancer/gogofly/service"
	"github.com/dotdancer/gogofly/service/dto"
	"github.com/dotdancer/gogofly/utils"
	"github.com/gin-gonic/gin"
)

const (
	ERR_CODE_ADD_USER       = 10011
	ERR_CODE_GET_USER_BY_ID = 10012
	ERR_CODE_GET_USER_LIST  = 10013
	ERR_CODE_UPDATE_USER    = 10014
)

type UserApi struct {
	BaseApi
	Service *service.UserService
}

func NewUserApi() UserApi {
	return UserApi{
		BaseApi: NewBaseApi(),
		Service: service.NewUserService(),
	}
}

// @Tag 用户管理
// @Summary 用户登录
// @Description 用户登录详情描述
// @Param name formData string true "用户名"
// @Param password formData string true "密码"
// @Success 200 {string} string "登录成功"
// @Failure 401 {string} string "登录失败"
// @Router /api/v1/public/user/login [post]
func (m UserApi) Login(c *gin.Context) {
	var iUserLoginDTO dto.UserLoginDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iUserLoginDTO}).GetError(); err != nil {
		return
	}

	iUser, err := m.Service.Login(iUserLoginDTO)
	if err != nil {
		m.Fail(ResponseJson{
			Msg: err.Error(),
		})
		return
	}

	token, _ := utils.GenerateToken(iUser.ID, iUser.Name)

	m.OK(ResponseJson{
		Data: gin.H{
			"token": token,
			"user":  iUser,
		},
	})

	//fmt.Println("Login 执行了")
	//c.JSON(http.StatusOK, gin.H{
	//	"msg": "Login Success",
	//})

	//OK(c, ResponseJson{
	//	Msg: "Login Success",
	//})

	//Fail(c, ResponseJson{
	//	Code: 9001,
	//	Msg:  "Login Failed",
	//})
}

func (m UserApi) AddUser(c *gin.Context) {
	var iUserAddDTO dto.UserAddDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iUserAddDTO}).GetError(); err != nil {
		return
	}

	err := m.Service.AddUser(&iUserAddDTO)

	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ERR_CODE_ADD_USER,
			Msg:  err.Error(),
		})

		return
	}

	m.OK(ResponseJson{
		Data: iUserAddDTO,
	})
}

func (m UserApi) GetUserById(c *gin.Context) {
	var iCommonIDDTO dto.CommonIDDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iCommonIDDTO, BindParamsFromUri: true}).GetError(); err != nil {
		return
	}

	iUser, err := m.Service.GetUserById(&iCommonIDDTO)
	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ERR_CODE_GET_USER_BY_ID,
			Msg:  err.Error(),
		})

		return
	}

	m.OK(ResponseJson{
		Data: iUser,
	})
}

func (m UserApi) GetUserList(c *gin.Context) {
	var iUserListDTO dto.UserListDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iUserListDTO}).GetError(); err != nil {
		return
	}

	giUserList, nTotal, err := m.Service.GetUserList(&iUserListDTO)

	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ERR_CODE_GET_USER_LIST,
			Msg:  err.Error(),
		})
		return
	}

	m.OK(ResponseJson{
		Data:  giUserList,
		Total: nTotal,
	})
}

func (m UserApi) UpdateUser(c *gin.Context) {
	var iUserUpdateDTO dto.UserUpdateDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iUserUpdateDTO}).GetError(); err != nil {
		return
	}

	err := m.Service.UpdateUser(&iUserUpdateDTO)

	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ERR_CODE_UPDATE_USER,
			Msg:  err.Error(),
		})
		return
	}

	m.OK(ResponseJson{})
}
