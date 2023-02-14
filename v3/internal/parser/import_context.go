package parser

import (
	"go/ast"
	"strings"
)

const WailsApplicationPackage = "github.com/wailsapp/wails/v3/pkg/application"

type ImportContext struct {
	Import   string
	ImportAs *string
	File     *AstFile
}

func NewImportContext(file *ast.File, target string) *ImportContext {
	for _, importObj := range file.Imports {
		if importObj == nil || importObj.Path == nil {
			continue
		}

		value := strings.Trim(importObj.Path.Value, `"`)
		if value == target {
			importAs := new(string)

			if importObj.Name == nil {
				importAs = nil
			} else {
				importAs = &importObj.Name.Name
			}

			f := AstFile(*file)
			return &ImportContext{
				Import:   value,
				ImportAs: importAs,
				File:     &f,
			}
		}
	}

	return nil
}
