package cmd

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/boywei/go-zero-check/internal/util/global"
	log "github.com/sirupsen/logrus"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

var (
	// TODO: 这里应该根据不同的模型创建不同的解释器，并在模型删除时销毁
	i = interp.New(interp.Options{})
)

// RunCode 执行前端的语句
func RunCode(id, code string) (err error) {
	model, ok := global.ModelIdMap[id]
	if !ok {
		log.Warnf("Model path of %s not exists\n", id)
		return errors.New("model path of " + id + " not exists")
	}
	modelPath := model.Path
	// TODO: 这里应该只需要执行一次
	RunStaticCode(modelPath)
	// 先导入包，不然找不到
	i.Eval(`import "` + modelPath + `"`)
	result, err := i.Eval(code)
	if err != nil {
		log.Printf("执行代码%s出错: %v\n", code, err)
		return err
	}
	log.Println(result)
	return nil
}

func RunStaticCode(modelPath string) {

	// 创建解释器
	i.Use(stdlib.Symbols)

	// 遍历文件夹并执行所有Go文件
	err := filepath.Walk(modelPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			if _, err := i.Eval(string(content)); err != nil {
				log.Printf("在执行文件 %s 时发生错误: %v\n", info.Name(), err)
			}
		}
		return nil
	})

	if err != nil {
		log.Println("遍历文件夹时出错:", err)
		return
	}

}
