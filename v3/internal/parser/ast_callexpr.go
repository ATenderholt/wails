package parser

import (
	"fmt"
	"go/ast"
)

// AstCallExpr is type definition of ast.CallExpr with functions to make it easier to use
type AstCallExpr ast.CallExpr

// String returns the string representation of an ast.CallExpr function name
func (expr *AstCallExpr) String() string {
	switch t := expr.Fun.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.SelectorExpr:
		return AstSelectorExpr(*t).String()
	default:
		msg := fmt.Sprintf("[AstCallExpr.String()] unsupported type for F: %T", t)
		panic(msg)
	}
}
