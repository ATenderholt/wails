package parser

import (
	"fmt"
	"go/ast"
)

type AstAssignStmt ast.AssignStmt

func (stmt *AstAssignStmt) String() string {
	switch t := stmt.Lhs[0].(type) {
	case *ast.Ident:
		return t.Name
	default:
		fmt.Printf("[AstAssignStmt.String()] not handling %T types yet", t)
		return ""
	}
}
