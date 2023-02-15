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

func (file AstFile) GetFunctionCalls() []*AstCallExpr {
	doneChan := make(chan bool)
	funcChan := make(chan *ast.CallExpr)

	go func() {
		unwrapped := ast.File(file)
		ast.Inspect(&unwrapped, func(n ast.Node) bool {
			callExpr, ok := n.(*ast.CallExpr)
			if ok {
				funcChan <- callExpr
			}
			return true
		})
		doneChan <- true
	}()

	var funcs []*AstCallExpr
	for {
		select {
		case callExpr := <-funcChan:
			wrapped := AstCallExpr(*callExpr)
			funcs = append(funcs, &wrapped)
		case <-doneChan:
			return funcs
		}
	}
}
