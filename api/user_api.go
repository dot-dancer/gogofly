package api

import (
	"github.com/dotdancer/gogofly/service"
	"github.com/dotdancer/gogofly/service/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	ERR_CODE_ADD_USER       = 10011
	ERR_CODE_GET_USER_BY_ID = 10012
	ERR_CODE_GET_USER_LIST  = 10013
	ERR_CODE_UPDATE_USER    = 10014
	ERR_CODE_DELETE_USER    = 10015
	ERR_CODE_LOGIN          = 10016
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

	iUser, token, err := m.Service.Login(iUserLoginDTO)
	if err == nil {
		err = service.SetLoginUserTokenToRedis(iUser.ID, token)
	}

	if err != nil {
		m.Fail(ResponseJson{
			Status: http.StatusUnauthorized,
			Code:   ERR_CODE_LOGIN,
			Msg:    err.Error(),
		})
		return
	}

	m.OK(ResponseJson{
		Data: gin.H{
			"token": token,
			"user":  iUser,
		},
	})
}

func (m UserApi) AddUser(c *gin.Context) {
	var iUserAddDTO dto.UserAddDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iUserAddDTO}).GetError(); err != nil {
		return
	}

	//file, _ := c.FormFile("file")
	//stFilePath := fmt.Sprintf("./upload/%s", file.Filename)
	//_ = c.SaveUploadedFile(file, stFilePath)
	//iUserAddDTO.Avatar = stFilePath

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
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iCommonIDDTO, BindUri: true}).GetError(); err != nil {
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
	//strId := c.Param("id")
	//fmt.Println("strId:" + strId)
	//
	//id, _ := strconv.Atoi(strId)
	//uid := uint(id)
	//iUserUpdateDTO.ID = uid

	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iUserUpdateDTO, BindAll: true}).GetError(); err != nil {
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

func (m UserApi) DeleteUserById(c *gin.Context) {
	var iCommonIDDTO dto.CommonIDDTO
	if err := m.BuildRequest(BuildRequestOption{Ctx: c, DTO: &iCommonIDDTO, BindUri: true}).GetError(); err != nil {
		return
	}

	err := m.Service.DeleteUserById(&iCommonIDDTO)
	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ERR_CODE_DELETE_USER,
			Msg:  err.Error(),
		})
		return
	}

	m.OK(ResponseJson{})
}
