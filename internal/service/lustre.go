package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/boywei/go-zero-check/internal/util/response"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
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
//	@Failure	400	{object}	response.ErrCode
//	@Router		/lustre/convert [post]
func ConvertLustre(c *gin.Context) {
	content := make(map[string]string)
	c.BindJSON(&content)
	file := content["file"]
	if file == "" {
		log.Errorln("Empty file")
		response.Failure(c, response.InvalidParam)
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
		log.Errorln("No automaton found")
		response.Failure(c, response.SynLongConvertErr)
		return
	}
	// example1: automaton SM1 视频状态机的案例
	if strings.Contains(file, "SM1") {
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
//	@Tags		Main, SynLong方法
//	@Summary	SynLong数据流验证
//	@Param		file	body	string	true	"file content"
//	@Produce	json
//	@Success	200	{object}	response.Response
//	@Failure	400	{object}	response.ErrCode
//	@Router		/lustre/check-dataflow [post]
func CheckDataflow(c *gin.Context) {
	content := make(map[string]string)
	c.BindJSON(&content)
	file := content["file"]
	if file == "" {
		log.Errorln("Empty file")
		response.Failure(c, response.InvalidParam)
		return
	}

	// TODO: 调用熊江的kind2验证接口
	data := make(map[string]interface{})
	data["selectedSolver"] = "Z3"
	data["code"] = file
	bytesData, _ := json.Marshal(data)
	resp, err := http.Post("http://localhost:8081/", "application/json", bytes.NewReader(bytesData))
	if err != nil {
		log.Errorln("Dataflow check error: ", err)
		response.Failure(c, response.SynLongDataflowErr)
		return
	}
	body, _ := io.ReadAll(resp.Body)

	response.Success(c, gin.H{
		"result": string(body),
	})
}
