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
	CompositeLits map[string][]*AstCompositeLit
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
	file.FunctionCalls = make(map[string][]*AstCallExpr)
	file.CompositeLits = make(map[string][]*AstCompositeLit)

	ast.Inspect(file.File, func(n ast.Node) bool {
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

		return true
	})
}
