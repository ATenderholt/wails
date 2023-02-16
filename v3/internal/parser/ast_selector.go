package parser

import (
	"fmt"
	"go/ast"
)

// AstSelectorExpr is a type definition of ast.SelectorExpr with functions to make it easier to use
type AstSelectorExpr ast.SelectorExpr

// String return the string representation
func (expr AstSelectorExpr) String() string {
	if expr.Sel == nil {
		panic("[AstSelectorExpr.String()] expr.Sel is nil")
	}

	switch t := expr.X.(type) {
	case *ast.Ident:
		return t.Name + "." + expr.Sel.Name
	default:
		msg := fmt.Sprintf("[AstSelectorExpr.X] is not an *ast.Ident: %T", t)
		panic(msg)
	}
}
