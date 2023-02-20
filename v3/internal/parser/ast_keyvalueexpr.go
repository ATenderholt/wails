package parser

import (
	"fmt"
	"go/ast"
)

type AstKeyValueExpr ast.KeyValueExpr

func (expr *AstKeyValueExpr) Name() string {
	switch t := expr.Key.(type) {
	case *ast.Ident:
		return t.Name
	default:
		msg := fmt.Sprintf("[AstKeyValueExpr.Name()]: unsupported type for Key: %T", t)
		panic(msg)
	}
}

func (expr *AstKeyValueExpr) ArrayValue() []ast.Expr {
	switch t := expr.Value.(type) {
	case *ast.CompositeLit:
		return AstCompositeLit(*t).Elts
	default:
		msg := fmt.Sprintf("[AstKeyValueExpr.ArrayValue()]: unsupported type for Value: %T", t)
		panic(msg)
	}
}
