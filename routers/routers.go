package routers

import (
	"api/pkg/setting"
	v1 "api/routers/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()        //创建gin框架路由实例
	r.Use(gin.Logger())   //使用gin框架中的打印中间件
	r.Use(gin.Recovery()) //使用gin框架中的恢复中间件，可以从任何恐慌中恢复，如果有，则写入500
	gin.SetMode(setting.ServerSetting.RunMode)
	apiv1 := r.Group("api/v1")
	{
		apiv1.GET("version", v1.GetAppVersionTest)
		// apiv1.POST("auth", v1.AuthStore)
	}
	return r
}
