package global

import (
	"os/exec"
	"sync"
)

var (
	// 全局的host和port，用于启动新的模型
	Host      string = "127.0.0.1"
	Port      int    = 8900
	PortMutex sync.Mutex

	// 全局模型执行的命令，用于模型终止
	ModelCmdMap map[string]*exec.Cmd
)

func IncreasePort() {
	PortMutex.Lock()
	Port++
	PortMutex.Unlock()
}
