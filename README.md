# go-zero-check
一个模型正确性检查系统的后端

## 开始
### 0. 预安装
需要先安装go1.20或以上版本
> 安装包下载地址为：https://go.dev/dl/  
> 如果打不开可以使用这个地址：https://golang.google.cn/dl/  

安装后显示输入`go version`检查是否输出版本号, 若正常输出版本号, 则继续设置:
```shell
go env -w GOPROXY=https://goproxy.io,direct
go env -w GO111MODULE=on
```

### 1. 启动
首先在终端**进入到项目根目录**,
#### 启动方式1: 直接启动
```shell
go run main.go
```

#### 启动方式2: 打包成可执行文件
| 系统        | 执行命令                           | 访问方式              |
|:----------|:-------------------------------|:------------------|
| Mac Intel | `./deploy/exec/darwin_amd.sh`  | `./check-backend` |
| Mac M1    | `./deploy/exec/darwin_arm.sh`  | `./check-backend` |
| Linux Amd | `./deploy/exec/linux_amd.sh`   | `./check-backend` |
| Linux Arm | `./deploy/exec/linux_arm.sh`   | `./check-backend` |
| Windows   | `./deploy/exec/windows_amd.sh` | `./check-backend` |

#### 启动方式3: 使用docker
需要先安装docker
```shell
docker build -f ./deploy/docker/Dockerfile -t check-backend .
docker run -d -p 8080:8080 check-backend
```

### 2. 访问
在浏览器输入: `http://localhost:8080/`, 若显示`not found`则正在运行

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
