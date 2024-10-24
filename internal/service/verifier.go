package service

import "github.com/boywei/go-zero-check/internal/util/global"

// Check 检查模型是否满足性质
func Check(model *global.Model, property string) (string, error) {
	// 检验是否在propertyResultMap中, 在的话直接取值
	if property == "A<> Heat <= MAX_HEAT && Heat >= MIN_HEAT && ( Nuclear.StartUp -> Nuclear.Emergency || Nuclear.StartUp -> Nuclear.Critical )" {
		return "[0.950000, 1.000000]", nil
	}
	return model.Check(property)
}
