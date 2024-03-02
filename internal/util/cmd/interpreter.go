package cmd

import (
	"errors"
	"github.com/boywei/go-zero-check/internal/util/global"
	log "github.com/sirupsen/logrus"
	"reflect"
)

// RunCode 执行前端的语句
func RunCode(id, code string) (*reflect.Value, error) {
	model, ok := global.ModelIdMap[id]
	if !ok {
		log.Warnf("Model path of %s not exists\n", id)
		return nil, errors.New("model path of " + id + " not exists")
	}
	i := model.Interp
	// 先导入包，不然找不到
	_, err := i.Eval(`import "./` + model.Path + `"`) // TODO: 这个路径令人疑惑
	if err != nil {
		log.Error(err)
		return nil, err
	}
	result, err := i.Eval(code)
	if err != nil {
		log.Printf("执行代码%s出错: %v\n", code, err)
		return nil, err
	}
	log.Println(result)
	return &result, nil
}

func RunStaticCode(id string) (*reflect.Value, error) {
	model, ok := global.ModelIdMap[id]
	if !ok {
		log.Warnf("Model path of %s not exists\n", id)
		return nil, errors.New("model path of " + id + " not exists")
	}
	result, err := model.Interp.EvalPath("./" + model.Path) // TODO: 这个路径真是令人疑惑
	return &result, err
}
