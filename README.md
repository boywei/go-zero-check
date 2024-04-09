# go-zero-check
一个模型正确性检查系统的后端

## 正常模型的项目结构
- Declaration
  - Variable
  - Constant
  - Clock
  - Chan
  - Function
- Automaton-1
  - Name
  - Parameter
  - Location
  - init Location
  - Transition
  - Nail
  - Declaration
- Automaton-2
- ...
- Automaton-N
- System Declaration

## Lustre和Uppaal对应的关系

## 模型检查算法步骤
```
while not deadlock:
	current_locations = map[automaton -> location] #获取所有自动机的当前状态
	next_transitions = map[automaton -> transitions[]] #获取所有自动机当前能迁移的边
	select_next(automaton, transition) #选择某条迁移前进
	step() # 前进并更新其他自动机
	forall property in properties: #检测所要验证的每条性质是否满足
	    check(property)
```
