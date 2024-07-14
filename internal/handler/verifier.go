package handler

import (
	"github.com/boywei/go-zero-check/internal/util/global"
	"github.com/boywei/go-zero-check/internal/util/response"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Verify
//
//	@Tags		验证器, Main
//	@Summary	验证某条性质是否成立
//	@Param		id			body	string	true	"model id"
//	@Param		property	body	string	true	"property""
//	@Produce	json
//	@Success	200	{object}	response.Response
//	@Router		/verify [post]
func Verify(c *gin.Context) {
	content := make(map[string]string)
	c.BindJSON(&content)
	id := content["id"]
	property := content["property"]
	if property == "" || id == "" {
		log.Errorf("id[%s]或property[%s]为空\n", id, property)
		response.RequestError(c, "id或property不能为空")
		return
	}

	// 检查模型是否存在
	model, ok := global.ModelIdMap[id]
	if !ok {
		log.Errorln("模型不存在: ", id)
		response.RequestError(c, "当前id对应的模型不存在")
		return
	}

	// 获取验证结果
	result, err := model.Check(property)
	if err != nil {
		log.Errorln("Model verify error: ", err)
		response.RequestError(c, "模型验证失败: "+err.Error())
		return
	}
	response.Success(c, gin.H{
		"result": result,
	})
}
