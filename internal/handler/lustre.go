package handler

import (
	"github.com/boywei/go-zero-check/internal/service"
	"github.com/boywei/go-zero-check/internal/util/response"
	"github.com/gin-gonic/gin"
)

// ConvertLustre
//
//	@Tags		Main, SynLong方法
//	@Summary	SynLong状态机转JSON
//	@Param		file	body	string	true	"file content"
//	@Produce	json
//	@Success	200	{object}	response.Response
//	@Router		/lustre/check [post]
func ConvertLustre(c *gin.Context) {
	content := make(map[string]string)
	c.BindJSON(&content)
	file := content["file"]
	if file == "" {
		response.ServiceError(c, "file不能为空")
		return
	}
	result, err := service.ConvertStateMachine(file)
	if err != nil {
		response.ServiceError(c, "解析错误: "+err.Error())
		return
	}
	response.Success(c, gin.H{
		"result": result,
	})
}

// CheckDataflow
//
//	@Tags		Main, SynLong方法
//	@Summary	SynLong数据流验证
//	@Param		file	body	string	true	"file content"
//	@Produce	json
//	@Success	200	{object}	response.Response
//	@Router		/lustre/check-dataflow [post]
func CheckDataflow(c *gin.Context) {
	content := make(map[string]string)
	c.BindJSON(&content)
	file := content["file"]
	if file == "" {
		response.ServiceError(c, "file不能为空")
		return
	}

	result, err := service.CheckDataflow(file)
	if err != nil {
		response.ServiceError(c, err.Error())
		return
	}

	response.Success(c, gin.H{
		"result": result,
	})
}
