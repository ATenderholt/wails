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

	file := parser.NewAstFile(f)
	assert.Nil(t, file.ImportMap[parser.WailsApplicationPackage])
	assert.Equal(t, "_", *file.ImportMap["embed"])

	assert.Contains(t, file.FunctionCalls, "application.New")
	assert.Contains(t, file.FunctionCalls, "app.NewWebviewWindow")
	assert.Contains(t, file.FunctionCalls, "app.Run")
	assert.Contains(t, file.FunctionCalls, "panic")

	assert.Contains(t, file.CompositeLits, "application.Options")
}
