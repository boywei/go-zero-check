package service

import (
	"fmt"
	"github.com/boywei/go-zero-check/internal/util/cmd"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/boywei/go-zero-check/internal/middleware"
	"github.com/boywei/go-zero-check/internal/util/enum"
	"github.com/boywei/go-zero-check/internal/util/global"
	"github.com/boywei/go-zero-check/internal/util/json"
	"github.com/boywei/go-zero-check/internal/util/response"
	log "github.com/sirupsen/logrus"
	"github.com/traefik/yaegi/interp"

	"github.com/gin-gonic/gin"
)

// Convert
//
//	@Tags		编辑器
//	@Summary	前端传过来的JSON转Go对象
//	@Param		file	formData	string	false	"json file path"
//	@Success	200		{string}	json	"{"code":"200","data":""}"
//	@Router		/model/convert [post]
func Convert(c *gin.Context) {
	file := c.PostForm("file")
	if file == "" {
		log.Errorln("Empty file")
		response.Failure(c, enum.InvalidParam)
		return
	}
	object, err := json.ConvertJson2Uppaal(file)
	if err != nil {
		log.Errorln("JSON decode error: ", err)
		response.Failure(c, enum.InvalidParam)
		return
	}

	// 0. 新建模型目录
	modelName := "test"
	modelPath := "demo/" + modelName + "/"
	err = os.Mkdir(modelPath, os.ModePerm)
	if err != nil {
		if os.IsExist(err) { // 当前目录已经存在, 那么清空
			_ = os.RemoveAll(modelPath)
			_ = os.Mkdir(modelPath, os.ModePerm)
		} else { // 其他错误
			log.Errorln("Mkdir err: ", err)
			response.Failure(c, enum.ModelConvertErr)
			return
		}
	}

	// 复制define.go文件到模型目录下
	sourceFile, err := os.Open("demo/define.txt")
	if err != nil {
		panic(err)
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create("demo/" + modelName + "/define.go")
	if err != nil {
		panic(err)
	}
	defer destinationFile.Close()
	// 使用io.Copy复制文件
	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		panic(err)
	}

	// 1. 处理Declaration
	df, err := os.Create(modelPath + "declaration.go")
	defer df.Close()
	if err != nil {
		log.Errorln("Create declaration file err: ", err)
		response.Failure(c, enum.ModelConvertErr)
		return
	}
	_, err = df.WriteString("package " + modelName + "\n\n" + string(object.Declaration))
	if err != nil {
		log.Errorln("Write global declaration err: ", err)
		response.Failure(c, enum.ModelConvertErr)
		return
	}
	// 2. 处理Automatons
	for _, automaton := range object.Automatons {
		// 用一个builder来拼接所有的内容
		var builder strings.Builder

		// 1) name string
		f, err := os.Create(modelPath + automaton.Name + ".go")
		defer f.Close()
		if err != nil {
			log.Errorln("Handle automaton err: ", err)
			response.Failure(c, enum.ModelConvertErr)
			return
		}
		builder.WriteString("package " + modelName + "\n\n")
		// TODO
		builder.WriteString("var (\n\t" + automaton.Name + " = &Automaton{\n")
		builder.WriteString("\t\tName: \"" + automaton.Name + "\",")

		// 2) Parameters  []Parameter
		builder.WriteString("\t\tParameters: []string{\n")
		for _, parameter := range automaton.Parameters {
			builder.WriteString("\t\t\t\"" + string(parameter) + "\",\n")
		}
		builder.WriteString("\t\t},\n")

		// 3) Locations   []Location
		builder.WriteString("\t\tLocations: []Location{\n")
		for _, location := range automaton.Locations {
			builder.WriteString("\t\t\t{\n")
			builder.WriteString("\t\t\t\tId: " + strconv.Itoa(location.Id) + ",\n")
			builder.WriteString("\t\t\t\tName: \"" + location.Name + "\",\n")
			builder.WriteString("\t\t\t\tInvariant: \"" + string(location.Invariant) + "\",\n")
			builder.WriteString("\t\t\t},\n")
		}
		builder.WriteString("\t\t},\n")

		// 4) Transitions []Transition
		builder.WriteString("\t\tTransitions: []Transition{\n")
		for _, transition := range automaton.Transitions {
			builder.WriteString("\t\t\t{\n")
			builder.WriteString("\t\t\t\tId: " + strconv.Itoa(transition.Id) + ",\n")
			builder.WriteString("\t\t\t\tSourceId: " + strconv.Itoa(transition.SourceId) + ",\n")
			builder.WriteString("\t\t\t\tDestinationId: " + strconv.Itoa(transition.DestinationId) + ",\n")
			builder.WriteString("\t\t\t\tSelect: \"" + string(transition.Select) + "\",\n")
			builder.WriteString("\t\t\t\tGuard: \"" + string(transition.Guard) + "\",\n")
			builder.WriteString("\t\t\t\tSync: \"" + string(transition.Sync) + "\",\n")
			builder.WriteString("\t\t\t\tUpdate: \"" + string(transition.Update) + "\",\n")
			builder.WriteString("\t\t\t},\n")
		}
		builder.WriteString("\t\t},\n")

		// TODO: 还得解析出source和target
		// 5) Init        Location
		builder.WriteString("\t\tInit: " + strconv.Itoa(automaton.Init) + ",\n")
		builder.WriteString("\t}\n\n")
		builder.WriteString(")\n\n")
		// 6) Declarations Declaration
		builder.WriteString(string(automaton.Declaration))

		// 将builder的内容写入文件
		_, err = f.WriteString(builder.String())
		if err != nil {
			log.Errorln("Write automaton err: ", err)
			response.Failure(c, enum.ModelConvertErr)
			return
		}
	}
	// 3. 处理SystemDeclaration: 执行该语句(执行方式待考量)

	// 4. 做好模型图中的语法检查

	// 5. 使用解释器运行代码, 若无报错则保存声明后的Uppaal对象(使用同一redis缓存！), 并返回生成的uuid字符串
	// 先保存id并返回新程序信息，再启动
	uuid, err := middleware.SetModel(object)
	if err != nil {
		log.Errorln("Cache set model error: ", err)
		response.Failure(c, enum.ModelRunErr)
		return
	}
	// 保存model的一些全局信息
	global.ModelIdMap[uuid] = &global.ModelMap{
		Name:   modelName,
		Path:   modelPath,
		Interp: interp.New(interp.Options{}),
	}
	if _, err = cmd.RunStaticCode(uuid); err != nil {
		log.Errorln("Model run error: ", err)
		response.Failure(c, enum.ModelRunErr)
		return
	}
	response.Success(c, gin.H{
		"id": uuid,
	})

}

// DeleteModel
//
//	@Tags		编辑器
//	@Summary	删除一个模型
//	@Param		id	formData	string	true	"model's id"
//	@Success	200	{string}	json	"{"code":"200","data":""}"
//	@Router		/model/delete [post]
func DeleteModel(c *gin.Context) {
	id := c.PostForm("id")
	if id == "" {
		log.Errorln("Empty id")
		response.Failure(c, enum.InvalidParam)
		return
	}
	err := middleware.DeleteModelById(id)
	if err != nil {
		response.Failure(c, enum.ModelNotExist)
		return
	}
	response.Success(c, nil)
}

// TestModel
//
//	@Tags		编辑器
//	@Summary	测试代码是否生成成功, 返回输入整数的下一个整数
//	@Param		id	query		string	true	"model's id"
//	@Param		num	query		int		true	"number"
//	@Success	200	{string}	json	"{"code":"200","data":"3"}"
//	@Router		/model/test [get]
func TestModel(c *gin.Context) {
	id := c.Query("id")
	num := c.Query("num")
	if id == "" || num == "" {
		log.Errorln("Empty param")
		response.Failure(c, enum.InvalidParam)
		return
	}
	value, err := strconv.Atoi(num)
	if err != nil {
		log.Errorln("Please input number")
		response.Failure(c, enum.InvalidParam)
		return
	}
	model, ok := global.ModelIdMap[id]
	if !ok {
		log.Errorf("%s not exists\n", id)
		response.Failure(c, enum.ModelNotExist)
		return
	}
	code := model.Name + "." + fmt.Sprintf("Add(%d, 1)", value)
	result, err := cmd.RunCode(id, code)
	if err != nil {
		log.Errorln("Model run error: ", err)
		response.Failure(c, enum.ModelRunErr)
		return
	}
	response.Success(c, gin.H{
		"result": fmt.Sprintf("%v", *result), // TODO: 返回结果的形式有待考量
	})
}

// RunModel
//
//	@Tags		编辑器
//	@Summary	运行一条语句(用于测试)
//	@Param		id	formData		string	true	"model's id"
//	@Param		code	formData		string		true	"code"
//	@Success	200	{string}	json	"{"code":"200","data":"3"}"
//	@Router		/model/run [post]
func RunModel(c *gin.Context) {
	id := c.PostForm("id")
	code := c.PostForm("code")
	if id == "" || code == "" {
		log.Errorln("Empty param")
		response.Failure(c, enum.InvalidParam)
		return
	}
	result, err := cmd.RunCode(id, code)
	if err != nil {
		log.Errorln("Model run error: ", err)
		response.Failure(c, enum.ModelRunErr)
		return
	}
	response.Success(c, gin.H{
		"result": fmt.Sprintf("%v", *result), // TODO: 返回结果的形式有待考量
	})
}
