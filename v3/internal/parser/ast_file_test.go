package parser_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wailsapp/wails/v3/internal/parser"
	lang "go/parser"
	"go/token"
	"testing"
)

const astFileTestSrc = `
package main

import (
	_ "embed"
	"github.com/wailsapp/wails/v3/pkg/application"
)

func main() {
	app := application.New(application.Options{
		Bind: []interface{}{},
	})

	app.NewWebviewWindow()

	err := app.Run()

	if err != nil {
		panic(err)
	}
}
`

func TestAstFile(t *testing.T) {
	fset := token.NewFileSet()
	f, err := lang.ParseFile(fset, "test.go", astFileTestSrc, lang.AllErrors)
	require.NoError(t, err, "unable to parse test src")

	file := parser.AstFile(*f)
	imports := file.ImportMap()
	assert.Nil(t, imports[parser.WailsApplicationPackage])
	assert.Equal(t, "_", *imports["embed"])

	funcs := file.GetFunctionCalls()
	funcNames := make([]string, len(funcs))
	for i, f := range funcs {
		funcNames[i] = f.String()
	}

	assert.Contains(t, funcNames, "application.New")
	assert.Contains(t, funcNames, "app.NewWebviewWindow")
	assert.Contains(t, funcNames, "app.Run")
	assert.Contains(t, funcNames, "panic")
}
