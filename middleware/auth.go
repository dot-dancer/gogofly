package middleware

import (
	"fmt"
	"github.com/dotdancer/gogofly/api"
	"github.com/dotdancer/gogofly/global"
	"github.com/dotdancer/gogofly/global/constants"
	"github.com/dotdancer/gogofly/model"
	"github.com/dotdancer/gogofly/service"
	"github.com/dotdancer/gogofly/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	ERR_CODE_INVALID_TOKEN_INVALID     = 10401
	ERR_CODE_INVALID_TOKEN_PARSE_ERROR = 10402
	ERR_CODE_INVALID_TOKEN_NOT_MATCHED = 10403
	ERR_CODE_INVALID_TOKEN_EXPIRED     = 10404
	ERR_CODE_INVALID_TOKEN_RENEW_ERROR = 10405
	TOKEN_NAME                         = "Authorization"
	TOKEN_PREFIX                       = "Bearer: "
	RENEW_TOKEN_DURATION               = 10 * 60 * time.Second
)

func tokenErr(c *gin.Context, code int) {
	api.Fail(c, api.ResponseJson{
		Status: http.StatusUnauthorized,
		Code:   code,
		Msg:    "Invalid Token",
	})
}

func Auth() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.GetHeader(TOKEN_NAME)

		// Token不存在, 直接返回
		if token == "" || !strings.HasPrefix(token, TOKEN_PREFIX) {
			tokenErr(c, ERR_CODE_INVALID_TOKEN_INVALID)
			return
		}

		// Token无法解析, 直接返回
		token = token[len(TOKEN_PREFIX):]
		iJwtCustClaims, err := utils.ParseToken(token)
		nUserId := iJwtCustClaims.ID
		if err != nil || nUserId == 0 {
			fmt.Println(err.Error())
			tokenErr(c, ERR_CODE_INVALID_TOKEN_PARSE_ERROR)
			return
		}

		stUserId := strconv.Itoa(int(nUserId))
		stRedisUserIdKey := strings.Replace(constants.LOGIN_USER_TOKEN_REDIS_KEY, "{ID}", stUserId, -1)

		// Token与访问者登录对应的token不一致, 直接返回
		stRedisToken, err := global.RedisClient.Get(stRedisUserIdKey)
		if err != nil || token != stRedisToken {
			tokenErr(c, ERR_CODE_INVALID_TOKEN_NOT_MATCHED)
			return
		}

		// Token已过期, 直接返回
		nTokenExpireDuration, err := global.RedisClient.GetExpireDuration(stRedisUserIdKey)
		if err != nil || nTokenExpireDuration <= 0 {
			tokenErr(c, ERR_CODE_INVALID_TOKEN_EXPIRED)
			return
		}

		// Token的续期
		if nTokenExpireDuration.Seconds() < RENEW_TOKEN_DURATION.Seconds() {
			err := service.SetLoginUserTokenToRedis(nUserId, token)
			if err != nil {
				tokenErr(c, ERR_CODE_INVALID_TOKEN_RENEW_ERROR)
				return
			}
		}

		// 将用户信息存入上下文, 方便后续处理继续使用
		//iUser, err := dao.NewUserDao().GetUserById(nUserId)
		//if err != nil {
		//	tokenErr(c)
		//	return
		//}
		//c.Set(constants.LOGIN_USER, iUser)
		c.Set(constants.LOGIN_USER, model.LoginUser{
			ID:   nUserId,
			Name: iJwtCustClaims.Name,
		})
		c.Next()
	}
}
