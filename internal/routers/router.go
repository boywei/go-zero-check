package routers

import (
	"github.com/boywei/go-zero-check/internal/middleware"
	"net/http"
	"runtime/debug"

	"github.com/boywei/go-zero-check/docs"
	"github.com/boywei/go-zero-check/internal/handler"
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
	r.Use(middleware.Cors())

	// Swagger 配置
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// 编辑器
	model := r.Group("/model")
	{
		model.POST("/convert", handler.Convert)
		model.POST("/delete", handler.DeleteModel)
	}

	// lustre
	lustre := r.Group("/lustre")
	{
		lustre.POST("/convert", handler.ConvertLustre)
		lustre.POST("/check-dataflow", handler.CheckDataflow)
	}

	// 模拟器
	simulator := r.Group("/simulator")
	{
		simulator.GET("/start", handler.Start)
		simulator.POST("/step", handler.Step)
		simulator.GET("/get-next", handler.GetNext)
		simulator.POST("/reset", handler.Reset)
		simulator.GET("/get-trace", handler.GetTrace)
		simulator.POST("/save-trace", handler.SaveTrace)
		simulator.POST("/open-trace", handler.OpenTrace)
		simulator.POST("/random-trace", handler.RandomTrace)
		simulator.GET("/get-global", handler.GetGlobal)
		simulator.GET("/get-local", handler.GetLocal)
		simulator.GET("/get-current-status", handler.GetCurrentStatus)
		simulator.GET("/get-sync", handler.GetSync)
	}

	// 验证器
	r.POST("/verify", handler.Verify)

	return r
}

// HandleNotFound 404
func HandleNotFound(c *gin.Context) {
	c.IndentedJSON(http.StatusNotFound, gin.H{
		"msg": "not found",
	})
}

// Recover 500
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
