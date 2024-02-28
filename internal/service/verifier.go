package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Verify
//	@Tags		验证器
//	@Summary	验证某条性质是否成立
//	@Param		property	formData	string	false	"property""
//	@Success	200			{string}	json	"{"code":"200","data":""}"
//	@Router		/verify [post]
func Verify(c *gin.Context) {
	property := c.PostForm("property")
	if property == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "性质不能为空",
		})
		return
	}
	data := check(property)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": data,
	})
}

func check(property string) float64 {
	// TODO: 检查性质是否成立，并返回性质成立的比例
	return float64(len(property))
}
