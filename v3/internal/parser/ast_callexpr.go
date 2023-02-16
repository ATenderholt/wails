package parser

import (
	"go/ast"
	"strings"
)

// AstCallExpr is type definition of ast.CallExpr with functions to make it easier to use
type AstCallExpr ast.CallExpr

// String returns the string representation of an ast.CallExpr function name
func (expr *AstCallExpr) String() string {
	// check if non-packaged (i.e. not a selector) function
	ident, ok := expr.Fun.(*ast.Ident)
	if ok {
		return ident.Name
	}

	selectorExpr, ok := expr.Fun.(*ast.SelectorExpr)
	if !ok {
		panic("AstCallExpr has Fun that cannot be cast to *ast.SelectorExpr")
	}

	if selectorExpr.Sel == nil {
		panic("AstCallExpr has nil Sel")
	}

	var builder strings.Builder
	if selectorExpr.X != nil {
		x, ok := selectorExpr.X.(*ast.Ident)
		if !ok {
			panic("AstCallExpr has X that cannot be cast to *ast.Ident")
		}
		builder.WriteString(x.Name)
		builder.WriteRune('.')
	}

	builder.WriteString(selectorExpr.Sel.Name)
	return builder.String()
}
