package service

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ConvertLustre
//
//	@Tags		Lustre方法
//	@Summary	Lustre转Uppaal
//	@Param		file	formData	string	false	"file path"
//	@Success	200		{string}	json	"{"code":"200","data":""}"
//	@Router		/convert-lustre [post]
func ConvertLustre(c *gin.Context) {
	file := c.PostForm("file")
	if file == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "文件不能为空",
		})
		return
	}
	// TODO: 解析Lustre -> Uppaal
	/*
		解析步骤：
		1. 映射为json模型文件
		2. 调用Convert接口
	*/
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": fmt.Sprintf("converting %s", file),
	})
}
