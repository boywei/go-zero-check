package response

import (
	"net/http"

	"github.com/boywei/go-zero-check/internal/util/enum"
	"github.com/gin-gonic/gin"
)

// Response 响应
type Response struct {
	Code    int         `json:"Code,omitempty"`    // 错误码
	Message string      `json:"Message,omitempty"` // 消息
	Data    interface{} `json:"Data,omitempty"`    // 数据
}

// New 响应
func New(errCode *enum.ErrCode, data interface{}) *Response {
	return &Response{
		Code:    errCode.Code,
		Message: errCode.Message,
		Data:    data,
	}
}

// Success 请求成功
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, New(enum.Success, data))
}

// Failure 请求失败
func Failure(c *gin.Context, errCode *enum.ErrCode) {
	c.JSON(http.StatusBadRequest, errCode)
}
