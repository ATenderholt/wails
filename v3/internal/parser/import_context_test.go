package parser_test

import (
	"github.com/stretchr/testify/require"
	"github.com/wailsapp/wails/v3/internal/parser"
	lang "go/parser"
	"go/token"
	"testing"
)

const importContextSrc = `
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

func TestNewImportContext(t *testing.T) {
	embed := "_"

	tests := []struct {
		name   string
		target string
		want   *parser.ImportContext
	}{
		{
			name:   "missing",
			target: "log",
			want:   nil,
		},
		{
			name:   "alias",
			target: "embed",
			want: &parser.ImportContext{
				Import:   "embed",
				ImportAs: &embed,
			},
		},
		{
			name:   "default",
			target: parser.WailsApplicationPackage,
			want: &parser.ImportContext{
				Import:   parser.WailsApplicationPackage,
				ImportAs: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fset := token.NewFileSet()
			file, err := lang.ParseFile(fset, "test.go", importContextSrc, lang.AllErrors)
			require.NoError(t, err, "unable to parse test src")

			importContext := parser.NewImportContext(file, tt.target)
			if tt.want == nil {
				require.Nil(t, importContext)
			} else {
				require.Equal(t, tt.want.Import, importContext.Import)
				require.Equal(t, tt.want.ImportAs, importContext.ImportAs)
				require.Equal(t, *parser.NewAstFile(file), *importContext.File)
			}
		})
	}
}
