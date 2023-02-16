package parser

import (
	"fmt"
	"go/ast"
)

// AstCompositeLit is type definition of ast.CompositeLit with functions to make it easier to use
type AstCompositeLit ast.CompositeLit

// String returns the string representation of an ast.CompositeLit name
func (lit AstCompositeLit) String() string {
	switch t := lit.Type.(type) {
	case *ast.SelectorExpr:
		return AstSelectorExpr(*t).String()
	case *ast.Ident:
		return t.Name
	case *ast.ArrayType:
		return "Array"
	default:
		msg := fmt.Sprintf("[AstCompositeLit.String()] unsupported type %T", t)
		panic(msg)
	}
}
