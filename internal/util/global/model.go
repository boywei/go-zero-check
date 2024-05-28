package global

import (
	log "github.com/sirupsen/logrus"
	"github.com/traefik/yaegi/interp"
	"os"
)

type Model struct {
	Name        string
	Path        string              // 模型的生成目录
	Interpreter *interp.Interpreter // 模型的解释器地址
}

var (
	ModelIdMap = make(map[string]*Model)
)

// RemoveDir 删除当前model时释放资源
func (model *Model) RemoveDir() {
	err := os.RemoveAll(model.Path)
	if err != nil {
		log.Error("No such model: ", model.Name)
		return
	}
}

/*
模型检查算法步骤
```
while not deadlock:
	current_locations = map[automaton -> location] #获取所有自动机的当前状态
	next_transitions = map[automaton -> transitions[]] #获取所有自动机当前能迁移的边
	select_next(automaton, transition) #选择某条迁移前进
	step() # 前进并更新其他自动机
	forall property in properties: #检测所要验证的每条性质是否满足
	    check(property)
```
*/

// Step 使模型前进一步
func (model *Model) Step() {
	// TODO: step
}

// Run 运行模型
func (model *Model) Run() {
	// TODO: run
}

// Check 检查模型是否满足性质
func (model *Model) Check(property string) (bool, error) {
	// TODO: 检查性质是否成立
	return true, nil
}