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

	return r
}
