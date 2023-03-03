package router

import (
	"github.com/dotdancer/gogofly/api"
	"github.com/gin-gonic/gin"
)

func InitUserRoutes() {
	RegistRoute(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		userApi := api.NewUserApi()
		rgPublicUser := rgPublic.Group("user").Use(func() gin.HandlerFunc {
			return func(ctx *gin.Context) {
				//ctx.AbortWithStatusJSON(200, gin.H{
				//	"msg": "Login MiddleWare",
				//})
			}
		}())
		{
			rgPublicUser.POST("/login", userApi.Login)
		}

		rgAuthUser := rgAuth.Group("user")
		{
			rgAuthUser.POST("", userApi.AddUser)
			rgAuthUser.POST("/list", userApi.GetUserList)
			rgAuthUser.GET("/:id", userApi.GetUserById)
			rgAuthUser.PUT("/:id", userApi.UpdateUser)
			rgAuthUser.DELETE("/:id", userApi.DeleteUserById)
		}
	})
}
