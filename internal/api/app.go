package api

import (
	"github.com/boywei/go-zero-check/docs"
	"github.com/boywei/go-zero-check/internal/service"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()

	// Swagger 配置
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// json
	r.POST("/convert", service.Convert)

	// lustre
	r.POST("/convert-lustre", service.ConvertLustre)

	// 模拟器
	r.POST("/next", service.Next)
	r.GET("/get-next", service.GetNext)
	r.POST("/reset", service.Reset)
	r.GET("/get-trace", service.GetTrace)
	r.POST("/save-trace", service.SaveTrace)
	r.POST("/open-trace", service.OpenTrace)
	r.POST("/random-trace", service.RandomTrace)
	r.GET("/get-global", service.GetGlobal)
	r.GET("/get-local", service.GetLocal)
	r.GET("/get-current-status", service.GetCurrentStatus)
	r.GET("/get-sync", service.GetSync)

	// 验证器
	r.POST("/verify", service.Verify)

	return r
}
