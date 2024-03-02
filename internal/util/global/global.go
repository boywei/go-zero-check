package global

import (
	"github.com/traefik/yaegi/interp"
	"os"
)

type ModelMap struct {
	Name   string
	Path   string              // 模型的生成目录
	Interp *interp.Interpreter // 模型的解释器地址
}

var (
	ModelIdMap = make(map[string]*ModelMap)
)

// Crush 删除当前model时释放资源
func (model *ModelMap) Crush() {
	os.RemoveAll(model.Path)
}
