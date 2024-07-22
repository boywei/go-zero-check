package response

import (
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 响应
type Response struct {
	Code    int         `json:"code,omitempty" example:"200"`                     // 错误码
	Message string      `json:"message,omitempty" example:"SUCCESS"`              // 消息
	Data    interface{} `json:"data,omitempty" swaggertype:"string" example:"{}"` // 数据
}

// Success 请求成功
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{Code: 200, Message: "SUCCESS", Data: data})
}

// RequestError 请求失败
func RequestError(c *gin.Context, errorMsg string) {
	log.Errorln(errorMsg)
	c.JSON(http.StatusBadRequest, Response{Code: -1, Message: errorMsg, Data: nil})
}

// ServiceError 业务错误
func ServiceError(c *gin.Context, errorMsg string) {
	log.Errorln(errorMsg)
	c.JSON(http.StatusOK, Response{Code: -1, Message: errorMsg, Data: nil})
}
