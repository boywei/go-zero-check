package service

import (
	"github.com/boywei/go-zero-check/internal/model"
	"github.com/boywei/go-zero-check/internal/util/cmd"
	"github.com/boywei/go-zero-check/internal/util/global"
	"github.com/pkg/errors"
	"github.com/traefik/yaegi/interp"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func Convert(modelName string, object *model.Uppaal) error {
	// 生成模型目录
	modelPath := filepath.Join("demo", modelName)
	if err := os.MkdirAll(modelPath, os.ModePerm); err != nil {
		return errors.Wrap(err, "make directory failed")
	}

	// 生成define.go文件
	src, dest := filepath.Join("demo", "define.txt"), filepath.Join(modelPath, "define.go")
	if err := copyFile(src, dest, modelName); err != nil {
		return errors.Wrap(err, "copy file failed")
	}

	// 生成declaration.go文件
	if err := generateDeclarationFile(modelPath, modelName, object.Declaration); err != nil {
		return errors.Wrap(err, "generate declaration file failed")
	}

	// 生成<automaton>.go文件
	for _, automaton := range object.Automatons {
		if err := generateAutomatonFile(modelPath, modelName, automaton); err != nil {
			return errors.Wrap(err, "generate automaton file failed")
		}
	}

	// 保存模型, 添加到全局变量中, 生成模型的解释器
	global.ModelIdMap[modelName] = &global.Model{
		Name:        modelName,
		Path:        modelPath,
		Interpreter: interp.New(interp.Options{}),
	}

	// 运行模型
	if _, err := cmd.RunStaticCode(modelName); err != nil {
		return errors.Wrap(err, "run model failed")
	}

	return nil
}

// copyFile 复制文件, 并添加package name到文件头部.
func copyFile(src, dst, modelName string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return errors.Wrap(err, "open file failed")
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(dst)
	if err != nil {
		return errors.Wrap(err, "create define file failed")
	}
	defer destinationFile.Close()
	// 添加包名
	_, err = destinationFile.WriteString("package " + modelName + "\n\n")
	if err != nil {
		return errors.Wrap(err, "write to define file failed")
	}
	// 使用io.Copy复制文件
	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return errors.Wrap(err, "copy to destination file failed")
	}
	return nil
}

// generateDeclarationFile 生成declaration文件
func generateDeclarationFile(modelPath, modelName string, declaration model.Declaration) error {
	file, err := os.Create(filepath.Join(modelPath, "declaration.go"))
	defer file.Close()
	if err != nil {
		return errors.Wrap(err, "创建文件失败")
	}
	// 添加包名
	_, err = file.WriteString("package " + modelName + "\n\n" + string(declaration))
	if err != nil {
		return errors.Wrap(err, "写入文件失败")
	}
	return nil
}

// generateAutomatonFile 处理Automatons -> 针对每个自动机，生成<automaton_name>.go文件
func generateAutomatonFile(modelPath, modelName string, automaton model.Automaton) error {
	// 用一个builder来拼接所有的内容
	var builder strings.Builder

	// 1) name string
	f, err := os.Create(filepath.Join(modelPath, automaton.Name+".go"))
	defer f.Close()
	if err != nil {
		return errors.Wrap(err, "创建文件失败")
	}
	builder.WriteString("package " + modelName + "\n\n")
	// TODO: 修改声明的方式，根据parameter来声明
	builder.WriteString("var " + string(automaton.Parameters) + "\n\n")
	builder.WriteString("var (\n\t" + automaton.Name + " = &Automaton{\n")
	builder.WriteString("\t\tName: \"" + automaton.Name + "\",\n")

	// 2) Parameters  Parameter
	builder.WriteString("\t\tParameters: \"" + string(automaton.Parameters) + "\",\n")

	// 3) Locations   []Location
	builder.WriteString("\t\tLocations: []Location{\n")
	for _, location := range automaton.Locations {
		builder.WriteString("\t\t\t{\n")
		builder.WriteString("\t\t\t\tId: " + strconv.Itoa(location.Id) + ",\n")
		builder.WriteString("\t\t\t\tName: \"" + location.Name + "\",\n")
		builder.WriteString("\t\t\t\tInvariant: func() bool {\n\t\t\t\t\treturn " + getValue(location.Invariant) + "\n\t\t\t\t},\n")
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
		builder.WriteString("\t\t\t\tSelect: \"" + transition.Select + "\",\n")
		builder.WriteString("\t\t\t\tGuard: func() bool {\n\t\t\t\t\treturn " + getValue(transition.Guard) + "\n\t\t\t\t},\n")
		builder.WriteString("\t\t\t\tSync: func() bool {\n\t\t\t\t\treturn " + getValue(transition.Sync) + "\n\t\t\t\t},\n")
		builder.WriteString("\t\t\t\tUpdate: func() {\n\t\t\t\t\t" + transition.Update + "\n\t\t\t\t},\n")
		builder.WriteString("\t\t\t},\n")
	}
	builder.WriteString("\t\t},\n")

	// 5) Init        Location
	builder.WriteString("\t\tInit: " + strconv.Itoa(automaton.Init) + ",\n")
	builder.WriteString("\t}\n\n")
	builder.WriteString(")\n\n")
	// 6) Declarations Declaration
	builder.WriteString(string(automaton.Declaration))

	// 将builder的内容写入文件
	_, err = f.WriteString(builder.String())
	if err != nil {
		return errors.Wrap(err, "写入文件失败")
	}
	return nil
}

// getValue 用于获取字符串（uppaal中的条件判断语句）的默认值true
func getValue(data string) string {
	if data == "" {
		return "true"
	}
	return data
}

// DeleteModelById 根据id删除对应的模型
func DeleteModelById(id string) error {
	// 删除modelMap中的内容
	m, ok := global.ModelIdMap[id]
	if !ok {
		return errors.New("该id对应的模型不存在: " + id)
	}
	err := m.RemoveDir()
	if err != nil {
		return errors.Wrap(err, "删除模型对应的目录失败")
	}
	delete(global.ModelIdMap, id)
	return nil
}
