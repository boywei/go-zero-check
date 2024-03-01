package global

import (
	"os"
	"os/exec"
	"sync"

	"github.com/traefik/yaegi/interp"
)

type ModelMap struct {
	Cmd    *exec.Cmd           // TODO: 弃用。模型的启动命令
	Path   string              // 模型的生成目录
	Interp *interp.Interpreter // 模型的解释器地址
}

var (
	// 全局的host和port，用于启动新的模型
	Host      string = "127.0.0.1"
	Port      int    = 8900
	PortMutex sync.Mutex

	ModelIdMap map[string]*ModelMap = make(map[string]*ModelMap)

	// // 全局模型执行的命令，用于模型终止; id -> cmd
	// ModelCmdMap map[string]*exec.Cmd = make(map[string]*exec.Cmd)

	// // 全局模型存放目录的地址; id -> modelPath
	// ModelPathMap map[string]string = make(map[string]string)

	// // 全局解释器; id -> interp
	// ModelInterpMap map[string]*interp.Interpreter = make(map[string]*interp.Interpreter)
)

func IncreasePort() {
	PortMutex.Lock()
	Port++
	PortMutex.Unlock()
}

// Crush 删除当前model时释放资源
func (model *ModelMap) Crush() {
	model.Cmd.Process.Kill()
	os.RemoveAll(model.Path)
}
