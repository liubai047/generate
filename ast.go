package main

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
	"strings"
)

// 入口。解析制定文件中指定结构体，将结果放在map中后，进行序列化，最后转为string类型数据
func astStruct(fileName string, structName string) ([]byte, error) {
	fSet := token.NewFileSet()
	node, err := parser.ParseFile(fSet, fileName, nil, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("打开结构体文件失败: %v", err)
	}
	// 定位待解析结构体
	var astData *ast.CompositeLit
	ast.Inspect(node, func(node ast.Node) bool {
		decl, ok := node.(*ast.GenDecl)
		if !ok || decl.Tok != token.VAR {
			return true
		}
		for _, spec := range decl.Specs {
			valueSpec, ok := spec.(*ast.ValueSpec)
			if !ok {
				return true
			}
			for i, name := range valueSpec.Names {
				if name.Name == structName {
					compositeLit, ok := valueSpec.Values[i].(*ast.CompositeLit)
					if !ok {
						return true
					}
					astData = compositeLit
					return false
				}
			}
		}
		return true
	})
	// 未定位到结构体
	if astData == nil {
		return nil, fmt.Errorf("未定位到结构体")
	}
	// 初始化返回值
	result := make(map[string]interface{})
	// 解析结构体各个字段数据
	for _, elt := range astData.Elts {
		if kv, ok := elt.(*ast.KeyValueExpr); ok {
			key := kv.Key.(*ast.Ident).Name
			value := extractValue(kv.Value)
			result[key] = value
		}
	}
	jsonData, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("结果序列化失败: %v", err)
	}
	return jsonData, nil
}

// 对结果进行赋值
func extractValue(expr ast.Expr) interface{} {
	switch v := expr.(type) {
	// 字面量直接提取结果进行赋值操作
	case *ast.BasicLit:
		switch v.Kind {
		case token.STRING:
			return strings.Trim(v.Value, `"`)
		case token.INT:
			vVal, err := strconv.Atoi(v.Value)
			if err != nil {
				panic("vVal INT类型错误")
			}
			return vVal
		case token.FLOAT:
			vVal, err := strconv.ParseFloat(v.Value, 64)
			if err != nil {
				panic("vVal FLOAT类型错误")
			}
			return vVal
		default:
			panic("类型未知")
		}
	// 如果是嵌套结构体，需要递归进行处理
	case *ast.CompositeLit:
		// 如果是数组结构体类型
		if _, ok := v.Type.(*ast.ArrayType); ok {
			var items []interface{}
			for _, elt := range v.Elts {
				items = append(items, extractValue(elt))
			}
			return items
		} else { // 普通的结构体嵌套类型
			itemMap := make(map[string]interface{})
			for _, elt := range v.Elts {
				if kv, ok := elt.(*ast.KeyValueExpr); ok {
					key := kv.Key.(*ast.Ident).Name
					value := extractValue(kv.Value)
					itemMap[key] = value
				}
			}
			return itemMap
		}
	}
	return nil
}
