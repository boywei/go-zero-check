package handler

import (
	"github.com/boywei/go-zero-check/internal/service"
	"strings"

	"github.com/boywei/go-zero-check/internal/util/json"
	"github.com/boywei/go-zero-check/internal/util/response"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Convert
//
//	@Tags		编辑器, Main
//	@Summary	前端传过来的JSON转Go对象
//	@Param		file	body	string	true	"JSON file"
//	@Produce	json
//	@Success	200	{object}	response.Response
//	@Router		/model/convert [post]
func Convert(c *gin.Context) {
	var content map[string]string
	if err := c.BindJSON(&content); err != nil {
		log.Errorln("JSON绑定错误: ", err)
		response.RequestError(c, "JSON格式错误: "+err.Error())
		return
	}

	file := content["file"]
	if file == "" {
		log.Errorln("file为空")
		response.RequestError(c, "file不能为空")
		return
	}

	object, err := json.ConvertJson2Uppaal(file)
	if err != nil {
		log.Errorln("JSON解析错误: ", err)
		response.ServiceError(c, "JSON解析错误: "+err.Error())
		return
	}

	// modelName直接设置为"model" + ip地址, 需要去除ip地址重的"."
	modelName := "model" + strings.ReplaceAll(c.ClientIP(), ".", "_")
	err = service.Convert(modelName, object)
	if err != nil {
		log.Errorln("模型转化错误: ", err)
		response.ServiceError(c, "模型转化错误"+err.Error())
		return
	}

	response.Success(c, gin.H{"id": modelName})
}

// DeleteModel
//
//	@Tags		编辑器
//	@Summary	删除一个模型
//	@Param		id	body	string	true	"model's id"
//	@Produce	json
//	@Success	200	{object}	response.Response
//	@Router		/model/delete [post]
func DeleteModel(c *gin.Context) {
	content := make(map[string]string)
	c.BindJSON(&content)
	id := content["id"]
	if id == "" {
		log.Errorln("id为空")
		response.RequestError(c, "id不能为空")
		return
	}
	err := service.DeleteModelById(id)
	if err != nil {
		log.Errorln("删除模型失败: ", err)
		response.ServiceError(c, "删除模型失败: "+err.Error())
		return
	}
	response.Success(c, nil)
}
