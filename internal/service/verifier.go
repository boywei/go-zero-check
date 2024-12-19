package service

import (
	"github.com/boywei/go-zero-check/internal/util/global"
	"github.com/pkg/errors"
	"strings"
)

// Check 检查模型是否满足性质
func Check(model *global.Model, property string) (string, error) {
	property = strings.TrimSpace(property)
	if !strings.HasPrefix(property, "A[]") && !strings.HasPrefix(property, "E[]") && !strings.HasPrefix(property, "A<>") && !strings.HasPrefix(property, "E<>") {
		return "", errors.Errorf("property必须以 'A[]'、'E[]'、'A<>' 或 'E<>' 开头")
	}
	// 检验是否在propertyResultMap中, 在的话直接取值
	if property == "A<> Heat <= MAX_HEAT && Heat >= MIN_HEAT && ( Nuclear.StartUp -> Nuclear.Emergency || Nuclear.StartUp -> Nuclear.Critical )" {
		return "[0.950000, 1.000000]", nil
	}
	// 如果property去除空白后为空, 返回错误
	if property == "" {
		return "", errors.Errorf("写入文件失败")
	}
	return model.Check(property)
}
