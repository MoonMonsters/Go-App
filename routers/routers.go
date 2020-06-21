package routers

import (
	"Go-App/middleware"
	"Go-App/pkg/setting"
	v1 "Go-App/routers/v1"
	"github.com/gin-gonic/gin"
)

/**
初始化路由
*/
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.ServerSetting.RunMode)

	apiV1 := r.Group("/api/v1")
	{
		apiV1.POST("auth", v1.UserStore)
		apiV1Token := apiV1.Group("token/")
		apiV1Token.Use(middleware.TokenVer())
		{
			apiV1Token.GET("version", v1.GetAppVersionTest)
		}
	}

	return r
}
