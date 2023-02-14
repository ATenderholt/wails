package parser_test

import (
	"github.com/stretchr/testify/require"
	"github.com/wailsapp/wails/v3/internal/parser"
	"go/ast"
	lang "go/parser"
	"go/token"
	"testing"
)

const astTypesTestSrc = `
package main

import (
	"log"
)

func f() bool {
	return false
}

func main() {
	f()
	log.Print("Hello world!")
}
`

func TestAstCallExpr_String(t *testing.T) {
	tests := []struct {
		name  string
		index int
		want  string
	}{
		{
			name:  "f()",
			index: 0,
			want:  "f",
		},
		{
			name:  "log.Print()",
			index: 1,
			want:  "log.Print",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fset := token.NewFileSet()
			file, err := lang.ParseFile(fset, "test.go", astTypesTestSrc, lang.AllErrors)
			require.NoError(t, err)

			// 0 - import; 1 - func f(); 2 - main()
			f := file.Decls[2].(*ast.FuncDecl)
			stmt := f.Body.List[tt.index].(*ast.ExprStmt)
			temp := stmt.X.(*ast.CallExpr)
			x := parser.AstCallExpr(*temp)
			require.Equal(t, tt.want, x.String())
		})
	}
}
