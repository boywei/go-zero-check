package routers

import (
	"net/http"
	"runtime/debug"

	"github.com/boywei/go-zero-check/docs"
	"github.com/boywei/go-zero-check/internal/service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.NoRoute(HandleNotFound)
	r.NoMethod(HandleNotFound)
	r.Use(Recover)

	// Swagger 配置
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// json
	model := r.Group("/model")
	{
		model.POST("/convert", service.Convert)
		model.POST("/delete", service.DeleteModel)
		model.GET("/test", service.TestModel)
		model.POST("/run", service.RunModel)
	}

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

// 404
func HandleNotFound(c *gin.Context) {
	c.IndentedJSON(http.StatusNotFound, gin.H{
		"msg": "not found",
	})
}

// 500
func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			//打印错误堆栈信息
			log.Printf("panic: %v\n", r)
			debug.PrintStack()
			c.IndentedJSON(http.StatusNotFound, gin.H{
				"msg": "服务器内部错误",
			})
		}
	}()
	//继续后续接口调用
	c.Next()
}
