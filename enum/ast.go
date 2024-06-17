package enum

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

func inspectNode(n ast.Node) bool {
	switch x := n.(type) {
	case *ast.GenDecl:
		if x.Tok != token.CONST {
			break
		}
		for _, spec := range x.Specs {
			valSpec := spec.(*ast.ValueSpec)
			for i, name := range valSpec.Names {
				if len(valSpec.Values) > i {
					value := valSpec.Values[i]
					fmt.Printf("Constant %s is declared with value: ", name.Name)
					inspectExpr(value)
				}
			}
		}
	}
	return true
}

func inspectExpr(expr ast.Expr) {
	// 确定表达式的类型
	switch x := expr.(type) {
	case *ast.CompositeLit:
		fmt.Println("Type: Struct Literal")
		// 进一步分析结构体字段
		for _, elt := range x.Elts {
			key, val := parseKeyVal(elt)
			fmt.Printf("  %s: ", key)
			inspectExpr(val)
		}
	case *ast.BasicLit:
		fmt.Println("Type: Basic Literal - Value:", x.Value)
		// 处理基本字面量
	default:
		fmt.Println("Type: Other")
		// 处理其他类型的表达式
	}
}

// 解析 CompositeLit 中的键值对
func parseKeyVal(expr ast.Expr) (string, ast.Expr) {
	switch x := expr.(type) {
	case *ast.KeyValueExpr:
		return x.Key.(*ast.Ident).Name, x.Value
	default:
		// 如果不是键值对，返回空字符串和原始表达式
		return "", expr
	}
}

func astFile() {
	// 解析 Go 源文件
	fileSet := token.NewFileSet()
	file, err := parser.ParseFile(fileSet, "./gen.go", nil, parser.ParseComments)
	if err != nil {
		fmt.Printf("Error parsing file: %v\n", err)
		os.Exit(1)
	}

	ast.Inspect(file, inspectNode)
}
