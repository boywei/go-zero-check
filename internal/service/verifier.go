package service

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
//	@Failure	400	{object}	response.ErrCode
//	@Router		/verify [post]
func Verify(c *gin.Context) {
	content := make(map[string]string)
	c.BindJSON(&content)
	id := content["id"]
	property := content["property"]
	if property == "" || id == "" {
		log.Errorln("Empty id or property")
		response.Failure(c, response.InvalidParam)
		return
	}

	// 检查模型是否存在
	model, ok := global.ModelIdMap[id]
	if !ok {
		log.Errorln("Model not exist")
		response.Failure(c, response.ModelNotExist)
		return
	}

	// 获取验证结果
	result, err := model.Check(property)
	if err != nil {
		log.Errorln("Model verify error: ", err)
		response.Failure(c, response.ModelVerifyErr)
		return
	}
	response.Success(c, gin.H{
		"result": result,
	})
}
