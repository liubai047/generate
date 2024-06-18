package test

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"strings"
	"testing"
)

func TestAstFile(t *testing.T) {
	// println("start")
	// v, err := json.Marshal(genReportStatus)
	// println(string(v), err)
	// main()
	// main2()
	str, err := astStruct("./gen.go", "genReportStatus")
	fmt.Printf("%#v,%x\n", str, err)
}

func main() {
	fSet := token.NewFileSet()
	node, err := parser.ParseFile(fSet, "gen.go", nil, parser.ParseComments)
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return
	}
	// 遍历 AST 并查找 genReportStatus 变量
	var genReportStatus *ast.ValueSpec
	ast.Inspect(node, func(n ast.Node) bool {
		if ident, ok := n.(*ast.Ident); ok && ident.Name == "genReportStatus" {
			if spec, ok := ident.Obj.Decl.(*ast.ValueSpec); ok {
				genReportStatus = spec
				return false
			}
		}
		return true
	})
	if genReportStatus == nil {
		fmt.Println("genReportStatus not found")
		return
	}
	// 获取 genReportStatus 的值
	if len(genReportStatus.Values) == 0 {
		fmt.Println("genReportStatus has no value")
		return
	}
	value := genReportStatus.Values[0]
	mapData := make(map[string]interface{})
	// 解析 genReportStatus 的值并写入 map
	parseValue(value, mapData)
	// 打印结果
	for k, v := range mapData {
		fmt.Printf("%s: %v\n", k, v)
	}
}

func parseValue(value ast.Expr, result map[string]interface{}) {
	switch v := value.(type) {
	case *ast.CompositeLit:
		fmt.Printf("v : %#v\n", v)
		for k, elt := range v.Elts {
			fmt.Printf("elt : %#v\n", elt)
			fmt.Printf("k : %#v\n", k)
			switch kv := elt.(type) {
			case *ast.KeyValueExpr:
				key := fmt.Sprint(kv.Key)
				val := kv.Value
				parseKeyValue(key, val, result)
			case *ast.CompositeLit:
				parseValue(elt, result)
			default:
				fmt.Printf("未知类型的值:%#v", kv)
				return
			}
		}
	case *ast.BasicLit:
		// 处理基本类型
		result[""] = v.Value
	case *ast.Ident:
		// 处理标识符
		result[""] = v.Name
	default:
		fmt.Printf("Unsupported type: %T\n", v)
	}
}

func parseKeyValue(key string, value ast.Expr, result map[string]interface{}) {
	switch v := value.(type) {
	case *ast.CompositeLit:
		println(key)
		subMap := make(map[string]interface{})
		parseValue(v, subMap)
		result[key] = subMap
	case *ast.ArrayType:
		println(key, "ArrayType")
	case *ast.BasicLit:
		// 处理基本类型
		result[key] = v.Value
	case *ast.Ident:
		// 处理标识符
		result[key] = v.Name
	default:
		fmt.Printf("Unsupported type in key-value pair: %T\n", v)
	}
}

func main2() {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "gen.go", nil, parser.ParseComments)
	if err != nil {
		log.Fatalf("Failed to parse file: %v", err)
	}

	var genReportStatus *ast.CompositeLit
	ast.Inspect(node, func(n ast.Node) bool {
		if decl, ok := n.(*ast.GenDecl); ok && decl.Tok == token.VAR {
			for _, spec := range decl.Specs {
				if valueSpec, ok := spec.(*ast.ValueSpec); ok {
					for i, name := range valueSpec.Names {
						if name.Name == "genReportStatus" {
							if compositeLit, ok := valueSpec.Values[i].(*ast.CompositeLit); ok {
								genReportStatus = compositeLit
								return false
							}
						}
					}
				}
			}
		}
		return true
	})

	if genReportStatus == nil {
		log.Fatal("genReportStatus not found")
	}

	result := make(map[string]interface{})
	for _, elt := range genReportStatus.Elts {
		if kv, ok := elt.(*ast.KeyValueExpr); ok {
			key := kv.Key.(*ast.Ident).Name
			value := extractValue2(kv.Value)
			result[key] = value
		}
	}

	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}
	fmt.Println(string(jsonData))
}

func extractValue2(expr ast.Expr) interface{} {
	switch v := expr.(type) {
	case *ast.BasicLit:
		switch v.Kind {
		case token.STRING:
			return strings.Trim(v.Value, `"`)
		case token.INT:
			return v.Value
		case token.FLOAT:
			return v.Value
		default:
			panic("unhandled default case")
		}
	case *ast.CompositeLit:
		if _, ok := v.Type.(*ast.ArrayType); ok {
			var items []interface{}
			for _, elt := range v.Elts {
				items = append(items, extractValue(elt))
			}
			return items
		} else {
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
