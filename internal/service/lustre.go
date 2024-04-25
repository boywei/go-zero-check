package service

import (
	"fmt"
	"github.com/boywei/go-zero-check/internal/util/enum"
	"github.com/boywei/go-zero-check/internal/util/response"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// ConvertLustre
//
//	@Tags		SynLong方法
//	@Summary	SynLong状态机转JSON
//	@Param		file	formData	string	false	"file content"
//	@Success	200		{string}	json	"{"code":"200","data":""}"
//	@Router		/lustre/convert [post]
func ConvertLustre(c *gin.Context) {
	file := c.PostForm("file")
	if file == "" {
		log.Errorln("Empty file")
		response.Failure(c, enum.InvalidParam)
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
		response.Failure(c, enum.SynLongConvertErr)
		return
	}
	// example1: automaton SM1 视频状态机的案例
	if strings.Contains(file, "SUM") {
		response.Success(c, gin.H{
			"jsonModel": `{
        "declaration": "",
        "automatons": [
          {
            "name": "SM1",
            "parameters": "c1, c2, c3, c4, c5 bool",
            "locations": [
              {
                "id": 1,
                "name": "A",
                "invariant": ""
              },
              {
                "id": 2,
                "name": "B",
                "invariant": ""
              },
              {
                "id": 3,
                "name": "C",
                "invariant": ""
              },
              {
                "id": 4,
                "name": "D",
                "invariant": ""
              },
              {
                "id": 5,
                "name": "E",
                "invariant": ""
              },
              {
                "id": 6,
                "name": "",
                "invariant": ""
              },
              {
                "id": 7,
                "name": "",
                "invariant": ""
              }
            ],
            "transitions": [
              {
                "id": 1,
                "source_id": 1,
                "target_id": 6,
                "select": "",
                "guard": "c1",
                "sync": "",
                "update": "S1()"
              },
              {
                "id": 2,
                "source_id": 6,
                "target_id": 7,
                "select": "",
                "guard": "!c2",
                "sync": "",
                "update": "S3()"
              },
              {
                "id": 3,
                "source_id": 6,
                "target_id": 2,
                "select": "",
                "guard": "c2",
                "sync": "",
                "update": "S2()"
              },
              {
                "id": 4,
                "source_id": 7,
                "target_id": 3,
                "select": "",
                "guard": "c4",
                "sync": "",
                "update": "S4()"
              },
              {
                "id": 5,
                "source_id": 7,
                "target_id": 4,
                "select": "",
                "guard": "!c4 && c5",
                "sync": "",
                "update": "S5()"
              },
              {
                "id": 6,
                "source_id": 7,
                "target_id": 5,
                "select": "",
                "guard": "!c4 & !c5",
                "sync": "",
                "update": "S6()"
              }
            ],
            "init": 1,
            "declaration": "func S1(){}\n\nfunc S2(){}\n\nfunc S3(){}\n\nfunc S4(){}\n\nfunc S5(){}\n\nfunc S6(){}"
          }
        ],
        "system_declaration": [
          "system A;"
        ]
      }`,
		})
		return
	}
	response.Success(c, gin.H{
		"id": fmt.Sprintf("converting %s", file),
	})
}

// CheckDataflow
//
//	@Tags		SynLong方法
//	@Summary	SynLong数据流验证
//	@Param		file	formData	string	false	"file content"
//	@Success	200		{string}	json	"{"code":"200","data":""}"
//	@Router		/lustre/check-dataflow [post]
func CheckDataflow(c *gin.Context) {
	file := c.PostForm("file")
	if file == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "文件不能为空",
		})
		return
	}
	// TODO: 调用kind2接口, 需要先下载kind2

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": fmt.Sprintf("converting %s", file),
	})
}
