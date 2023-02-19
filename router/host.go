package router

import (
	"github.com/dotdancer/gogofly/api"
	"github.com/gin-gonic/gin"
)

func InitHostRoutes() {
	RegistRoute(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		hostApi := api.NewHostApi()
		rgAuthUser := rgAuth.Group("host")
		{
			rgAuthUser.POST("/shutdown", hostApi.Shutdown)
		}
	})
}
