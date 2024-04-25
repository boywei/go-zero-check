package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 响应
type Response struct {
	Code    int         `json:"code,omitempty" example:"200"`                     // 错误码
	Message string      `json:"message,omitempty" example:"SUCCESS"`              // 消息
	Data    interface{} `json:"data,omitempty" swaggertype:"string" example:"{}"` // 数据
}

// New 响应
func New(errCode *ErrCode, data interface{}) *Response {
	return &Response{
		Code:    errCode.Code,
		Message: errCode.Message,
		Data:    data,
	}
}

// Success 请求成功
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, New(SuccessResp, data))
}

// Failure 请求失败
func Failure(c *gin.Context, errCode *ErrCode) {
	c.JSON(http.StatusBadRequest, errCode)
}
