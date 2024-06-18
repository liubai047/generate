package main

import (
	"encoding/json"
	"flag"
	"log"
	"testing"

	"generate/enum"
)

func TestMainTry(m *testing.T) {
	tmplFile := "./gen.go.tmpl"
	structFile := "./enum/gen.go"
	structName := "genReportStatus"
	dstFile := "./ReportStatusEnum/c.go"
	mainTry(&tmplFile, &structFile, &structName, &dstFile)
}

func mainTry(tmplFile, structFile, structName, dstFile *string) {
	// tmplFile := flag.String("tmplFile", "./gen.go.tmpl", "一期不填.模板文件位置，推荐使用相对路径，例如：./template/a.go.tmpl")
	// structFile := flag.String("struct_file", "", "结构体所在文件路径，也推荐使用相对路径")
	// structName := flag.String("struct_name", "", "结构体所在文件名字")
	// dstFile := flag.String("dstFile", "", "生成文件所在位置，默认为./genData.PkgName/enum.go")
	flag.Parse()
	if *tmplFile == "" || *structFile == "" || *structName == "" {
		log.Fatalf("参数有误")
	}
	data, err := astStruct(*structFile, *structName)
	if err != nil {
		log.Fatalln(err.Error())
	}
	var genData enum.GenEnum
	err = json.Unmarshal(data, &genData)
	if err != nil {
		log.Fatalf("反序列化数据到genData失败:%s", err.Error())
	}
	// 处理待写入路径
	if *dstFile == "" {
		*dstFile = "./" + genData.PkgName + "/enum.go"
	}
	generateEnum(*dstFile, genData)
}
