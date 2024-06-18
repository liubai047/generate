package test

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"log"
	"os"
	"strings"
	"testing"
	"text/template"

	"generate/enum"
)

var testMyEnum = enum.GenEnum{
	PkgName:  "RpStatusEnum",
	EnumPath: "generate/enum",
	Data: []enum.Item{
		{Name: "Success4", Val: "assc", Annotation: "成功咯"},
		{Name: "Fail3", Val: "fffc", Annotation: "失败咯"},
		{Name: "Wait2", Val: "wwwc", Annotation: "待拉咯"},
		{Name: "Deal1", Val: "dddc", Annotation: "待处理咯"},
	},
}

func TestDm2(t *testing.T) {
	generate("./gen.go.tmpl", testMyEnum)
}

func generate(tmplFile string, genData enum.GenEnum) {
	// 判断数据是否合法
	if genData.PkgName == "" || genData.EnumPath == "" || len(genData.Data) < 1 {
		log.Fatalf("GenEnum结构体数据不合规,请检查\n")
		return
	}
	// 打开并读取模板文件
	tfRes, err := os.ReadFile(tmplFile)
	if err != nil {
		log.Fatalf("模板文件读取失败: %s\n", err.Error())
		return
	}
	// 处理待写入路径
	dstFile := "./" + genData.PkgName + "/enum.go"
	// 处理待写入的文件路径
	var pathDir = ""
	if slashIndex := strings.LastIndex(dstFile, "/"); slashIndex != -1 {
		pathDir = dstFile[:slashIndex]
	}
	// 检查路径是否存在，如果不存在则创建
	if _, err := os.Stat(pathDir); os.IsNotExist(err) {
		err = os.MkdirAll(pathDir, 0755)
		if err != nil {
			log.Fatalf("创建路径失败: %s\n", err.Error())
			return
		}
		fmt.Printf("路径 %s 已创建\n", pathDir)
	}
	// 打开待写入文件
	f, err := os.OpenFile(dstFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("打开待写入文件失败: %s\n", err.Error())
		return
	}
	defer f.Close()
	// 生成模板函数
	tmpl := template.Must(template.New("generateEnum").Funcs(template.FuncMap{
		"lowerFirst": lowerFirst,
		"upFirst":    upFirst,
		"camelCase":  camelCase,
		"quoteIfStr": quoteIfString,
	}).Parse(string(tfRes)))
	// 获取模板最终生成结果
	var codeBuf bytes.Buffer
	cGenData := createGenData(genData)
	err = tmpl.Execute(&codeBuf, cGenData)
	if err != nil {
		log.Fatalf("模板执行失败: %s\n", err.Error())
		return
	}
	// 代码格式化
	formatterCode, err := format.Source(codeBuf.Bytes())
	if err != nil {
		log.Fatalf("代码格式化失败: %s\n", err.Error())
		return
	}
	// 最后将字节写入文件
	_, err = io.WriteString(f, string(formatterCode))
	if err != nil {
		log.Fatalf("模板数据写入文件失败: %s\n", err.Error())
		return
	}
}

func createGenData(gen enum.GenEnum) map[string]interface{} {
	var res = make(map[string]interface{})
	res["PkgName"] = gen.PkgName
	res["EnumPath"] = gen.EnumPath
	res["valType"] = ""
	res["antType"] = ""
	res["Data"] = gen.Data
	// 根据gen.Data中的类型，赋值valType
	switch gen.Data[0].Val.(type) {
	case string:
		res["valType"] = "String"
	default:
		res["valType"] = "Int"
	}
	// 根据gen.Data中的类型，赋值antType
	switch gen.Data[0].Annotation.(type) {
	case string:
		res["antType"] = "String"
	default:
		res["antType"] = "Int"
	}
	fmt.Printf("%v\n", res)
	return res
}

// ------------------------------- 以下是带入模板的方法 -----------------------------

// 首字母转小写
func lowerFirst(s string) string {
	if len(s) > 0 {
		return strings.ToLower(s[:1]) + s[1:]
	}
	return s
}

// 首字母转大写
func upFirst(s string) string {
	if len(s) > 0 {
		return strings.ToUpper(s[:1]) + s[1:]
	}
	return s
}

// 下划线转驼峰
func camelCase(s string) string {
	parts := strings.Split(s, "_")
	var camelCaseParts []string
	for _, part := range parts {
		if part != "" {
			camelCaseParts = append(camelCaseParts, strings.Title(part))
		}
	}
	return strings.Join(camelCaseParts, "")
}

// 判断是否是字符串类型，是的话key返回"key"，否则返回key
func quoteIfString(val interface{}) interface{} {
	if str, ok := val.(string); ok {
		return "\"" + str + "\""
	}
	return val
}
