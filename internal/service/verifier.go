package service

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"net/http"
)

const (
	kind2Url = "http://localhost:8081/"
	contentType = "application/json"
)

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