package parser

import (
	"fmt"
	"go/ast"
	"strings"
)

// AstFile is type definition of ast.File with functions to make it easier to use
type AstFile struct {
	*ast.File
	ImportMap     map[string]*string
	AssignStmts   map[string][]*AstAssignStmt
	CompositeLits map[string][]*AstCompositeLit
	FunctionCalls map[string][]*AstCallExpr
	TypeSpecs     map[string][]*AstTypeSpec
}

func NewAstFile(file *ast.File) *AstFile {
	astFile := AstFile{
		File: file,
	}

	astFile.buildImportMap()
	astFile.build()

	return &astFile

}

// buildImportMap iterates through Imports and builds a map of package names & import names
func (file *AstFile) buildImportMap() {
	file.ImportMap = make(map[string]*string)
	for _, importObj := range file.Imports {
		if importObj == nil || importObj.Path == nil {
			continue
		}

		key := strings.Trim(importObj.Path.Value, `"`)
		value := new(string)
		if importObj.Name == nil {
			value = nil
		} else {
			value = &importObj.Name.Name
		}

		file.ImportMap[key] = value
	}
}

func (file *AstFile) build() {
	file.AssignStmts = make(map[string][]*AstAssignStmt)
	file.FunctionCalls = make(map[string][]*AstCallExpr)
	file.CompositeLits = make(map[string][]*AstCompositeLit)
	file.TypeSpecs = make(map[string][]*AstTypeSpec)

	ast.Inspect(file.File, func(n ast.Node) bool {
		switch t := n.(type) {
		case *ast.TypeSpec:
			fmt.Printf("Found typespec\n")
		default:
			fmt.Printf("Found node of type %T\n", t)
		}

		stmt, ok := n.(*ast.AssignStmt)
		if ok {
			wrapped := AstAssignStmt(*stmt)
			name := wrapped.String()
			list := file.AssignStmts[name]
			list = append(list, &wrapped)
			file.AssignStmts[name] = list
			return true
		}

		callExpr, ok := n.(*ast.CallExpr)
		if ok {
			wrapped := AstCallExpr(*callExpr)
			name := wrapped.String()
			list := file.FunctionCalls[name]
			list = append(list, &wrapped)
			file.FunctionCalls[name] = list
			return true
		}

		compositeLit, ok := n.(*ast.CompositeLit)
		if ok {
			wrapped := AstCompositeLit(*compositeLit)
			name := wrapped.String()
			list := file.CompositeLits[name]
			list = append(list, &wrapped)
			file.CompositeLits[name] = list

			return true
		}

		typeSpec, ok := n.(*ast.TypeSpec)
		if ok {
			wrapped := AstTypeSpec(*typeSpec)
			name := wrapped.String()
			list := file.TypeSpecs[name]
			list = append(list, &wrapped)
			file.TypeSpecs[name] = list
		}

		return true
	})
}
