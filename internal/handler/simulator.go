package handler

import (
	"fmt"
	"net/http"

	"github.com/boywei/go-zero-check/internal/util/response"
	"github.com/gin-gonic/gin"
)

// Start 获取使能迁移中所有可能的下一步
//
//	@Tags		模拟器
//	@Summary	获取使能迁移中所有可能的下一步
//	@Param		id	query	string	false	"model id"
//	@Produce	json
//	@Success	200	{object}	response.Response
//	@Router		/simulator/start [get]
func Start(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": "next",
	})
}

// Step 使能迁移中的下一步
//
//	@Tags		模拟器
//	@Summary	使能迁移下一步，根据自动机的id/transition的id/select的参数使模型步进一次，并返回步进结果和下一次可步进的自动机
//	@Param		id		body	string	true	"automaton's id"
//	@Param		param	body	string	false	"parameters of the automaton"
//	@Produce	json
//	@Success	200	{object}	response.Response
//	@Router		/simulator/step [post]
func Step(c *gin.Context) {
	content := make(map[string]string)
	c.BindJSON(&content)
	id := content["id"]
	if id == "" {
		response.ServiceError(c, "模型id不能为空")
		return
	}
	// 先选择自动机，再选择该自动机可选的参数，再运行得到下一步；
	// 运行时需要考虑信号和变量对其他自动机的影响，即需要逐个检查各个自动机location对应的下一条边sync和guard的情况，获取下一步可以主动迁移的自动机；
	// 循环以上过程
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": fmt.Sprintf("next %s", id),
	})
}

// GetNext 获取使能迁移中所有可能的下一步
//
//	@Tags		模拟器
//	@Summary	获取使能迁移中所有可能的下一步
//	@Param		id	query	string	false	"model id"
//	@Produce	json
//	@Success	200	{object}	response.Response
//	@Router		/simulator/get-next [get]
func GetNext(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": "next",
	})
}

// Reset 重置使能迁移, 复位
//
//	@Tags		模拟器
//	@Summary	重置使能迁移, 复位
//	@Param		id	body	string	true	"model id"
//	@Produce	json
//	@Success	200	{object}	response.Response
//	@Router		/simulator/reset [post]
func Reset(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": "reset",
	})
}

// GetTrace 获取模拟Trace
//
//	@Tags		模拟器
//	@Summary	获取模拟Trace
//	@Param		id	query	string	false	"model id"
//	@Produce	json
//	@Success	200	{object}	response.Response
//	@Router		/simulator/get-trace [get]
func GetTrace(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": "trace",
	})
}

// SaveTrace 保存模拟Trace
//
//	@Tags		模拟器
//	@Summary	保存模拟Trace
//	@Param		id	body	string	true	"model id"
//	@Produce	json
//	@Success	200	{object}	response.Response
//	@Router		/simulator/save-trace [post]
func SaveTrace(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": "save trace",
	})
}

// OpenTrace 打开模拟Trace
//
//	@Tags		模拟器
//	@Summary	打开模拟Trace
//	@Param		file	body	string	true	"trace path"
//	@Produce	json
//	@Success	200	{object}	response.Response
//	@Router		/simulator/open-trace [post]
func OpenTrace(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": "open trace",
	})
}

// RandomTrace 随机模拟Trace
//
//	@Tags		模拟器
//	@Summary	随机模拟Trace
//	@Param		id	body	string	true	"model id"
//	@Produce	json
//	@Success	200	{object}	response.Response
//	@Router		/simulator/random-trace [post]
func RandomTrace(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": "random trace",
	})
}

// GetGlobal 获取全局变量(对应模拟器中间的全局变量)
//
//	@Tags		模拟器
//	@Summary	获取全局变量(对应模拟器中间的全局变量)
//	@Param		id	query	string	false	"model id"
//	@Produce	json
//	@Success	200	{object}	response.Response
//	@Router		/simulator/get-global [get]
func GetGlobal(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": "globals",
	})
}

// GetLocal 获取各自动机的局部变量(对应模拟器中间的局部变量)
//
//	@Tags		模拟器
//	@Summary	获取各自动机的局部变量(对应模拟器中间的局部变量)
//	@Param		id	query	string	false	"model id"
//	@Produce	json
//	@Success	200	{object}	response.Response
//	@Router		/simulator/get-local [get]
func GetLocal(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": "locals",
	})
}

// GetCurrentStatus 获取当前状态(对应模拟器右上角的各自动机状态图)
//
//	@Tags		模拟器
//	@Summary	获取当前状态
//	@Param		id	query	string	false	"model id"
//	@Produce	json
//	@Success	200	{object}	response.Response
//	@Router		/simulator/get-current-status [get]
func GetCurrentStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": "current status",
	})
}

// GetSync 获取同步情况(对应模拟器右下角的同步图)
//
//	@Tags		模拟器
//	@Summary	获取同步情况(对应模拟器右下角的同步图)
//	@Param		id	query	string	false	"model id"
//	@Produce	json
//	@Success	200	{object}	response.Response
//	@Router		/simulator/get-sync [get]
func GetSync(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": "sync",
	})
}
