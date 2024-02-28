package service

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Next 使能迁移中的下一步
//	@Tags		模拟器
//	@Summary	使能迁移下一步
//	@Param		id	formData	string	false	"model id"
//	@Success	200	{string}	json	"{"code":"200","data":""}"
//	@Router		/next [post]
func Next(c *gin.Context) {
	id := c.PostForm("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "id不能为空",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": fmt.Sprintf("next %s", id),
	})
}

// GetNext 获取使能迁移中所有可能的下一步
//	@Tags		模拟器
//	@Summary	获取使能迁移中所有可能的下一步
//	@Param		id	query		string	false	"model id"
//	@Success	200	{string}	json	"{"code":"200","data":""}"
//	@Router		/get-next [get]
func GetNext(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": "next",
	})
}

// Reset 重置使能迁移, 复位
//	@Tags		模拟器
//	@Summary	重置使能迁移, 复位
//	@Param		id	formData	string	false	"model id"
//	@Success	200	{string}	json	"{"code":"200","data":""}"
//	@Router		/reset [post]
func Reset(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": "reset",
	})
}

// GetTrace 获取模拟Trace
//	@Tags		模拟器
//	@Summary	获取模拟Trace
//	@Param		id	query		string	false	"model id"
//	@Success	200	{string}	json	"{"code":"200","data":""}"
//	@Router		/get-trace [get]
func GetTrace(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": "trace",
	})
}

// SaveTrace 保存模拟Trace
//	@Tags		模拟器
//	@Summary	保存模拟Trace
//	@Param		id	formData	string	false	"model id"
//	@Success	200	{string}	json	"{"code":"200","data":""}"
//	@Router		/save-trace [post]
func SaveTrace(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": "save trace",
	})
}

// OpenTrace 打开模拟Trace
//	@Tags		模拟器
//	@Summary	打开模拟Trace
//	@Param		file	formData	string	false	"trace path"
//	@Success	200		{string}	json	"{"code":"200","data":""}"
//	@Router		/open-trace [post]
func OpenTrace(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": "open trace",
	})
}

// RandomTrace 随机模拟Trace
//	@Tags		模拟器
//	@Summary	随机模拟Trace
//	@Param		id	formData	string	false	"model id"
//	@Success	200	{string}	json	"{"code":"200","data":""}"
//	@Router		/random-trace [post]
func RandomTrace(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": "random trace",
	})
}

// GetGlobal 获取全局变量(对应模拟器中间的全局变量)
//	@Tags		模拟器
//	@Summary	获取全局变量(对应模拟器中间的全局变量)
//	@Param		id	query		string	false	"model id"
//	@Success	200	{string}	json	"{"code":"200","data":""}"
//	@Router		/get-global [get]
func GetGlobal(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": "globals",
	})
}

// GetLocal 获取各自动机的局部变量(对应模拟器中间的局部变量)
//	@Tags		模拟器
//	@Summary	获取各自动机的局部变量(对应模拟器中间的局部变量)
//	@Param		id	query		string	false	"model id"
//	@Success	200	{string}	json	"{"code":"200","data":""}"
//	@Router		/get-global [get]
func GetLocal(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": "locals",
	})
}

// GetCurrentStatus 获取当前状态(对应模拟器右上角的各自动机状态图)
//	@Tags		模拟器
//	@Summary	获取当前状态
//	@Param		id	query		string	false	"model id"
//	@Success	200	{string}	json	"{"code":"200","data":""}"
//	@Router		/get-current-status [get]
func GetCurrentStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": "current status",
	})
}

// GetSync 获取同步情况(对应模拟器右下角的同步图)
//	@Tags		模拟器
//	@Summary	获取同步情况(对应模拟器右下角的同步图)
//	@Param		id	query		string	false	"model id"
//	@Success	200	{string}	json	"{"code":"200","data":""}"
//	@Router		/get-sync [get]
func GetSync(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": "sync",
	})
}
