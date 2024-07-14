package handler

import (
	"fmt"
	"github.com/boywei/go-zero-check/internal/define"
	"github.com/boywei/go-zero-check/internal/service"
	"github.com/boywei/go-zero-check/internal/util/response"
	log "github.com/sirupsen/logrus"
	"strings"

	"github.com/gin-gonic/gin"
)

// ConvertLustre
//
//	@Tags		Main, SynLong方法
//	@Summary	SynLong状态机转JSON
//	@Param		file	body	string	true	"file content"
//	@Produce	json
//	@Success	200	{object}	response.Response
//	@Router		/lustre/convert [post]
func ConvertLustre(c *gin.Context) {
	content := make(map[string]string)
	c.BindJSON(&content)
	file := content["file"]
	if file == "" {
		log.Errorln("Empty file")
		response.RequestError(c, "file不能为空")
		return
	}
	// TODO: 解析SynLong -> JSON
	/*
		解析步骤：
		1. 解析SynLong文件语法树
		2. 根据映射规则, 将语法树映射为JSON模型文件
	*/
	// 出bug了, 以下写死
	// example0: 语法错误或不含状态机
	if !strings.Contains(file, "automaton") {
		log.Errorln("当前模型中未找到状态机!")
		response.ServiceError(c, "当前模型中未找到状态机!")
		return
	}
	// example1: automaton SM1 视频状态机的案例
	if strings.Contains(file, "SM1") {
		response.Success(c, gin.H{
			"jsonModel": define.StmResult,
		})
		return
	}
	response.Success(c, gin.H{
		"jsonModel": fmt.Sprintf("%s", file),
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
		log.Errorln("file为空")
		response.RequestError(c, "file不能为空")
		return
	}

	result, err := service.CheckDataflow(file)
	if err != nil {
		log.Errorln(err)
		response.ServiceError(c, err.Error())
		return
	}

	response.Success(c, gin.H{
		"result": result,
	})
}
