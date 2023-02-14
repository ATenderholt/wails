package parser

import (
	"go/ast"
	"strings"
)

// AstCallExpr extends ast.CallExpr with functions to make it easier to use
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

type AstFile ast.File

//func (file AstFile) HasNewApplication() bool {
//	for _, decl := range file.Decls {
//		funcDecl, ok := decl.(*ast.FuncDecl)
//		if !ok {
//			continue
//		}
//
//		for _, stmt := range funcDecl.Body.List {
//			assignStmt, ok := stmt.(*ast.AssignStmt)
//			if !ok {
//				continue
//			}
//
//			for
//		}
//	}
//	return false
//}
