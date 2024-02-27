package service

import (
	"fmt"
	"github.com/boywei/go-zero-check/internal/util/json"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// Convert
// @Tags 转换方法
// @Summary 前端传过来的JSON转Go对象
// @Param file formData string false "json file path"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /convert [post]
func Convert(c *gin.Context) {
	file := c.PostForm("file")
	if file == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "文件不能为空",
		})
		return
	}
	object, err := json.ConvertJson2Uppaal(file)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  fmt.Sprintf("JSON decode error: %v", err),
		})
		return
	}
	// TODO: 处理转化后的object
	/*
		处理步骤包括:
		0. 在demo目录下新建一个目录: <模型名>_demo, 后续生成的模型文件都放在此目录下
		1. 处理Declaration: 将全局的Declaration生成新的go文件: declaration.go
		2. 处理Automatons: 将Automatons数组的每个元素生成一个新的go文件: <自动机名>.go
		3. 处理SystemDeclaration: 执行该语句(执行方式待考量), 保存声明后的Uppaal对象(可使用缓存?)
		4. 做好模型图中的语法检查
		5. 在新的端口启动项目, 给出错误信息
	*/
	// 0. 新建模型目录
	modelName := "test"
	modelPath := "internal/demo/" + modelName + "/"
	err = os.Mkdir(modelPath, os.ModePerm)
	if err != nil {
		log.Fatal("Mkdir err: ", err)
	}
	// 1. 处理Declaration
	df, err := os.Create(modelPath + "declaration.go")
	defer df.Close()
	if err != nil {
		log.Fatal("Create declaration file err: ", err)
	}
	_, err = df.WriteString(string(object.Declaration))
	if err != nil {
		log.Fatal("Write global declaration err: ", err)
	}
	// 2. 处理Automatons
	for _, automaton := range object.Automatons {
		// 用一个builder来拼接所有的内容
		var builder strings.Builder

		// 1) name string
		f, err := os.Create(modelPath + automaton.Name + ".go")
		defer f.Close()
		if err != nil {
			log.Fatal("Handle automaton err: ", err)
		}
		// TODO
		// 2) Parameters  []Parameter
		// 3) Locations   []Location
		// 4) Transitions []Transition
		// 5) Init        Location
		// 6) Declarations Declaration
		builder.WriteString(string(automaton.Declaration))

		// 将builder的内容写入文件
		_, err = f.WriteString(builder.String())
		if err != nil {
			log.Fatal("Write automaton declaration err: ", err)
		}
	}
	// 3. 处理SystemDeclaration: 执行该语句(执行方式待考量), 保存声明后的Uppaal对象(可使用缓存?)

	// 4. 做好模型图中的语法检查

	// 5. 在新的端口启动项目, 给出错误信息

	data := object
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": data,
	})
}
