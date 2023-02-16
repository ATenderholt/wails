package parser

import (
	"go/ast"
	"strings"
)

// AstFile is type definition of ast.File with functions to make it easier to use
type AstFile ast.File

// ImportMap returns a map of import names and their aliases (if any)
func (file AstFile) ImportMap() map[string]*string {
	imports := make(map[string]*string)

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

		imports[key] = value
	}
	return imports
}

func (file AstFile) GetFunctionCalls() map[string][]*AstCallExpr {
	callExprs := make(map[string][]*AstCallExpr)
	unwrapped := ast.File(file)
	ast.Inspect(&unwrapped, func(n ast.Node) bool {
		callExpr, ok := n.(*ast.CallExpr)
		if ok {
			wrapped := AstCallExpr(*callExpr)
			name := wrapped.String()
			list := callExprs[name]
			list = append(list, &wrapped)
			callExprs[name] = list
		}
		return true
	})

	return callExprs
}
