package parser

import (
	"fmt"
	"go/ast"
)

type BoundUnaryExpr ast.UnaryExpr

func (expr BoundUnaryExpr) String() string {
	switch t := expr.X.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.SelectorExpr:
		return AstSelectorExpr(*t).String()
	default:
		msg := fmt.Sprintf("[BoundUnaryExpr.String()] unsupported type for X: %T", t)
		panic(msg)
	}
}

func (expr BoundUnaryExpr) GetTypeSpec() *ast.TypeSpec {
	switch t := expr.X.(type) {
	case *ast.CompositeLit:
		return expr.getTypeSpecForCompositeLit(t)
	}
	return nil
}

func (expr BoundUnaryExpr) getTypeSpecForCompositeLit(lit *ast.CompositeLit) *ast.TypeSpec {
	switch t := lit.Type.(type) {
	case *ast.Ident:
		if t.Obj == nil {
			return nil
		}

		ts, ok := t.Obj.Decl.(*ast.TypeSpec)
		if ok {
			return ts
		}

		return nil
	}

	return nil
}