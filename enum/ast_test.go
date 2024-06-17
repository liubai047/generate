package enum

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"testing"
)

func TestAstFile(t *testing.T) {
	main()
}

var isLoop = false

func main() {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "gen.go", nil, parser.ParseComments)
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

func parseArrValue(value ast.Expr, result []map[string]interface{}) {
	switch v := value.(type) {
	case *ast.CompositeLit:
		for k, elt := range v.Elts {
			var vMap = map[string]interface{}{}
			result = append(result, vMap)
			switch kv := elt.(type) {
			case *ast.KeyValueExpr:
				key := fmt.Sprint(kv.Key)
				val := kv.Value
				parseKeyValue(key, val, vMap, k)
			default:
				fmt.Printf("未知类型的值:%#v", kv)
				return
			}
			//if kv, ok := elt.(*ast.KeyValueExpr); ok {
			//	key := fmt.Sprint(kv.Key)
			//	val := kv.Value
			//	parseKeyValue(key, val, result, k)
			//}
		}
	}
}

func parseValue(value ast.Expr, result map[string]interface{}) {
	switch v := value.(type) {
	case *ast.CompositeLit:
		//fmt.Printf("v : %#v\n", v)
		for k, elt := range v.Elts {
			//fmt.Printf("elt : %#v\n", elt)
			//fmt.Printf("k : %#v\n", k)
			switch kv := elt.(type) {
			case *ast.KeyValueExpr:
				key := fmt.Sprint(kv.Key)
				val := kv.Value
				parseKeyValue(key, val, result, k)
			case *ast.CompositeLit:
				parseValue(elt, result)
			default:
				fmt.Printf("未知类型的值:%#v", kv)
				return
			}
			//if kv, ok := elt.(*ast.KeyValueExpr); ok {
			//	key := fmt.Sprint(kv.Key)
			//	val := kv.Value
			//	parseKeyValue(key, val, result, k)
			//}
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

//func parseKeyArrValue(key string, value ast.Expr, result []map[string]interface{}, idx int) {
//	switch v := value.(type) {
//	case *ast.CompositeLit:
//		println(key)
//		subMap := make(map[string]interface{})
//		parseValue(v, subMap)
//		result[key] = subMap
//	case *ast.ArrayType:
//		println(key, "ArrayType")
//	case *ast.BasicLit:
//		// 处理基本类型
//		result[key] = v.Value
//	case *ast.Ident:
//		// 处理标识符
//		result[key] = v.Name
//	default:
//		fmt.Printf("Unsupported type in key-value pair: %T\n", v)
//	}
//}

func parseKeyValue(key string, value ast.Expr, result map[string]interface{}, idx int) {
	switch v := value.(type) {
	case *ast.CompositeLit:
		//println(key)
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

func serialize(v interface{}) string {
	var buf bytes.Buffer
	err := format.Node(&buf, token.NewFileSet(), v)
	if err != nil {
		fmt.Println("Error serializing value:", err)
		return ""
	}
	return buf.String()
}
