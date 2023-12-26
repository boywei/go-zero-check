package service

import (
	"fmt"
	"github.com/boywei/go-zero-check/internal/util/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Convert
// @Tags 转换方法
// @Summary 前端传过来的JSON转Go对象
// @Param file formData string false "json file path"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /convert [post]
func Convert(c *gin.Context) {
	file := c.PostForm("file")
	if file == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "文件不能为空",
		})
		return
	}
	object, err := json.ConvertJson2Uppaal(file)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  fmt.Sprintf("JSON decode error! %v", err),
		})
		return
	}
	// TODO: 处理转化后的object
	data := object
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": data,
	})
}
