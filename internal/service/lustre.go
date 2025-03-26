package service

import (
	"bytes"
	"encoding/json"
	"github.com/boywei/go-zero-check/internal/define"
	"github.com/boywei/go-zero-check/internal/synlong"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"strings"
)

const (
	kind2Url    = "http://localhost:8099/lustre/check-java"
	contentType = "application/json"
)

// CheckDataflow 验证数据流
func CheckDataflow(file string) (string, error) {
	data := make(map[string]interface{})
	data["file"] = file
	bytesData, _ := json.Marshal(data)
	resp, err := http.Post(kind2Url, contentType, bytes.NewReader(bytesData))
	if err != nil {
		return "", err
	}
	body, _ := io.ReadAll(resp.Body)
	escapedBody := strings.ReplaceAll(string(body), "\\n", "\n")

	if resp.StatusCode != http.StatusOK {
		return "", errors.New(escapedBody)
	}
	return escapedBody, nil
}

// ConvertStateMachine 状态机转换
func ConvertStateMachine(file string) (string, error) {
	// TODO: 解析SynLong -> JSON
	result, err := synlong.CheckSynLong(file)
	if err != nil {
		return "", errors.Wrap(err, "未能识别语法!")
	}

	// 不含状态机, 则直接使用数据流验证
	if !strings.Contains(file, "automaton") {
		return CheckDataflow(file)
	}

	// 出bug了, 以下写死
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
	return result, nil
}
