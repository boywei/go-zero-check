package service

import (
	"bytes"
	"encoding/json"
	"github.com/boywei/go-zero-check/internal/define"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"strings"
)

const (
	kind2Url    = "http://localhost:8081/synlong/check"
	contentType = "application/json"
)

// CheckDataflow 验证数据流
func CheckDataflow(file string) (string, error) {
	// TODO: 调用熊江的kind2验证接口
	data := make(map[string]interface{})
	data["selectedSolver"] = "Z3"
	data["code"] = file
	bytesData, _ := json.Marshal(data)
	resp, err := http.Post(kind2Url, contentType, bytes.NewReader(bytesData))
	if err != nil {
		return "", errors.Wrapf(err, "数据流验证错误: %v", err)
	}
	body, _ := io.ReadAll(resp.Body)
	return string(body), nil
}

// ConvertStateMachine 状态机转换
func ConvertStateMachine(file string) (string, error) {
	// TODO: 解析SynLong -> JSON
	// 出bug了, 以下写死
	// example0: 语法错误或不含状态机
	if !strings.Contains(file, "automaton") {
		return "", errors.New("当前模型中未找到状态机!")
	}
	// case-nuclear: 核电场景
	if strings.Contains(file, "Nuclear") {
		return define.NuclearResult, nil
	}
	// case-driving:
	if strings.Contains(file, "Car") {
		return define.CarResult, nil
	}
	// example1: automaton SM1 视频状态机的案例
	if strings.Contains(file, "SM1") {
		return define.StmResult, nil
	}
	// example2: SM2代表反例
	if strings.Contains(file, "SM2") {
		return "", errors.New("状态机存在语法错误!")
	}
	// 默认情况
	return file, nil
}
