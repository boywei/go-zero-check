package json

import (
	"encoding/json"
	"github.com/pkg/errors"

	"github.com/boywei/go-zero-check/internal/model"
)

// ConvertJson2Uppaal JSON转Uppaal对象
func ConvertJson2Uppaal(jsonInput string) (*model.Uppaal, error) {
	var uppaal model.Uppaal
	err := json.Unmarshal([]byte(jsonInput), &uppaal)
	if err != nil {
		return nil, errors.Wrapf(err, "json解析错误: %v", err)
	}
	return &uppaal, nil
}

// ConvertUppaal2Json Uppaal对象转JSON
func ConvertUppaal2Json(uppaalModel model.Uppaal) (string, error) {
	bytes, err := json.MarshalIndent(uppaalModel, "", "  ")
	if err != nil {
		return "", errors.Wrapf(err, "json序列化错误: %v", err)
	}
	return string(bytes), nil
}
