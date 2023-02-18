package parser

import (
	"go/ast"
	"strings"
)

// AstFile is type definition of ast.File with functions to make it easier to use
type AstFile struct {
	*ast.File
	ImportMap     map[string]*string
	FunctionCalls map[string][]*AstCallExpr
}

func NewAstFile(file *ast.File) *AstFile {
	astFile := AstFile{
		File: file,
	}

	astFile.buildImportMap()
	astFile.buildFunctionCalls()

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

func (file *AstFile) buildFunctionCalls() {
	file.FunctionCalls = make(map[string][]*AstCallExpr)
	ast.Inspect(file.File, func(n ast.Node) bool {
		callExpr, ok := n.(*ast.CallExpr)
		if ok {
			wrapped := AstCallExpr(*callExpr)
			name := wrapped.String()
			list := file.FunctionCalls[name]
			list = append(list, &wrapped)
			file.FunctionCalls[name] = list
		}
		return true
	})
}
